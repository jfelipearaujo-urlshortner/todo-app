package list_controller

import (
	"github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/constants/layouts"
	"github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/entities"
	"github.com/jfelipearaujo-urlshortner/todo-app/internal/external/shared"
)

type Response struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DeadlineAt  string `json:"deadline_at,omitempty"`
	CompletedAt string `json:"completed_at,omitempty"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func Map(items []entities.Item) []Response {
	var response []Response

	if len(items) == 0 {
		return response
	}

	for _, item := range items {
		response = append(response, Response{
			ID:          item.ID,
			Title:       item.Title,
			Description: item.Description,
			DeadlineAt:  shared.FormatTime(layouts.DatetimeLayout, item.DeadlineAt),
			CompletedAt: shared.FormatTime(layouts.DatetimeLayout, item.CompletedAt),
			CreatedAt:   item.CreatedAt.Format(layouts.DatetimeLayout),
			UpdatedAt:   item.UpdatedAt.Format(layouts.DatetimeLayout),
		})
	}

	return response
}
