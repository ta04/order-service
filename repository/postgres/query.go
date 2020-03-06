// order-service/repository/postgres/query.go

package postgres

import (
	"database/sql"
	"fmt"

	orderPB "github.com/SleepingNext/order-service/proto"
)

type Repository struct {
	DB *sql.DB
}

func (repo *Repository) Index() (orders []*orderPB.Order, err error) {
	var id, productId, userId int32
	var status bool

	query := "SELECT * FROM orders"
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&id, &productId, &userId, &status)
		if err != nil {
			return nil, err
		}
		order := &orderPB.Order{
			Id:        id,
			ProductId: productId,
			UserId:    userId,
			Status:    status,
		}
		orders = append(orders, order)
	}

	return orders, err
}

func (repo *Repository) Show(order *orderPB.Order) (*orderPB.Order, error) {
	var id, productId, userId int32
	var status bool

	query := fmt.Sprintf("SELECT * FROM orders WHERE id = %d", order.Id)
	err := repo.DB.QueryRow(query).Scan(&id, &productId, &userId, &status)
	if err != nil {
		return nil, err
	}

	return &orderPB.Order{
		Id:        id,
		ProductId: productId,
		UserId:    userId,
		Status:    status,
	}, err
}
