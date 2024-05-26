package event

import "time"

type OrderListingEvent struct {
	Name    string
	Payload interface{}
}

func NewOrderListingEvent() *OrderListingEvent {
	return &OrderListingEvent{
		Name: "OrderListingEvent",
	}
}

func (e *OrderListingEvent) GetName() string {
	return e.Name
}

func (e *OrderListingEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *OrderListingEvent) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *OrderListingEvent) GetDateTime() time.Time {
	return time.Now()
}
