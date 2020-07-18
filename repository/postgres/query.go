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

	proto "github.com/ta04/order-service/model/proto"
)

// Postgres is the implementor of Postgres interface
type Postgres struct {
	DB *sql.DB
}

// NewPostgres will create a new postgres instance
func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{
		DB: db,
	}
}

// GetAllByUserID will get all orders by user id
func (postgres *Postgres) GetAllByUserID(request *proto.GetAllOrdersRequest) (*[]*proto.Order, error) {
	var id, productID, userID int32
	var status string
	var orders []*proto.Order

	query := fmt.Sprintf("SELECT * FROM orders WHERE user_id = %d AND status = '%s'", request.UserId, request.Status)
	rows, err := postgres.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&id, &productID, &userID, &status)
		if err != nil {
			return nil, err
		}
		order := &proto.Order{
			Id:        id,
			ProductId: productID,
			UserId:    userID,
			Status:    status,
		}
		orders = append(orders, order)
	}

	return &orders, err
}

// GetAll will get all orders
func (postgres *Postgres) GetAll(request *proto.GetAllOrdersRequest) (*[]*proto.Order, error) {
	var id, productID, userID int32
	var status string
	var orders []*proto.Order

	query := fmt.Sprintf("SELECT * FROM orders WHERE status = '%s'", request.Status)
	rows, err := postgres.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&id, &productID, &userID, &status)
		if err != nil {
			return nil, err
		}
		order := &proto.Order{
			Id:        id,
			ProductId: productID,
			UserId:    userID,
			Status:    status,
		}
		orders = append(orders, order)
	}

	return &orders, err
}

// GetOne will get an order by id
func (postgres *Postgres) GetOne(request *proto.GetOneOrderRequest) (*proto.Order, error) {
	var id, productID, userID int32
	var status string

	query := fmt.Sprintf("SELECT * FROM orders WHERE id = %d", request.Id)
	err := postgres.DB.QueryRow(query).Scan(&id, &productID, &userID, &status)
	if err != nil {
		return nil, err
	}

	return &proto.Order{
		Id:        id,
		ProductId: productID,
		UserId:    userID,
		Status:    status,
	}, err
}
