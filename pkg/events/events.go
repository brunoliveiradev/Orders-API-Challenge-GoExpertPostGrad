package events

import (
	"errors"
	"log"
	"sync"
	"time"
)

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
	SetPayload(payload interface{})
}

type EventHandlerInterface interface {
	Handle(event EventInterface) error
}

type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dispatch(event EventInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Clear()
}

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) Dispatch(event EventInterface) error {
	if handlers, ok := ed.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}
		for _, handler := range handlers {
			wg.Add(1)
			go func(h EventHandlerInterface) {
				defer wg.Done()
				if err := h.Handle(event); err != nil {
					log.Printf("Error handling event %s: %v", event.GetName(), err)
				}
			}(handler)
		}
		wg.Wait()
	}
	return nil
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if ed.Has(eventName, handler) {
		return ErrHandlerAlreadyRegistered
	}
	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

func (ed *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	if handlers, ok := ed.handlers[eventName]; ok {
		for i, h := range handlers {
			if h == handler {
				ed.handlers[eventName] = append(handlers[:i], handlers[i+1:]...)
				return nil
			}
		}
	}
	return nil
}

func (ed *EventDispatcher) Clear() {
	ed.handlers = make(map[string][]EventHandlerInterface)
}
