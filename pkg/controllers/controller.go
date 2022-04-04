package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github/hasnatsaeed/go-fiber-mongo-hrms/pkg/models"
)

func CreateEmployee(ctx *fiber.Ctx) error {
	employee := models.Employee{}

	errParsing := ctx.BodyParser(&employee)
	if errParsing != nil {
		panic(errParsing)
	}
	createdEmployee, errCreation := employee.CreateEmployee()
	if errCreation != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(errCreation.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(createdEmployee)

}

func GetEmployees(ctx *fiber.Ctx) error {
	employees, err := models.GetEmployees()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(employees)
}

func GetEmployee(ctx *fiber.Ctx) error {

	employee, err := models.GetEmployee(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(employee)
}

func UpdateEmployee(ctx *fiber.Ctx) error {
	employee := models.Employee{}
	errParsing := ctx.BodyParser(&employee)
	if errParsing != nil {
		panic(errParsing)
	}
	updateCount, errUpdate := models.UpdateEmployee(ctx.Params("id"), employee)

	if errUpdate != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(errUpdate.Error())
	}

	if updateCount < 1 {
		return ctx.Status(fiber.StatusNotFound).SendString("record not found")
	}

	return ctx.Status(fiber.StatusOK).JSON("Updated")
}

func DeleteEmployee(ctx *fiber.Ctx) error {
	deleteCount, err := models.DeleteEmployee(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if deleteCount < 1 {
		return ctx.Status(fiber.StatusNotFound).SendString("record not found")
	}

	return ctx.Status(fiber.StatusOK).JSON("Deleted")
}
