// order-service/repository/repository.go

package repository

import (
	orderPB "github.com/SleepingNext/order-service/proto"
)

type Repository interface {
	Index() ([]*orderPB.Order, error)
	Show(*orderPB.Order) (*orderPB.Order, error)
	Store(*orderPB.Order) (*orderPB.Order, error)
	Update(*orderPB.Order) (*orderPB.Order, error)
	Destroy(*orderPB.Order) (*orderPB.Order, error)
}
