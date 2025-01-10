package get_controller

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	get_contract "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/usecases/item/get"
)

type controller struct {
	validator *validator.Validate
	useCase   get_contract.UseCase
}

func New(validator *validator.Validate, useCase get_contract.UseCase) *controller {
	return &controller{
		validator: validator,
		useCase:   useCase,
	}
}

func (c *controller) Handle(ctx *fiber.Ctx) error {
	var request Request

	request.ID = ctx.Params("id")

	if err := c.validator.Struct(&request); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(err)
	}

	item, err := c.useCase.Execute(request.ID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}

	if item == nil {
		return ctx.SendStatus(http.StatusNotFound)
	}

	return ctx.JSON(item)
}
