//go:build wireinject
// +build wireinject

package main

//
//import (
//	"GoExpertPostGrad-Orders-Challenge/internal/entity"
//	"GoExpertPostGrad-Orders-Challenge/internal/event"
//	"GoExpertPostGrad-Orders-Challenge/internal/infra/database"
//	"GoExpertPostGrad-Orders-Challenge/internal/infra/rest"
//	"GoExpertPostGrad-Orders-Challenge/internal/usecase"
//	"GoExpertPostGrad-Orders-Challenge/pkg/events"
//	"database/sql"
//
//	"github.com/google/wire"
//)
//
//var setOrderRepositoryDependency = wire.NewSet(
//	database.NewOrderRepository,
//	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
//)
//
//var setEventDispatcherDependency = wire.NewSet(
//	events.NewEventDispatcher,
//	event.NewOrderCreated,
//	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
//	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
//)
//
//var setOrderCreatedEvent = wire.NewSet(
//	event.NewOrderCreated,
//	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
//)
//
//func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
//	wire.Build(
//		setOrderRepositoryDependency,
//		setOrderCreatedEvent,
//		usecase.NewCreateOrderUseCase,
//	)
//	return &usecase.CreateOrderUseCase{}
//}
//
//func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *rest.WebOrderHandler {
//	wire.Build(
//		setOrderRepositoryDependency,
//		setOrderCreatedEvent,
//		rest.NewWebOrderHandler,
//	)
//	return &rest.WebOrderHandler{}
//}
