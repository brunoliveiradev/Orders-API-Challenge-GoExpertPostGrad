package main

import (
	"GoExpertPostGrad-Orders-Challenge/configs"
	"GoExpertPostGrad-Orders-Challenge/internal/entity"
	"GoExpertPostGrad-Orders-Challenge/internal/event"
	"GoExpertPostGrad-Orders-Challenge/internal/event/handler"
	"GoExpertPostGrad-Orders-Challenge/internal/infra/database"
	"GoExpertPostGrad-Orders-Challenge/internal/infra/graphql/graph"
	"GoExpertPostGrad-Orders-Challenge/internal/infra/grpc/pb"
	"GoExpertPostGrad-Orders-Challenge/internal/infra/grpc/service"
	"GoExpertPostGrad-Orders-Challenge/internal/infra/rest"
	"GoExpertPostGrad-Orders-Challenge/internal/infra/rest/web"
	"GoExpertPostGrad-Orders-Challenge/internal/usecase"
	"GoExpertPostGrad-Orders-Challenge/pkg/cache"
	"GoExpertPostGrad-Orders-Challenge/pkg/events"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"

	graphqlhandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, db, rabbitMQChannels, eventDispatcher, memoryCache := initialize()

	defer db.Close()
	defer rabbitMQChannels.Close()

	orderCreationUseCase, orderListingUseCase, orderRepository, orderCreatedEvent, orderListingEvent := setupUseCases(db, eventDispatcher, memoryCache)

	startServers(config, eventDispatcher, orderCreationUseCase, orderListingUseCase, orderRepository, orderCreatedEvent, orderListingEvent, memoryCache)
}

func initialize() (*configs.Envs, *sql.DB, *events.RabbitMQService, *events.EventDispatcher, cache.Cache) {
	config := loadConfig()
	db := setupDatabase(config)
	rabbitMQChannels := setupRabbitMQService(config)
	memoryCache := cache.NewMemoryCache()
	eventDispatcher := setupEventDispatcher(rabbitMQChannels, memoryCache)
	return config, db, rabbitMQChannels, eventDispatcher, memoryCache
}

func setupUseCases(db *sql.DB, eventDispatcher *events.EventDispatcher, memoryCache cache.Cache) (*usecase.OrderCreationUseCase, *usecase.OrderListingUseCase, entity.OrderRepositoryInterface, usecase.OrderCreatedEventInterface, usecase.OrderListEventInterface) {
	orderRepository := database.NewOrderRepository(db)
	orderCreatedEvent := event.NewOrderCreated()
	orderCreationUseCase := usecase.NewOrderCreationUseCase(orderRepository, orderCreatedEvent, eventDispatcher)

	orderListingEvent := event.NewOrderListingEvent()
	orderListingUseCase := usecase.NewOrderListAllUseCase(orderRepository, orderListingEvent, eventDispatcher, memoryCache)

	return orderCreationUseCase, orderListingUseCase, orderRepository, orderCreatedEvent, orderListingEvent
}

func loadConfig() *configs.Envs {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	return config
}

func setupDatabase(config *configs.Envs) *sql.DB {
	db, err := database.SetupDatabase(config)
	if err != nil {
		log.Fatalf("failed to setup database: %v", err)
	}
	return db
}

func setupRabbitMQService(config *configs.Envs) *events.RabbitMQService {
	return events.NewRabbitMQService(config, "orders")
}

func setupEventDispatcher(rabbitMQChannels *events.RabbitMQService, memoryCache cache.Cache) *events.EventDispatcher {
	eventDispatcher := events.NewEventDispatcher()

	orderCreatedHandler := handler.NewOrderCreatedHandler(rabbitMQChannels, memoryCache)
	eventDispatcher.Register(orderCreatedHandler.GetEventName(), orderCreatedHandler)

	orderListingHandler := handler.NewOrderListingHandler(rabbitMQChannels)
	eventDispatcher.Register(orderListingHandler.GetEventName(), orderListingHandler)

	return eventDispatcher
}

func startServers(config *configs.Envs, eventDispatcher *events.EventDispatcher, orderCreationUseCase *usecase.OrderCreationUseCase, orderListingUseCase *usecase.OrderListingUseCase, orderRepository entity.OrderRepositoryInterface, orderCreatedEvent usecase.OrderCreatedEventInterface, orderListingEvent usecase.OrderListEventInterface, memoryCache cache.Cache) {
	startWebServer(config, eventDispatcher, orderRepository, orderCreatedEvent, orderListingEvent, memoryCache)
	startGRPCServer(config, orderCreationUseCase, orderListingUseCase)
	startGraphQLServer(config, orderCreationUseCase, orderListingUseCase)

	fmt.Println("Starting servers...")
}

func startWebServer(config *configs.Envs, eventDispatcher *events.EventDispatcher, orderRepository entity.OrderRepositoryInterface, orderCreatedEvent usecase.OrderCreatedEventInterface, orderListingEvent usecase.OrderListEventInterface, memoryCache cache.Cache) {
	webserver := web.NewWebServer(config.WebServerPort)
	webOrderHandler := rest.NewWebOrderHandler(eventDispatcher, orderRepository, orderCreatedEvent, orderListingEvent, memoryCache)

	webserver.AddHandler("/order", webOrderHandler.Create)
	webserver.AddHandler("/orders", webOrderHandler.ListAll)

	fmt.Println("Starting web server on port", config.WebServerPort)
	go webserver.Start()
}

func startGRPCServer(config *configs.Envs, createOrderUseCase *usecase.OrderCreationUseCase, listOrdersUseCase *usecase.OrderListingUseCase) {
	grpcServer := grpc.NewServer()
	newOrderService := service.NewOrderService(createOrderUseCase, listOrdersUseCase)
	pb.RegisterOrderServiceServer(grpcServer, newOrderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", config.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)
}

func startGraphQLServer(config *configs.Envs, createOrderUseCase *usecase.OrderCreationUseCase, orderListingUseCase *usecase.OrderListingUseCase) {
	srv := graphqlhandler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		OrderCreationUseCase: *createOrderUseCase,
		OrderListingUseCase:  *orderListingUseCase,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", config.GraphQLServerPort)
	if err := http.ListenAndServe(":"+config.GraphQLServerPort, nil); err != nil {
		log.Fatalf("Failed to start GraphQL server: %v", err)
	}
}
