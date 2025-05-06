package mock

import (
	"errors"
	"sync"
	"time"
	"todo-app-project/internal/domain"
)

type TodoListMockRepository struct {
	todoLists    map[int]*domain.TodoList
	todoItemRepo domain.TodoItemRepository
	mu           sync.Mutex
	nextID       int
}

func NewTodoListMockRepository(itemRepo domain.TodoItemRepository) *TodoListMockRepository {
	return &TodoListMockRepository{
		todoLists:    make(map[int]*domain.TodoList),
		todoItemRepo: itemRepo,
		nextID:       1,
	}
}

func (r *TodoListMockRepository) Create(todoList *domain.TodoList) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	todoList.ID = r.nextID
	todoList.CreatedAt = now
	todoList.UpdatedAt = now
	r.todoLists[todoList.ID] = todoList
	r.nextID++
	return nil
}

func (r *TodoListMockRepository) Update(todoList *domain.TodoList) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	existingList, exists := r.todoLists[todoList.ID]
	if !exists || existingList.IsDeleted() {
		return errors.New("todo list not found")
	}
	todoList.UpdatedAt = time.Now()
	r.todoLists[todoList.ID] = todoList
	return nil

}

func (r *TodoListMockRepository) SoftDelete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	todoList, exists := r.todoLists[id]
	if !exists || todoList.IsDeleted() {
		return errors.New("todo list not found")
	}
	now := time.Now()
	todoList.DeletedAt = &now
	todoList.UpdatedAt = now
	return nil

}

func (r *TodoListMockRepository) GetAll(userID int) ([]*domain.TodoList, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var todoLists []*domain.TodoList
	for _, todoList := range r.todoLists {
		if !todoList.IsDeleted() && todoList.UserID == userID {
			todoLists = append(todoLists, todoList)
		}
	}
	return todoLists, nil
}

func (r *TodoListMockRepository) GetByID(id int) (*domain.TodoList, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	todoList, exists := r.todoLists[id]
	if !exists || todoList.IsDeleted() {
		return nil, errors.New("todo list not found")
	}
	return todoList, nil
}
func (r *TodoListMockRepository) CalculateCompletionRate(listID int) (float64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	todoList, exists := r.todoLists[listID]
	if !exists || todoList.IsDeleted() {
		return 0, errors.New("todo list not found")
	}

	items, err := r.todoItemRepo.GetByListID(listID)
	if err != nil {
		return 0, err
	}
	if len(items) == 0 {
		return 0, nil
	}

	completed := 0
	for _, item := range items {
		if item.IsCompleted {
			completed++
		}
	}

	rate := (float64(completed) / float64(len(items))) * 100
	return rate, nil
}
