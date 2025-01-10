package complete_controller_test

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	complete_controller "github.com/jfelipearaujo-urlshortner/todo-app/internal/external/http/controllers/item/complete"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Run("Should not return an error if Request is valid", func(t *testing.T) {
		// Arrange
		validator := validator.New()

		sut := complete_controller.Request{
			ID: uuid.NewString(),
		}

		// Act
		err := validator.Struct(&sut)

		// Assert
		assert.NoError(t, err)
	})

	t.Run("Should return an error if Request does not have an ID", func(t *testing.T) {
		// Arrange
		validator := validator.New()

		sut := complete_controller.Request{
			ID: "",
		}

		// Act
		err := validator.Struct(&sut)

		// Assert
		assert.Error(t, err)
	})

	t.Run("Should return an error if Request does not have a valid ID", func(t *testing.T) {
		// Arrange
		validator := validator.New()

		sut := complete_controller.Request{
			ID: "not-uuid",
		}

		// Act
		err := validator.Struct(&sut)

		// Assert
		assert.Error(t, err)
	})
}
