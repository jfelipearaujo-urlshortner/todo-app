package create_controller

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	create_contract "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/usecases/item/create"
)

const datetimeLayout = "2006-01-02T15:04:05.000Z"

type controller struct {
	validator *validator.Validate
	useCase   create_contract.UseCase
}

func New(validator *validator.Validate, useCase create_contract.UseCase) *controller {
	return &controller{
		validator: validator,
		useCase:   useCase,
	}
}

func (c *controller) Handle(ctx *fiber.Ctx) error {
	var request Request

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	if err := c.validator.Struct(&request); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(err)
	}

	deadlineAt, err := ParseDeadlineAt(request.DeadlineAt)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(err)
	}

	item, err := c.useCase.Execute(request.Title, request.Description, deadlineAt)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}

	return ctx.Status(http.StatusCreated).JSON(Map(item))
}
