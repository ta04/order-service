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
	orderPB "github.com/ta04/order-service/proto"
)

// Repository is the interface of repositories.
// As there are number of repositories can be used.
type Repository interface {
	Index(*orderPB.IndexOrdersRequest) ([]*orderPB.Order, error)
	IndexByUserID(*orderPB.User) ([]*orderPB.Order, error)
	Show(*orderPB.Order) (*orderPB.Order, error)
	Store(*orderPB.Order) (*orderPB.Order, error)
	Update(*orderPB.Order) (*orderPB.Order, error)
}
