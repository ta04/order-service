// order-service/repository/postgres/repository.go

package postgres

import (
	"fmt"

	orderPB "github.com/SleepingNext/order-service/proto"
)

func (repo *Repository) Store(order *orderPB.Order) (*orderPB.Order, error) {
	query := fmt.Sprintf("INSERT INTO orders (product_id, user_id, status)"+
		" VALUES (%d, %d, '%s')", order.ProductId, order.UserId, order.Status)
	_, err := repo.DB.Exec(query)

	return order, err
}

func (repo *Repository) Update(order *orderPB.Order) (*orderPB.Order, error) {
	query := fmt.Sprintf("UPDATE orders SET product_id = %d, user_id = %d, status = '%s'"+
		" WHERE id = %d", order.ProductId, order.UserId, order.Status, order.Id)
	_, err := repo.DB.Exec(query)

	return order, err
}

func (repo *Repository) Destroy(order *orderPB.Order) (*orderPB.Order, error) {
	query := fmt.Sprintf("DELETE FROM orders WHERE id = %d", order.Id)
	_, err := repo.DB.Exec(query)

	return order, err
}
