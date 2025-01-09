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
