package update_controller

import (
	"github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/constants/layouts"
	"github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/entities"
)

type Request struct {
	ID          string  `json:"id" validate:"required,uuid"`
	Title       *string `json:"title" validate:"required,min=3,max=100"`
	Description *string `json:"description" validate:"required,min=3,max=255"`
	DeadlineAt  *string `json:"deadline_at" validate:"omitempty,datetime=2006-01-02T15:04:05.000Z"`
}

type Response struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DeadlineAt  string `json:"deadline_at,omitempty"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func Map(item *entities.Item) Response {
	var deadlineAt string

	if item.DeadlineAt != nil {
		deadlineAt = item.DeadlineAt.Format(layouts.DatetimeLayout)
	}

	return Response{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		DeadlineAt:  deadlineAt,
		CreatedAt:   item.CreatedAt.Format(layouts.DatetimeLayout),
		UpdatedAt:   item.UpdatedAt.Format(layouts.DatetimeLayout),
	}
}
