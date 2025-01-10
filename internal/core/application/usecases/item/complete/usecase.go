package complete_usecase

import (
	"fmt"

	item_contract "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/repositories/item"
	complete_contract "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/usecases/item/complete"
)

type useCase struct {
	repository item_contract.Repository
}

func New(repository item_contract.Repository) complete_contract.UseCase {
	return &useCase{
		repository: repository,
	}
}

func (u *useCase) Execute(id string) error {
	item, err := u.repository.Get(id)
	if err != nil {
		return err
	}

	if item == nil {
		return fmt.Errorf("item with id %s does not exist", id)
	}

	if item.Done {
		return fmt.Errorf("item with id %s is already completed", id)
	}

	item.Complete()

	if err := u.repository.Update(item); err != nil {
		return err
	}

	return nil
}
