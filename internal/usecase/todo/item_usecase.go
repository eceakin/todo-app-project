package todo

import (
	"errors"
	"time"
	"todo-app-project/internal/domain"
)

type TodoItemUseCase struct {
	todoItemRepo domain.TodoItemRepository // TodoItem repository'si
	todoListRepo domain.TodoListRepository // TodoList repository'si
}

func NewTodoItemUseCase(todoItemRepo domain.TodoItemRepository, todoListRepo domain.TodoListRepository) *TodoItemUseCase {
	return &TodoItemUseCase{
		todoItemRepo: todoItemRepo,
		todoListRepo: todoListRepo,
	}
}

func (u *TodoItemUseCase) Create(userID, listID int, content string) error {
	todoList, err := u.todoListRepo.GetByID(listID) // TodoList'i ID'sine göre alırız
	if err != nil {
		return err // hata varsa döneriz
	}
	if todoList.UserID != userID {
		return errors.New("not authorized") // yetkisiz erişim hatası döneriz
	}
	todoItem := &domain.TodoItem{
		ListID:    listID,
		Content:   content,
		CreatedAt: time.Now(),
	}
	return u.todoItemRepo.Create(todoItem) // TodoItem'i oluştururuz
}
func (u *TodoItemUseCase) Update(userID, itemID int, newContent string) error {
	todoItem, err := u.todoItemRepo.GetByID(itemID) // TodoItem'i ID'sine göre alırız
	if err != nil {
		return err // hata varsa döneriz
	}
	todoList, err := u.todoListRepo.GetByID(todoItem.ListID) // TodoList'i ID'sine göre alırız
	if err != nil {
		return errors.New("not the owner") // hata varsa döneriz
	}
	if todoList.UserID != userID {
		return errors.New("not the owner") // yetkisiz erişim hatası döneriz
	}
	todoItem.Content = newContent          // yeni içeriği atarız
	todoItem.UpdatedAt = time.Now()        // güncellenme zamanını atarız
	return u.todoItemRepo.Update(todoItem) // TodoItem'i güncelleriz
}

func (u *TodoItemUseCase) SoftDelete(userID, itemID int) error {
	todoItem, err := u.todoItemRepo.GetByID(itemID) // TodoItem'i ID'sine göre alırız
	if err != nil {
		return err // hata varsa döneriz
	}
	todoList, err := u.todoListRepo.GetByID(todoItem.ListID) // TodoList'i ID'sine göre alırız
	if err != nil {
		return errors.New("not the owner") // hata varsa döneriz
	}
	if todoList.UserID != userID {
		return errors.New("not the owner") // yetkisiz erişim hatası döneriz
	}
	return u.todoItemRepo.SoftDelete(itemID) // TodoItem'i sileriz
}
func (u *TodoItemUseCase) GetByListID(listID, userID int) ([]*domain.TodoItem, error) {
	todoList, err := u.todoListRepo.GetByID(listID) // TodoList'i ID'sine göre alırız
	if err != nil {
		return nil, err // hata varsa döneriz
	}
	if todoList.UserID != userID {
		return nil, errors.New("not authorized") // yetkisiz erişim hatası döneriz
	}
	return u.todoItemRepo.GetByListID(listID) // TodoItem'leri list ID'sine göre alırız
}
func (u *TodoItemUseCase) CompleteItem(userID, itemID int) error {
	todoItem, err := u.todoItemRepo.GetByID(itemID) // TodoItem'i ID'sine göre alırız
	if err != nil {
		return err // hata varsa döneriz
	}
	todoList, err := u.todoListRepo.GetByID(todoItem.ListID) // TodoList'i ID'sine göre alırız
	if err != nil {
		return errors.New("not the owner") // hata varsa döneriz
	}
	if todoList.UserID != userID {
		return errors.New("not the owner") // yetkisiz erişim hatası döneriz
	}
	return u.todoItemRepo.CompleteItem(itemID) // TodoItem'i tamamlanmış olarak işaretleriz

}
func (u *TodoItemUseCase) GetByID(userID, id int) (*domain.TodoItem, error) {
	item, err := u.todoItemRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if item.UserID != userID {
		return nil, errors.New("not authorized to view this item")
	}
	return item, nil
}
