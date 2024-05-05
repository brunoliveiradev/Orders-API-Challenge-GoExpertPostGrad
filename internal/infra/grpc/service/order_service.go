package service

import (
	"GoExpertPostGrad-Orders-Challenge/internal/infra/grpc/pb"
	"GoExpertPostGrad-Orders-Challenge/internal/usecase"
	"context"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	input := usecase.OrderInputDTO{
		ID:    req.Id,
		Price: float64(req.Price),
		Tax:   float64(req.Tax),
	}

	output, err := s.CreateOrderUseCase.Execute(input)
	if err != nil {
		return nil, err
	}

	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}
