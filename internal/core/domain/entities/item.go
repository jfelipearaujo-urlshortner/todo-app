package entities

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID          string
	Title       string
	Description string
	Done        bool
	DeadlineAt  *time.Time
	CompletedAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func New(title string, description string, deadlineAt *time.Time) Item {
	return Item{
		ID:          uuid.NewString(),
		Title:       title,
		Description: description,
		DeadlineAt:  deadlineAt,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (i *Item) Update(title *string, description *string, deadlineAt *time.Time) {
	if title != nil {
		i.Title = *title
	}

	if description != nil {
		i.Description = *description
	}

	if deadlineAt != nil {
		i.DeadlineAt = deadlineAt
	}

	if title != nil || description != nil || deadlineAt != nil {
		i.UpdatedAt = time.Now()
	}
}

func (i *Item) Complete() {
	i.Done = true
	completedAt := time.Now()
	i.CompletedAt = &completedAt
}
