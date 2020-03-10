// order-service/main.go

package main

import (
	"github.com/SleepingNext/order-service/repository/postgres"
	"log"

	"github.com/SleepingNext/order-service/database"
	"github.com/SleepingNext/order-service/handler"
	orderPB "github.com/SleepingNext/order-service/proto"
	_ "github.com/lib/pq"
	"github.com/micro/go-micro"
)

func main() {
	// Create a new service
	s := micro.NewService(
		micro.Name("com.ta04.srv.order"),
	)

	// Initialize the service
	s.Init()

	// Connect to Postgres
	db, err := database.OpenPostgresConnection()
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Create a new handler
	h := handler.NewHandler(&postgres.Repository{
		DB: db,
	})

	// Register the handler
	orderPB.RegisterOrderServiceHandler(s.Server(), h)

	// Run the server
	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
