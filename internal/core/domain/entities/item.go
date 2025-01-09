package entities

import "time"

type Item struct {
	ID          string
	Title       string
	Description string
	Done        bool
	DeadlineAt  time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
