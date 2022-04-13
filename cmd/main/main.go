package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github/hasnatsaeed/go-fiber-mongo-hrms/pkg/configs/storage"
	"github/hasnatsaeed/go-fiber-mongo-hrms/pkg/handlers"
	"github/hasnatsaeed/go-fiber-mongo-hrms/pkg/repositories"
	"github/hasnatsaeed/go-fiber-mongo-hrms/pkg/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	application := fiber.New()

	//Connect mongodb client.
	db := connectMongoDb()

	// Create mongo db repository
	repository := repositories.NewEmployeeMongoDbRepository(db.Collection("employees"))

	// Create api request handler
	handler := handlers.NewEmployeeHandler(repository)

	//Register API routes.
	routes.RegisterRoutes(application, handler)

	//Register shutdown hook to gracefully shut down the application.
	registerShutdownHook(application)

	//Disconnect mongodb client on `main` exit.
	defer func() {
		err := storage.MongoDbDisconnect()
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Println(application.Listen(":9010"))

}

func connectMongoDb() *mongo.Database {
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	db, err := storage.MongoDbConnect(dbHost, dbPort, dbUser, dbPass, dbName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func registerShutdownHook(application *fiber.App) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		_ = <-c
		fmt.Println("Application gracefully shutting down...")
		_ = application.Shutdown()
	}()
}
