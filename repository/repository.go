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

package repository

import (
	proto "github.com/ta04/order-service/model/proto"
)

// Repository is the interface of repositories.
// As there are number of repositories can be used.
type Repository interface {
	GetAllByUserID(request *proto.GetAllOrdersRequest) (*[]*proto.Order, error)
	GetAll(request *proto.GetAllOrdersRequest) (*[]*proto.Order, error)
	GetOne(request *proto.GetOneOrderRequest) (*proto.Order, error)
	CreateOne(order *proto.Order) (*proto.Order, error)
	UpdateOne(order *proto.Order) (*proto.Order, error)
}
