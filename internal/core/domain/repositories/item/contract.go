package item_contract

import "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/entities"

type Repository interface {
	Create(item entities.Item) error
	Get(id string) (*entities.Item, error)
	Update(item *entities.Item) error
	Delete(id string) error
	List() []entities.Item
}
