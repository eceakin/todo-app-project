package domain

import (
	"time"
)

type TodoList struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
	CompletionRate float64    `json:"completion_rate"`
	UserID         int        `json:"user_id"`
}

func (l *TodoList) IsDeleted() bool {
	return l.DeletedAt != nil
}
