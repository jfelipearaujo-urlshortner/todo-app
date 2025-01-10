package get_controller

import (
	"github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/constants/layouts"
	"github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/entities"
	"github.com/jfelipearaujo-urlshortner/todo-app/internal/external/shared"
)

type Request struct {
	ID string `json:"id" validate:"required,uuid"`
}

type Response struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DeadlineAt  string `json:"deadline_at,omitempty"`
	CompletedAt string `json:"completed_at,omitempty"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func Map(item *entities.Item) Response {
	return Response{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		DeadlineAt:  shared.FormatTime(layouts.DatetimeLayout, item.DeadlineAt),
		CompletedAt: shared.FormatTime(layouts.DatetimeLayout, item.CompletedAt),
		CreatedAt:   item.CreatedAt.Format(layouts.DatetimeLayout),
		UpdatedAt:   item.UpdatedAt.Format(layouts.DatetimeLayout),
	}
}
