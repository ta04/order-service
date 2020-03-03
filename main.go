// product-service/main.go

package main

import (
	"fmt"
	"log"

	"github.com/SleepingNext/product-service/datastore"
	"github.com/SleepingNext/product-service/handler"
	pb "github.com/SleepingNext/product-service/proto"
	"github.com/SleepingNext/product-service/repository"
	_ "github.com/lib/pq"
	"github.com/micro/go-micro"
)

func main() {
	// Setup the micro instance
	s := micro.NewService(
		micro.Name("com.ta04.srv.order"),
	)

	s.Init()

	// Connect to Postgres
	db, err := datastore.ConnectPostgres()
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}

	// Initialize the handler
	h := &handler.Handler{
		Repo: &repository.PostgresRepository{DB: db},
	}

	// Register the handler
	pb.RegisterProductServiceHandler(s.Server(), h)

	// Run the server
	err = s.Run()
	if err != nil {
		fmt.Println(err)
	}
}
