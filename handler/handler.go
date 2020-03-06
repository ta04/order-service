// order-service/handler/handler.go

package handler

import (
	"context"

	orderPB "github.com/SleepingNext/order-service/proto"
	orderRepo "github.com/SleepingNext/order-service/repository"
)

type handler struct {
	repository orderRepo.Repository
}

func NewHandler(repo orderRepo.Repository) *handler {
	return &handler{
		repository: repo,
	}
}

func (h *handler) IndexOrders(ctx context.Context, req *orderPB.IndexOrdersRequest, res *orderPB.Response) error {
	orders, err := h.repository.Index()
	if err != nil {
		return err
	}

	res.Orders = orders
	res.Error = nil

	return err
}

func (h *handler) ShowOrder(ctx context.Context, req *orderPB.Order, res *orderPB.Response) error {
	order, err := h.repository.Show(req)
	if err != nil {
		return err
	}

	res.Order = order
	res.Error = nil

	return nil
}

func (h *handler) StoreOrder(ctx context.Context, req *orderPB.Order, res *orderPB.Response) error {
	order, err := h.repository.Store(req)
	if err != nil {
		return err
	}

	res.Order = order
	res.Error = nil

	return err
}

func (h *handler) UpdateOrder(ctx context.Context, req *orderPB.Order, res *orderPB.Response) error {
	order, err := h.repository.Update(req)
	if err != nil {
		return err
	}

	res.Order = order
	res.Error = nil

	return nil
}

func (h *handler) DestroyOrder(ctx context.Context, req *orderPB.Order, res *orderPB.Response) error {
	order, err := h.repository.Destroy(req)
	if err != nil {
		return err
	}

	res.Order = order
	res.Error = nil

	return err
}
