package complete_controller_test

import (
	"net/http"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	complete_usecase "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/application/usecases/item/complete"
	complete_controller "github.com/jfelipearaujo-urlshortner/todo-app/internal/external/http/controllers/item/complete"
	"github.com/stretchr/testify/assert"
)

func TestHandle(t *testing.T) {
	t.Run("Should complete an item successfully", func(t *testing.T) {
		// Arrange
		id := uuid.NewString()

		validator := validator.New()
		useCase := complete_usecase.NewMockUseCase(t)

		useCase.EXPECT().
			Execute(id).
			Return(nil).
			Once()

		sut := complete_controller.New(validator, useCase)

		app := fiber.New()
		app.Post("/test/:id", sut.Handle)

		req, err := http.NewRequest(http.MethodPost, "/test/"+id, nil)
		assert.NoError(t, err)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, resp.StatusCode)
		useCase.AssertExpectations(t)
	})

	t.Run("Should return validation error", func(t *testing.T) {
		// Arrange
		validator := validator.New()
		useCase := complete_usecase.NewMockUseCase(t)

		sut := complete_controller.New(validator, useCase)

		app := fiber.New()
		app.Post("/test/:id", sut.Handle)

		req, err := http.NewRequest(http.MethodPost, "/test/not-uuid", nil)
		assert.NoError(t, err)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
		useCase.AssertExpectations(t)
	})

	t.Run("Should return internal server error", func(t *testing.T) {
		// Arrange
		id := uuid.NewString()

		validator := validator.New()
		useCase := complete_usecase.NewMockUseCase(t)

		useCase.EXPECT().
			Execute(id).
			Return(assert.AnError).
			Once()

		sut := complete_controller.New(validator, useCase)

		app := fiber.New()
		app.Post("/test/:id", sut.Handle)

		req, err := http.NewRequest(http.MethodPost, "/test/"+id, nil)
		assert.NoError(t, err)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		useCase.AssertExpectations(t)
	})
}
