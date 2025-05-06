package domain

type TodoListRepository interface {
	Create(todoList *TodoList) error
	Update(todoList *TodoList) error
	SoftDelete(id int) error
	GetByID(id int) (*TodoList, error)
	GetAll(userID int) ([]*TodoList, error)
	CalculateCompletionRate(listID int) (float64, error)
}
type TodoItemRepository interface {
	Create(todoItem *TodoItem) error
	Update(todoItem *TodoItem) error
	SoftDelete(id int) error
	GetByID(id int) (*TodoItem, error)
	GetByListID(listID int) ([]*TodoItem, error)
	CompleteItem(id int) error
}
type UserRepository interface {
	GetByUsername(username string) (*User, error)
	GetByID(id int) (*User, error)
}
