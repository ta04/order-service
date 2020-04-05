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

// Index will index all active orders
func (repo *Repository) Index() (orders []*orderPB.Order, err error) {
	var id, productID, userID int32
	var status string

	query := "SELECT * FROM orders WHERE status = 'active'"
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&id, &productID, &userID, &status)
		if err != nil {
			return nil, err
		}
		order := &orderPB.Order{
			Id:        id,
			ProductId: productID,
			UserId:    userID,
			Status:    status,
		}
		orders = append(orders, order)
	}

	return orders, err
}

// IndexByUserID will index all active orders by it's userID
func (repo *Repository) IndexByUserID(user *orderPB.User) (orders []*orderPB.Order, err error) {
	var id, productID, userID int32
	var status string

	query := fmt.Sprintf("SELECT * FROM orders WHERE user_id = %d AND status = 'active'", user.Id)
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&id, &productID, &userID, &status)
		if err != nil {
			return nil, err
		}
		order := &orderPB.Order{
			Id:        id,
			ProductId: productID,
			UserId:    userID,
			Status:    status,
		}
		orders = append(orders, order)
	}

	return orders, err
}

// Show will show an active order by it's id
func (repo *Repository) Show(order *orderPB.Order) (*orderPB.Order, error) {
	var id, productID, userID int32
	var status string

	query := fmt.Sprintf("SELECT * FROM orders WHERE id = %d AND status = 'active'", order.Id)
	err := repo.DB.QueryRow(query).Scan(&id, &productID, &userID, &status)
	if err != nil {
		return nil, err
	}

	return &orderPB.Order{
		Id:        id,
		ProductId: productID,
		UserId:    userID,
		Status:    status,
	}, err
}
