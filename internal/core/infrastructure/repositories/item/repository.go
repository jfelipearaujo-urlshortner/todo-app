package item_repository

import (
	"fmt"

	"github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/entities"
	item_contract "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/repositories/item"
)

type repository struct {
	database map[string]entities.Item
}

func New() item_contract.Repository {
	return &repository{}
}

func (r *repository) Create(item entities.Item) error {
	_, exists := r.database[item.ID]
	if exists {
		return fmt.Errorf("item with id %s already exists", item.ID)
	}

	r.database[item.ID] = item

	return nil
}

func (r *repository) Get(id string) (*entities.Item, error) {
	item, exists := r.database[id]
	if !exists {
		return nil, nil
	}

	return &item, nil
}

func (r *repository) Update(item *entities.Item) error {
	_, exists := r.database[item.ID]
	if !exists {
		return fmt.Errorf("item with id %s does not exist", item.ID)
	}

	r.database[item.ID] = *item

	return nil
}

func (r *repository) Delete(id string) error {
	_, exists := r.database[id]
	if !exists {
		return fmt.Errorf("item with id %s does not exist", id)
	}

	delete(r.database, id)

	return nil
}

func (r *repository) List() []entities.Item {
	var items []entities.Item

	for _, value := range r.database {
		items = append(items, value)
	}

	return items
}
