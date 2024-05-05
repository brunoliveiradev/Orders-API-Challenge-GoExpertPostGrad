package web

import (
	"GoExpertPostGrad-Orders-Challenge/internal/entity"
	"GoExpertPostGrad-Orders-Challenge/internal/usecase"
	"GoExpertPostGrad-Orders-Challenge/pkg/events"
	"encoding/json"
	"net/http"
)

type WebOrderHandler struct {
	EventDispatcher   events.EventDispatcherInterface
	OrderRepository   entity.OrderRepositoryInterface
	OrderCreatedEvent events.EventInterface
}

func NewWebOrderHandler(eventDispatcher events.EventDispatcherInterface, orderRepository entity.OrderRepositoryInterface, orderCreatedEvent events.EventInterface) *WebOrderHandler {
	return &WebOrderHandler{
		EventDispatcher:   eventDispatcher,
		OrderRepository:   orderRepository,
		OrderCreatedEvent: orderCreatedEvent,
	}
}

func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input usecase.OrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Create order
	createOrder := usecase.NewCreateOrderUseCase(h.OrderRepository, h.OrderCreatedEvent, h.EventDispatcher)
	output, err := createOrder.Execute(input)
	if err != nil {
		http.Error(w, "Error creating order: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
