package delete_usecase

import (
	item_contract "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/repositories/item"
	delete_contract "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/usecases/item/delete"
)

type useCase struct {
	repository item_contract.Repository
}

func New(repository item_contract.Repository) delete_contract.UseCase {
	return &useCase{
		repository: repository,
	}
}

func (u *useCase) Execute(id string) error {
	if err := u.repository.Delete(id); err != nil {
		return err
	}

	return nil
}
