package handler

import (
	"context"

	pb "github.com/SleepingNext/order-service/proto"
	"github.com/SleepingNext/order-service/repository"
)

type Handler struct {
	Repo repository.Repository
}

func (h *Handler) IndexOrders(ctx context.Context, req *pb.IndexOrdersRequest, res *pb.Response) error {
	orders, err := h.Repo.Index()
	if err != nil {
		return err
	}

	res.Orders = orders
	res.Error = nil

	return err
}

func (h *Handler) ShowOrder(ctx context.Context, req *pb.Order, res *pb.Response) error {
	order, err := h.Repo.Show(req)
	if err != nil {
		return err
	}

	res.Order = order
	res.Error = nil

	return nil
}

func (h *Handler) StoreOrder(ctx context.Context, req *pb.Order, res *pb.Response) error {
	order, err := h.Repo.Store(req)
	if err != nil {
		return err
	}

	res.Order = order
	res.Error = nil

	return err
}

func (h *Handler) UpdateOrder(ctx context.Context, req *pb.Order, res *pb.Response) error {
	order, err := h.Repo.Update(req)
	if err != nil {
		return err
	}

	res.Order = order
	res.Error = nil

	return nil
}

func (h *Handler) DestroyOrder(ctx context.Context, req *pb.Order, res *pb.Response) error {
	order, err := h.Repo.Destroy(req)
	if err != nil {
		return err
	}

	res.Order = order
	res.Error = nil

	return err
}
