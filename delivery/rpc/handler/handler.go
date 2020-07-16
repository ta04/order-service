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

package handler

import (
	"context"
	"errors"

	proto "github.com/ta04/order-service/model/proto"
	"github.com/ta04/order-service/usecase"
)

// Handler is the handler of order service
type Handler struct {
	Usecase usecase.Usecase
}

// NewHandler will create a new order handler
func NewHandler(usecase usecase.Usecase) *Handler {
	return &Handler{
		Usecase: usecase,
	}
}

// GetAllOrders will get all orders
func (handler *Handler) GetAllOrders(ctx context.Context, req *proto.GetAllOrdersRequest, res *proto.Response) error {
	orders, err := handler.Usecase.GetAll(req)
	if err != nil {
		res.Orders = nil
		res.Error = err

		return errors.New(err.Message)
	}

	res.Orders = orders
	res.Error = nil

	return nil
}

// GetOneOrder will get an order
func (handler *Handler) GetOneOrder(ctx context.Context, req *proto.GetOneOrderRequest, res *proto.Response) error {
	order, err := handler.Usecase.GetOne(req)
	if err != nil {
		res.Order = nil
		res.Error = err

		return errors.New(err.Message)
	}

	res.Order = order
	res.Error = nil

	return nil
}

// CreateOneOrder will create a new order
func (handler *Handler) CreateOneOrder(ctx context.Context, req *proto.Order, res *proto.Response) error {
	order, err := handler.Usecase.CreateOne(req)
	if err != nil {
		res.Order = nil
		res.Error = err

		return errors.New(err.Message)
	}

	res.Order = order
	res.Error = nil

	return nil
}

// UpdateOneOrder will update an order
func (handler *Handler) UpdateOneOrder(ctx context.Context, req *proto.Order, res *proto.Response) error {
	order, err := handler.Usecase.UpdateOne(req)
	if err != nil {
		res.Order = nil
		res.Error = err

		return errors.New(err.Message)
	}

	res.Order = order
	res.Error = nil

	return nil
}
