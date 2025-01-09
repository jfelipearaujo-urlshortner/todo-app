package create_usecase

import (
	"time"

	"github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/entities"
	item_contract "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/repositories/item"
	create_contract "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/usecases/item/create"
)

type useCase struct {
	repository item_contract.Repository
}

func New(repository item_contract.Repository) create_contract.UseCase {
	return &useCase{
		repository: repository,
	}
}

func (u *useCase) Execute(title string, description string, deadlineAt *time.Time) (*entities.Item, error) {
	item := entities.New(title, description, deadlineAt)

	if err := u.repository.Create(item); err != nil {
		return nil, err
	}

	return &item, nil
}
