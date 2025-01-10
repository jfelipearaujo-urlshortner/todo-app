package update_usecase

import (
	"time"

	"github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/entities"
	item_contract "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/repositories/item"
	update_contract "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/usecases/item/update"
)

type useCase struct {
	repository item_contract.Repository
}

func New(repository item_contract.Repository) update_contract.UseCase {
	return &useCase{
		repository: repository,
	}
}

func (u *useCase) Execute(id string, title *string, description *string, deadlineAt *time.Time) (*entities.Item, error) {
	item, err := u.repository.Get(id)
	if err != nil {
		return nil, err
	}

	item.Update(title, description, deadlineAt)

	if err := u.repository.Update(item); err != nil {
		return nil, err
	}

	return item, nil
}
