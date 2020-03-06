// order-service/repository/repository.go

package repository

import (
	pb "github.com/SleepingNext/order-service/proto"
)

type Repository interface {
	Index() ([]*pb.Order, error)
	Show(*pb.Order) (*pb.Order, error)
	Store(*pb.Order) (*pb.Order, error)
	Update(*pb.Order) (*pb.Order, error)
	Destroy(*pb.Order) (*pb.Order, error)
}
