package auth

import (
	"errors"
	"todo-app-project/internal/domain"
	"todo-app-project/internal/utils"
)

type AuthUseCase struct {
	userRepo domain.UserRepository // kullanıcı repository'si
	jwtUtil  *utils.JWTUtil        // JWT utilitesi
}

func NewAuthUseCase(userRepo domain.UserRepository, jwtUtil *utils.JWTUtil) *AuthUseCase {
	return &AuthUseCase{
		userRepo: userRepo,
		jwtUtil:  jwtUtil,
	}
}

func (a *AuthUseCase) Login(username, password string) (string, error) {
	user, err := a.userRepo.GetByUsername(username)
	if err != nil {
		return "", errors.New("invalid") // hata varsa döneriz
	}

	if user.Password != password {
		return "", errors.New("invalid") // şifre yanlışsa döneriz
	}

	token, err := a.jwtUtil.GenerateToken(user) // token oluştururuz
	if err != nil {
		return "", errors.New("failed to generate token") // hata varsa döneriz
	}
	return token, nil // tokeni döneriz

}
