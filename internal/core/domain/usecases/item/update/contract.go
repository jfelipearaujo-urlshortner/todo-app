package update_contract

import (
	"time"

	"github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/entities"
)

type UseCase interface {
	Execute(id string, title *string, description *string, deadlineAt *time.Time) (*entities.Item, error)
}
