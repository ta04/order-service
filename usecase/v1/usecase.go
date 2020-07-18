package v1

import (
	"net/http"

	proto "github.com/ta04/order-service/model/proto"
	"github.com/ta04/order-service/repository"
)

// usecase is the struct of order usecase
type Usecase struct {
	Repository repository.Repository
}

// NewUsecase will create a new order usecase
func NewUsecase(repository repository.Repository) *Usecase {
	return &Usecase{
		Repository: repository,
	}
}

func (usecase *Usecase) GetAll(request *proto.GetAllOrdersRequest) (*[]*proto.Order, *proto.Error) {
	if request == nil {
		return nil, &proto.Error{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}
	}

	if request.Status == "" {
		request.Status = "waiting for payment"
	}

	var orders *[]*proto.Order
	var err error
	if request.UserId != 0 {
		orders, err = usecase.Repository.GetAllByUserID(request)
		
	} else {
		orders, err = usecase.Repository.GetAll(request)
	}	
	if orders == nil || len(*orders) <= 0 {
		return nil, &proto.Error{
			Code:    http.StatusInternalServerError,
			Message: "no orders found",
		}
	}
	if err != nil {
		return nil, &proto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}

	return orders, nil
}

func (usecase *Usecase) GetOne(request *proto.GetOneOrderRequest) (*proto.Order, *proto.Error) {
	if request == nil || (request.Id == 0) {
		return nil, &proto.Error{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}
	}

	order, err := usecase.Repository.GetOne(request)
	if err != nil {
		return nil, &proto.Error{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	return order, nil
}

func (usecase *Usecase) CreateOne(order *proto.Order) (*proto.Order, *proto.Error) {
	if order == nil {
		return nil, &proto.Error{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}
	}

	order, err := usecase.Repository.CreateOne(order)
	if err != nil {
		return nil, &proto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return order, nil
}

func (usecase *Usecase) UpdateOne(order *proto.Order) (*proto.Order, *proto.Error) {
	if order == nil {
		return nil, &proto.Error{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}
	}

	order, err := usecase.Repository.UpdateOne(order)
	if err != nil {
		return nil, &proto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return order, nil
}
