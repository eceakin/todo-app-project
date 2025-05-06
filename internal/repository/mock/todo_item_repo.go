package mock

import (
	"errors"
	"sync"
	"time"
	"todo-app-project/internal/domain"
)

type TodoItemMockRepository struct {
	todoItems map[int]*domain.TodoItem
	mu        sync.Mutex
	nextID    int
}

func NewTodoItemMockRepository() *TodoItemMockRepository {
	return &TodoItemMockRepository{
		todoItems: make(map[int]*domain.TodoItem),
		nextID:    1,
	}
}

func (r *TodoItemMockRepository) Create(todoItem *domain.TodoItem) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	todoItem.ID = r.nextID
	todoItem.CreatedAt = now
	todoItem.UpdatedAt = now
	r.todoItems[todoItem.ID] = todoItem
	r.nextID++
	return nil
}

func (r *TodoItemMockRepository) Update(todoItem *domain.TodoItem) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	existingItem, exists := r.todoItems[todoItem.ID]
	if !exists || existingItem.IsDeleted() {
		return errors.New("todo item not found")
	}
	todoItem.UpdatedAt = time.Now()
	r.todoItems[todoItem.ID] = todoItem
	return nil
}

func (r *TodoItemMockRepository) SoftDelete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	todoItem, exists := r.todoItems[id]
	if !exists || todoItem.IsDeleted() {
		return errors.New("todo item not found")
	}
	now := time.Now()
	todoItem.DeletedAt = &now
	todoItem.UpdatedAt = now
	return nil
}

func (r *TodoItemMockRepository) GetByID(id int) (*domain.TodoItem, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	todoItem, exists := r.todoItems[id]
	if !exists || todoItem.IsDeleted() {
		return nil, errors.New("todo item not found")
	}
	return todoItem, nil
}

func (r *TodoItemMockRepository) GetByListID(listID int) ([]*domain.TodoItem, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var todoItems []*domain.TodoItem
	for _, todoItem := range r.todoItems {
		if !todoItem.IsDeleted() && todoItem.ListID == listID {
			todoItems = append(todoItems, todoItem)
		}
	}
	return todoItems, nil
}

func (r *TodoItemMockRepository) CompleteItem(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	todoItem, exists := r.todoItems[id]
	if !exists || todoItem.IsDeleted() {
		return errors.New("todo item not found")
	}
	todoItem.IsCompleted = true
	todoItem.UpdatedAt = time.Now()
	return nil
}
