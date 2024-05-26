package service

import (
	"GoExpertPostGrad-Orders-Challenge/internal/infra/grpc/pb"
	"GoExpertPostGrad-Orders-Challenge/internal/usecase"
	"context"
)

var cacheKey = "orders"

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	createOrderUseCase *usecase.OrderCreationUseCase
	listOrdersUseCase  *usecase.OrderListingUseCase
}

func NewOrderService(createOrderUseCase *usecase.OrderCreationUseCase, listOrdersUseCase *usecase.OrderListingUseCase) *OrderService {
	return &OrderService{
		createOrderUseCase: createOrderUseCase,
		listOrdersUseCase:  listOrdersUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		Name:  req.Name,
		Price: float64(req.Price),
		Tax:   float64(req.Tax),
	}

	output, err := s.createOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}

	// Clear cache to ensure consistency
	s.listOrdersUseCase.Cache.Delete(cacheKey)

	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Name:       output.Name,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	output, err := s.listOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var orders []*pb.Order
	for _, order := range output.Orders {
		orders = append(orders, &pb.Order{
			Id:         order.ID,
			Name:       order.Name,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		})
	}

	return &pb.ListOrdersResponse{Orders: orders}, nil
}
