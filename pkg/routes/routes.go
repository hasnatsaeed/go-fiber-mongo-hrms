package routes

import (
	"github.com/gofiber/fiber/v2"
	"github/hasnatsaeed/go-fiber-mongo-hrms/pkg/controllers"
)

func RegisterRoutes(app *fiber.App) {

	app.Post("/api/v1/employee", controllers.CreateEmployee)
	app.Get("/api/v1/employee", controllers.GetEmployees)
	app.Get("/api/v1/employee/:id", controllers.GetEmployee)
	app.Put("/api/v1/employee/:id", controllers.UpdateEmployee)
	app.Delete("/api/v1/employee/:id", controllers.DeleteEmployee)

}
