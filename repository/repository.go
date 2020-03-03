package repository

import (
	"database/sql"
	"fmt"
	pb "github.com/SleepingNext/order-service/proto"
)

type Repository interface {
	Index() ([]*pb.Order, error)
	Show(*pb.Order) (*pb.Order, error)
	Store(*pb.Order) (*pb.Order, error)
	Update(*pb.Order) (*pb.Order, error)
	Destroy(*pb.Order) (*pb.Order, error)
}

type PostgresRepository struct {
	DB *sql.DB
}

func (repo *PostgresRepository) Index() (orders []*pb.Order, err error) {
	var id, productId, userId int32
	var status bool

	query := "SELECT * FROM orders"
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&id, &productId, &userId, &status)
		if err != nil {
			return nil, err
		}
		order := &pb.Order{
			Id:        id,
			ProductId: productId,
			UserId:    userId,
			Status:    status,
		}
		orders = append(orders, order)
	}

	return orders, err
}

func (repo *PostgresRepository) Show(order *pb.Order) (*pb.Order, error) {
	var id, productId, userId int32
	var status bool

	query := fmt.Sprintf("SELECT * FROM orders WHERE id = %d", order.Id)
	err := repo.DB.QueryRow(query).Scan(&id, &productId, &userId, &status)
	if err != nil {
		return nil, err
	}

	return &pb.Order{
		Id:        id,
		ProductId: productId,
		UserId:    userId,
		Status:    status,
	}, err
}

func (repo *PostgresRepository) Store(order *pb.Order) (*pb.Order, error) {
	query := fmt.Sprintf("INSERT INTO orders (product_id, user_id, status)"+
		" VALUES (%d, %d, %t)", order.ProductId, order.UserId, order.Status)
	_, err := repo.DB.Exec(query)

	return order, err
}

func (repo *PostgresRepository) Update(order *pb.Order) (*pb.Order, error) {
	query := fmt.Sprintf("UPDATE orders SET product_id = %d, user_id = %d, status = %t"+
		" WHERE id = %d", order.ProductId, order.UserId, order.Status, order.Id)
	_, err := repo.DB.Exec(query)

	return order, err
}

func (repo *PostgresRepository) Destroy(order *pb.Order) (*pb.Order, error) {
	query := fmt.Sprintf("DELETE FROM orders WHERE id = %d", order.Id)
	_, err := repo.DB.Exec(query)

	return order, err

}
