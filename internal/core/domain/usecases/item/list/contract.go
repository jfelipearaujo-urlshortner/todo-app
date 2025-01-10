package list_contract

import "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/entities"

type UseCase interface {
	Execute() []entities.Item
}
