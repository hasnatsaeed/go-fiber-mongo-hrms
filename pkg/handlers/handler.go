package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github/hasnatsaeed/go-fiber-mongo-hrms/pkg/models"
	"github/hasnatsaeed/go-fiber-mongo-hrms/pkg/repositories"
)

func NewEmployeeHandler(repository repositories.Repository) *Handler {
	return &Handler{
		repository: repository,
	}
}

type Handler struct {
	repository repositories.Repository
}

func (handler *Handler) CreateEmployee(ctx *fiber.Ctx) error {
	employeeToCreate := &models.Employee{}

	errParsingBody := ctx.BodyParser(employeeToCreate)
	if errParsingBody != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": errParsingBody.Error()})
	}
	createdEmployee, errCreateEmployee := handler.repository.SaveEmployee(ctx.Context(), employeeToCreate)
	if errCreateEmployee != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": errParsingBody.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(createdEmployee)

}

func (handler *Handler) GetAllEmployees(ctx *fiber.Ctx) error {
	employees, errGetEmployees := handler.repository.FetchAllEmployees(ctx.Context())
	if errGetEmployees != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": errGetEmployees.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(employees)
}

func (handler *Handler) GetEmployeeById(ctx *fiber.Ctx) error {

	employee, errGetEmployee := handler.repository.FetchEmployeeById(ctx.Context(), ctx.Params("id"))
	if errGetEmployee != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": errGetEmployee.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(employee)
}

func (handler *Handler) UpdateEmployee(ctx *fiber.Ctx) error {

	employeeToUpdate := &models.Employee{}
	errParsingBody := ctx.BodyParser(employeeToUpdate)
	if errParsingBody != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": errParsingBody.Error()})
	}
	updatedEmployee, errUpdateEmployee := handler.repository.UpdateEmployee(ctx.Context(), ctx.Params("id"), employeeToUpdate)

	if errUpdateEmployee != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": errUpdateEmployee.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(updatedEmployee)
}

func (handler *Handler) DeleteEmployee(ctx *fiber.Ctx) error {
	deletedEmployee, errDeleteEmployee := handler.repository.DeleteEmployee(ctx.Context(), ctx.Params("id"))
	if errDeleteEmployee != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": errDeleteEmployee.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(deletedEmployee)
}
