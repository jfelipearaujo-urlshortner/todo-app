package complete_usecase_test

import (
	"testing"

	"github.com/google/uuid"
	complete_usecase "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/application/usecases/item/complete"
	"github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/entities"
	item_repository "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/infrastructure/repositories/item"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestExecute(t *testing.T) {
	t.Run("Should create an item successfully", func(t *testing.T) {
		// Arrange
		repository := item_repository.NewMockRepository(t)

		id := uuid.NewString()

		repository.EXPECT().
			Get(id).
			Return(&entities.Item{
				ID:   id,
				Done: false,
			}, nil).
			Once()

		repository.EXPECT().
			Update(mock.Anything).
			Return(nil).
			Once()

		sut := complete_usecase.New(repository)

		// Act
		err := sut.Execute(id)

		// Assert
		assert.NoError(t, err)
		repository.AssertExpectations(t)
	})

	t.Run("Should return an error while try to get an item by ID", func(t *testing.T) {
		// Arrange
		repository := item_repository.NewMockRepository(t)

		id := uuid.NewString()

		repository.EXPECT().
			Get(id).
			Return(nil, assert.AnError).
			Once()

		sut := complete_usecase.New(repository)

		// Act
		err := sut.Execute(id)

		// Assert
		assert.Error(t, err)
		repository.AssertExpectations(t)
	})

	t.Run("Should return an error while try to complete an item that does not exist", func(t *testing.T) {
		// Arrange
		repository := item_repository.NewMockRepository(t)

		id := uuid.NewString()

		repository.EXPECT().
			Get(id).
			Return(nil, nil).
			Once()

		sut := complete_usecase.New(repository)

		// Act
		err := sut.Execute(id)

		// Assert
		assert.Error(t, err)
		repository.AssertExpectations(t)
	})

	t.Run("Should return an error when try to complete an item that is already completed", func(t *testing.T) {
		// Arrange
		repository := item_repository.NewMockRepository(t)

		id := uuid.NewString()

		repository.EXPECT().
			Get(id).
			Return(&entities.Item{
				ID:   id,
				Done: true,
			}, nil).
			Once()

		sut := complete_usecase.New(repository)

		// Act
		err := sut.Execute(id)

		// Assert
		assert.Error(t, err)
		repository.AssertExpectations(t)
	})

	t.Run("Should return an error when try to update an item", func(t *testing.T) {
		// Arrange
		repository := item_repository.NewMockRepository(t)

		id := uuid.NewString()

		repository.EXPECT().
			Get(id).
			Return(&entities.Item{
				ID:   id,
				Done: false,
			}, nil).
			Once()

		repository.EXPECT().
			Update(mock.Anything).
			Return(assert.AnError).
			Once()

		sut := complete_usecase.New(repository)

		// Act
		err := sut.Execute(id)

		// Assert
		assert.Error(t, err)
		repository.AssertExpectations(t)
	})
}
