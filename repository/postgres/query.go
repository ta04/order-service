/*
Dear Programmers,

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
*                                                 *
*	This file belongs to Kevin Veros Hamonangan   *
*	and	Fandi Fladimir Dachi and is a part of     *
*	our	last project as the student of Del        *
*	Institute of Technology, Sitoluama.           *
*	Please contact us via Instagram:              *
*	sleepingnext and fandi_dachi                  *
*	before copying this file.                     *
*	Thank you, buddy. 😊                          *
*                                                 *
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
*/

package postgres

import (
	"database/sql"
	"fmt"

	orderPB "github.com/ta04/order-service/proto"
)

// Postgres is the implementor of Postgres interface
type Postgres struct {
	DB *sql.DB
}

// Index indexes all orders
func (repo *Postgres) Index(req *orderPB.IndexOrdersRequest) (orders []*orderPB.Order, err error) {
	var id, productID, userID int32
	var status string

	query := "SELECT * FROM orders"
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

// IndexByUserID indexes all orders by userID
func (repo *Postgres) IndexByUserID(user *orderPB.User) (orders []*orderPB.Order, err error) {
	var id, productID, userID int32
	var status string

	query := fmt.Sprintf("SELECT * FROM orders WHERE user_id = %d", user.Id)
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

// Show shows an order by id
func (repo *Postgres) Show(order *orderPB.Order) (*orderPB.Order, error) {
	var id, productID, userID int32
	var status string

	query := fmt.Sprintf("SELECT * FROM orders WHERE id = %d", order.Id)
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
