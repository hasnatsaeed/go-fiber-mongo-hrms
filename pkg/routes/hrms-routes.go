package routes

import (
	"github.com/gofiber/fiber/v2"
	"github/hasnatsaeed/go-fiber-mongo-hrms/pkg/handlers"
)

func RegisterRoutes(app *fiber.App, handler *handlers.Handler) {

	app.Post("/api/v1/employee", handler.CreateEmployee)
	app.Get("/api/v1/employee", handler.GetAllEmployees)
	app.Get("/api/v1/employee/:id", handler.GetEmployeeById)
	app.Put("/api/v1/employee/:id", handler.UpdateEmployee)
	app.Delete("/api/v1/employee/:id", handler.DeleteEmployee)

}
