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
	"errors"
	"fmt"

	proto "github.com/ta04/order-service/model/proto"
)

// CreateOne will create a new order
func (postgres *Postgres) CreateOne(order *proto.Order) (*proto.Order, error) {
	query := fmt.Sprintf("INSERT INTO orders (product_id, user_id, status)"+
		" VALUES (%d, %d, 'waiting for payment')", order.ProductId, order.UserId)
	_, err := postgres.DB.Exec(query)
	if err != nil {
		return nil, err
	}

	return order, err
}

// UpdateOne will update a product
func (postgres *Postgres) UpdateOne(order *proto.Order) (*proto.Order, error) {
	query := fmt.Sprintf("UPDATE orders SET product_id = %d, user_id = %d, status = '%s'"+
		" WHERE id = %d", order.ProductId, order.UserId, order.Status, order.Id)
	res, err := postgres.DB.Exec(query)
	if err != nil {
		return nil, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if count <= 0 {
		return nil, errors.New("sql: no rows found")
	}

	return order, err
}
