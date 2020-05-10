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

	orderPB "github.com/ta04/order-service/proto"
	orderRepo "github.com/ta04/order-service/repository"
)

// Handler is the handler of order service
type Handler struct {
	repository orderRepo.Repository
}

// NewHandler creates a new order service handler
func NewHandler(repo orderRepo.Repository) *Handler {
	return &Handler{
		repository: repo,
	}
}

// IndexOrders indexes the orders
func (h *Handler) IndexOrders(ctx context.Context, req *orderPB.IndexOrdersRequest, res *orderPB.Response) error {
	orders, err := h.repository.Index(req)
	if err != nil {
		return err
	}

	res.Orders = orders
	res.Error = nil

	return err
}

// IndexOrdersByUserID indexes the orders by user ID
func (h *Handler) IndexOrdersByUserID(ctx context.Context, req *orderPB.User, res *orderPB.Response) error {
	orders, err := h.repository.IndexByUserID(req)
	if err != nil {
		return err
	}

	res.Orders = orders
	res.Error = nil

	return err
}

// ShowOrder shows an order by ID
func (h *Handler) ShowOrder(ctx context.Context, req *orderPB.Order, res *orderPB.Response) error {
	order, err := h.repository.Show(req)
	if err != nil {
		return err
	}

	res.Order = order
	res.Error = nil

	return nil
}

// StoreOrder stores a new order
func (h *Handler) StoreOrder(ctx context.Context, req *orderPB.Order, res *orderPB.Response) error {
	order, err := h.repository.Store(req)
	if err != nil {
		return err
	}

	res.Order = order
	res.Error = nil

	return err
}

// UpdateOrder updates an order
func (h *Handler) UpdateOrder(ctx context.Context, req *orderPB.Order, res *orderPB.Response) error {
	order, err := h.repository.Update(req)
	if err != nil {
		return err
	}

	res.Order = order
	res.Error = nil

	return nil
}
