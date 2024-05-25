package entity

import (
	"errors"
	"github.com/google/uuid"
)

var ErrInvalidEntity = errors.New("invalid entity")

type Order struct {
	ID         string
	Name       string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(name string, price float64, tax float64) (*Order, error) {
	order := &Order{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
		Tax:   tax,
	}
	err := order.IsValid()
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) IsValid() error {
	if o.Name == "" {
		return ErrInvalidEntity
	}
	if o.Price <= 0 {
		return ErrInvalidEntity
	}
	if o.Tax <= 0 {
		return ErrInvalidEntity
	}
	return nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax

	err := o.IsValid()
	if err != nil {
		return err
	}
	return nil
}
