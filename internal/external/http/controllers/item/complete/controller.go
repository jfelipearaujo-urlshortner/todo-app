package complete_controller

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	complete_contract "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/usecases/item/complete"
)

type Controller struct {
	validator *validator.Validate
	useCase   complete_contract.UseCase
}

func New(validator *validator.Validate, useCase complete_contract.UseCase) *Controller {
	return &Controller{
		validator: validator,
		useCase:   useCase,
	}
}

func (c *Controller) Handle(ctx *fiber.Ctx) error {
	var request Request

	request.ID = ctx.Params("id")

	if err := c.validator.Struct(&request); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(err)
	}

	if err := c.useCase.Execute(request.ID); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}

	return ctx.SendStatus(http.StatusNoContent)
}
