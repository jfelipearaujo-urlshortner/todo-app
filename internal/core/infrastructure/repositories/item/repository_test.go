package item_repository_test

import (
	"testing"

	"github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/entities"
	item_repository "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/infrastructure/repositories/item"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	t.Run("Should create an item successfully", func(t *testing.T) {
		// Arrange
		sut := item_repository.New()

		item := entities.New("Buy milk", "Milk is the best drink in the world", nil)

		// Act
		err := sut.Create(item)

		// Assert
		assert.NoError(t, err)

		exists, err := sut.Get(item.ID)
		assert.NoError(t, err)
		assert.NotNil(t, exists)
		assert.Equal(t, item, *exists)
	})

	t.Run("Should return an error if the item id already exists", func(t *testing.T) {
		// Arrange
		sut := item_repository.New()

		item := entities.New("Buy milk", "Milk is the best drink in the world", nil)

		err := sut.Create(item)
		assert.NoError(t, err)

		// Act
		err = sut.Create(item)

		// Assert
		assert.Error(t, err)
	})
}

func TestGet(t *testing.T) {
	t.Run("Should return an item successfully", func(t *testing.T) {
		// Arrange
		sut := item_repository.New()

		item := entities.New("Buy milk", "Milk is the best drink in the world", nil)

		err := sut.Create(item)
		assert.NoError(t, err)

		// Act
		result, err := sut.Get(item.ID)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, item, *result)
	})

	t.Run("Should return nil if the item does not exist", func(t *testing.T) {
		// Arrange
		sut := item_repository.New()

		item := entities.New("Buy milk", "Milk is the best drink in the world", nil)

		// Act
		result, err := sut.Get(item.ID)

		// Assert
		assert.NoError(t, err)
		assert.Nil(t, result)
	})
}
