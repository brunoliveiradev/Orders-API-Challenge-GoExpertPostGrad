package database

import (
	"GoExpertPostGrad-Orders-Challenge/internal/entity"
	"database/sql"
)

const (
	insertQuery    = "INSERT INTO orders (id, name, price, tax, final_price) VALUES (?, ?, ?, ?, ?)"
	selectAllQuery = "SELECT id, name, price, tax, final_price FROM orders"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	// When we use the Prepare method, we need to close the statement after using it
	// The advantage of using Prepare is that it is possible to reuse the statement and avoid SQL injection
	stmt, err := r.Db.Prepare(insertQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(order.ID, order.Name, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) ListAllOrders() ([]*entity.Order, error) {
	rows, err := r.Db.Query(selectAllQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*entity.Order
	for rows.Next() {
		order := new(entity.Order)
		if err := rows.Scan(&order.ID, &order.Name, &order.Price, &order.Tax, &order.FinalPrice); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
