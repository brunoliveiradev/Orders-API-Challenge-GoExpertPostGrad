package database

import (
	"GoExpertPostGrad-Orders-Challenge/internal/entity"
	"database/sql"
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
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotalCount() (int, error) {
	var count int
	err := r.Db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
