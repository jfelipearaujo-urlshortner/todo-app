package get_usecase

import (
	"fmt"

	"github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/entities"
	item_contract "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/repositories/item"
	get_contract "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/usecases/item/get"
)

type useCase struct {
	repository item_contract.Repository
}

func New(repository item_contract.Repository) get_contract.UseCase {
	return &useCase{
		repository: repository,
	}
}

func (u *useCase) Execute(id string) (*entities.Item, error) {
	item, err := u.repository.Get(id)
	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, fmt.Errorf("item with id %s does not exist", id)
	}

	return item, nil
}
