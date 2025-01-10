package update_controller

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/constants/layouts"
	update_contract "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/usecases/item/update"
	"github.com/jfelipearaujo-urlshortner/todo-app/internal/external/shared"
)

type controller struct {
	validator *validator.Validate
	useCase   update_contract.UseCase
}

func New(validator *validator.Validate, useCase update_contract.UseCase) *controller {
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

	request.ID = ctx.Params("id")

	if err := c.validator.Struct(&request); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(err)
	}

	deadlineAt, err := shared.ParseStringToTime(layouts.DatetimeLayout, request.DeadlineAt)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(err)
	}

	item, err := c.useCase.Execute(request.ID, request.Title, request.Description, deadlineAt)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}

	return ctx.Status(http.StatusOK).JSON(Map(item))
}
