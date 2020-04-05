// order-service/repository/postgres/repository.go

package postgres

import (
	"fmt"

	orderPB "github.com/SleepingNext/order-service/proto"
)

// Store will store a new order
func (repo *Repository) Store(order *orderPB.Order) (*orderPB.Order, error) {
	query := fmt.Sprintf("INSERT INTO orders (product_id, user_id, status)"+
		" VALUES (%d, %d, 'active')", order.ProductId, order.UserId)
	_, err := repo.DB.Exec(query)

	return order, err
}

// Update will update an existing order's data
func (repo *Repository) Update(order *orderPB.Order) (*orderPB.Order, error) {
	query := fmt.Sprintf("UPDATE orders SET product_id = %d, user_id = %d, status = 'active'"+
		" WHERE id = %d", order.ProductId, order.UserId, order.Id)
	_, err := repo.DB.Exec(query)

	return order, err
}

// Destroy will update an existing order's status to inactive
func (repo *Repository) Destroy(order *orderPB.Order) (*orderPB.Order, error) {
	query := fmt.Sprintf("UPDATE orders SET status = 'inactive' WHERE id = %d", order.Id)
	_, err := repo.DB.Exec(query)

	return order, err
}
