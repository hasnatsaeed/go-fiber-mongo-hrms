package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github/hasnatsaeed/go-fiber-mongo-hrms/pkg/configs/database"
	"github/hasnatsaeed/go-fiber-mongo-hrms/pkg/routes"
	"log"
	"os"
	"os/signal"
)

func main() {

	application := fiber.New()

	//Register shutdown hook to gracefully shut down the application.
	registerShutdownHook(application)

	//Register API routes.
	routes.RegisterRoutes(application)

	//Connect mongodb client.
	database.Connect()

	//Disconnect mongodb client on `main` exit.
	defer database.Disconnect()

	log.Panic(application.Listen(":9010"))

}

func registerShutdownHook(application *fiber.App) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("Application gracefully shutting down...")
		_ = application.Shutdown()
	}()
}
