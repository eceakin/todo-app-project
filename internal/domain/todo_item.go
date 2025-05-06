package domain

import (
	"time"
)

type TodoItem struct {
	ID          int        `json:"id"`
	ListID      int        `json:"list_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	Content     string     `json:"content"`
	IsCompleted bool       `json:"is_completed"`
	UserID      int        `json:"user_id"`
}

func (i *TodoItem) IsDeleted() bool {
	return i.DeletedAt != nil
}
