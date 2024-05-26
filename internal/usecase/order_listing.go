package usecase

import (
	"GoExpertPostGrad-Orders-Challenge/internal/entity"
	"GoExpertPostGrad-Orders-Challenge/pkg/cache"
	"GoExpertPostGrad-Orders-Challenge/pkg/events"
	"fmt"
	"log"
	"time"
)

const cacheKey = "orders"

type OrdersDTO struct {
	Orders []*OrderOutputDTO
}

type OrderListEventInterface interface {
	events.EventInterface
}

type OrderListingUseCase struct {
	OrderRepository   entity.OrderRepositoryInterface
	OrderListAllEvent OrderListEventInterface
	EventDispatcher   events.EventDispatcherInterface
	Cache             cache.Cache
}

func NewOrderListAllUseCase(orderRepository entity.OrderRepositoryInterface, orderListAllEvent OrderListEventInterface, eventDispatcher events.EventDispatcherInterface, cache cache.Cache) *OrderListingUseCase {
	return &OrderListingUseCase{
		OrderRepository:   orderRepository,
		OrderListAllEvent: orderListAllEvent,
		EventDispatcher:   eventDispatcher,
		Cache:             cache,
	}
}

func (uc *OrderListingUseCase) Execute() (OrdersDTO, error) {
	if cachedOrders, ok := uc.Cache.Get(cacheKey); ok {
		log.Println("Orders fetched from cache")
		return cachedOrders.(OrdersDTO), nil
	}

	log.Println("Orders not found in cache, fetching from database...")
	orders, err := uc.fetchAndCacheOrders()
	if err != nil {
		return OrdersDTO{}, fmt.Errorf("failed to fetch and cache orders: %w", err)
	}

	return orders, nil
}

func (uc *OrderListingUseCase) fetchAndCacheOrders() (OrdersDTO, error) {
	startTime := time.Now()

	log.Println("Listing all orders from the repository...")
	orders, err := uc.OrderRepository.ListAllOrders()
	if err != nil {
		return OrdersDTO{}, fmt.Errorf("failed to list all orders: %w", err)
	}

	dto := OrdersDTO{}
	for _, order := range orders {
		dto.Orders = append(dto.Orders, toDTOOrder(order))
	}

	duration := time.Since(startTime)

	uc.OrderListAllEvent.SetPayload(duration)
	err = uc.EventDispatcher.Dispatch(uc.OrderListAllEvent)
	if err != nil {
		return OrdersDTO{}, fmt.Errorf("failed to dispatch event: %w", err)
	}

	uc.Cache.Set(cacheKey, dto)
	log.Println("Orders fetched and cached successfully, duration: ", duration)
	return dto, nil
}

func toDTOOrder(order *entity.Order) *OrderOutputDTO {
	return &OrderOutputDTO{
		ID:         order.ID,
		Name:       order.Name,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}
}
