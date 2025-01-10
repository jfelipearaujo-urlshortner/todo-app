package get_contract

import "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/entities"

type UseCase interface {
	Execute(id string) (*entities.Item, error)
}
