package mock

import (
	"errors"
	"sync"
	"todo-app-project/internal/domain"
)

type UserMockRepository struct {
	users map[int]*domain.User
	mu    sync.Mutex
}

func NewUserMockRepository() *UserMockRepository {
	repo := &UserMockRepository{
		users: make(map[int]*domain.User),
	}
	repo.initializeUsers()
	return repo
}

func (r *UserMockRepository) initializeUsers() {
	r.mu.Lock()
	defer r.mu.Unlock()

	defaultUsers := []domain.User{
		{
			ID:       1,
			Username: "admin",
			Password: "admin",
			Role:     domain.AdminRole,
		},
		{
			ID:       2,
			Username: "user",
			Password: "user",
			Role:     domain.UserRole,
		},
		{
			ID:       3,
			Username: "guest",
			Password: "guest",
			Role:     domain.UserRole},
	}

	for _, user := range defaultUsers {
		r.users[user.ID] = &user
	}
}

func (r *UserMockRepository) GetByUsername(username string) (*domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range r.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *UserMockRepository) GetByID(id int) (*domain.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}
