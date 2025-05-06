package todo

import (
	"errors"
	"time"
	"todo-app-project/internal/domain"
)

type ListUseCase struct {
	todolistRepo domain.TodoListRepository // TodoList repository'si
	todoItemRepo domain.TodoItemRepository // TodoItem repository'si
}

func NewListUseCase(todolistRepo domain.TodoListRepository, todoItemRepo domain.TodoItemRepository) *ListUseCase {
	return &ListUseCase{
		todolistRepo: todolistRepo,
		todoItemRepo: todoItemRepo,
	}
}

func (l *ListUseCase) Create(userID int, name string) error {
	todoList := &domain.TodoList{
		UserID:    userID,
		Name:      name,
		CreatedAt: time.Now(),
	}
	return l.todolistRepo.Create(todoList)
}
func (l *ListUseCase) Update(userID, listID int, newName string) error {
	todoList, err := l.todolistRepo.GetByID(listID)
	if err != nil {
		return err
	}
	if todoList.UserID != userID {
		return errors.New("not authorized")
	}
	todoList.Name = newName
	todoList.UpdatedAt = time.Now()
	return l.todolistRepo.Update(todoList)
}

func (l *ListUseCase) SoftDelete(userID, listID int) error {
	todoList, err := l.todolistRepo.GetByID(listID)
	if err != nil {
		return err
	}
	if todoList.UserID != userID {
		return errors.New("not authorized") // yetkisiz erişim hatası döneriz
	} //
	return l.todolistRepo.SoftDelete(listID)
}
func (l *ListUseCase) GetByID(id int) (*domain.TodoList, error) {
	return l.todolistRepo.GetByID(id) // TodoList'i ID'sine göre alırız
}

func (l *ListUseCase) GetAll(userID int, isAdmin bool) ([]*domain.TodoList, error) {
	if isAdmin {
		return l.todolistRepo.GetAll(0)
	}
	return l.todolistRepo.GetAll(userID)
}
