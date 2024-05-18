package main

import (
	"GoExpertPostGrad-Orders-Challenge/configs"
	"GoExpertPostGrad-Orders-Challenge/internal/event/handler"
	"GoExpertPostGrad-Orders-Challenge/internal/infra/database"
	"GoExpertPostGrad-Orders-Challenge/internal/infra/graphql/graph"
	"GoExpertPostGrad-Orders-Challenge/internal/infra/grpc/pb"
	"GoExpertPostGrad-Orders-Challenge/internal/infra/grpc/service"
	"GoExpertPostGrad-Orders-Challenge/internal/infra/rest/web"
	"GoExpertPostGrad-Orders-Challenge/internal/usecase"
	"GoExpertPostGrad-Orders-Challenge/pkg/events"
	"database/sql"
	"fmt"
	graphqlhandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := database.SetupDatabase(config)
	if err != nil {
		log.Fatalf("failed to setup database: %v", err)
	}
	defer db.Close()

	rabbitMQChannel := getRabbitMQChannel(config)
	defer rabbitMQChannel.Close()

	orderCreatedHandler := handler.NewOrderCreatedHandler(rabbitMQChannel)

	eventDispatcher := setupEventDispatcher(orderCreatedHandler)

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)

	startWebServer(config, db, eventDispatcher)
	startGRPCServer(config, createOrderUseCase)
	startGraphQLServer(config, createOrderUseCase)
}

func getRabbitMQChannel(config *configs.Envs) *amqp.Channel {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", config.RabbitMQUser, config.RabbitMQPassword, config.RabbitMQHost, config.RabbitMQPort)
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open a channel: %v", err)
	}
	return ch
}

func setupEventDispatcher(orderCreatedHandler *handler.OrderCreatedHandler) *events.EventDispatcher {
	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", orderCreatedHandler)
	return eventDispatcher
}

func startWebServer(config *configs.Envs, db *sql.DB, eventDispatcher *events.EventDispatcher) {
	webserver := web.NewWebServer(config.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)

	webserver.AddHandler("/order", webOrderHandler.Create)

	fmt.Println("Starting web server on port", config.WebServerPort)
	go webserver.Start()
}

func startGRPCServer(config *configs.Envs, createOrderUseCase *usecase.CreateOrderUseCase) {
	grpcServer := grpc.NewServer()
	createOrderService := service.NewOrderService(*createOrderUseCase)
	pb.RegisterOrderServiceServer(grpcServer, createOrderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", config.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)
}

func startGraphQLServer(config *configs.Envs, createOrderUseCase *usecase.CreateOrderUseCase) {
	srv := graphqlhandler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", config.GraphQLServerPort)
	http.ListenAndServe(":"+config.GraphQLServerPort, nil)
}
