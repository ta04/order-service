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
	"fmt"

	orderPB "github.com/ta04/order-service/proto"
)

// Store stores a new order
func (repo *Postgres) Store(order *orderPB.Order) (*orderPB.Order, error) {
	query := fmt.Sprintf("INSERT INTO orders (product_id, user_id, status)"+
		" VALUES (%d, %d, 'waiting for payment')", order.ProductId, order.UserId)
	_, err := repo.DB.Exec(query)

	return order, err
}

// Update updates an order
func (repo *Postgres) Update(order *orderPB.Order) (*orderPB.Order, error) {
	query := fmt.Sprintf("UPDATE orders SET product_id = %d, user_id = %d, status = '%s'"+
		" WHERE id = %d", order.ProductId, order.UserId, order.Status, order.Id)
	_, err := repo.DB.Exec(query)

	return order, err
}
