package list_controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	list_contract "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/usecases/item/list"
)

type controller struct {
	validator *validator.Validate
	useCase   list_contract.UseCase
}

func New(validator *validator.Validate, useCase list_contract.UseCase) *controller {
	return &controller{
		validator: validator,
		useCase:   useCase,
	}
}

func (c *controller) Handle(ctx *fiber.Ctx) error {
	items := c.useCase.Execute()

	return ctx.JSON(items)
}
