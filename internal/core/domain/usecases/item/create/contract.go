package create_contract

import (
	"time"

	"github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/entities"
)

type UseCase interface {
	Execute(title string, description string, deadlineAt *time.Time) (*entities.Item, error)
}
