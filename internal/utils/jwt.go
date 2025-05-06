package utils

import (
	"time"
	"todo-app-project/internal/domain"

	"github.com/golang-jwt/jwt/v4"
)

type JWTUtil struct {
	secretKey string
}

func (j *JWTUtil) SecretKey() string {
	return j.secretKey
}
func NewJWTUtil(secretKey string) *JWTUtil {
	return &JWTUtil{
		secretKey: secretKey,
	}
}

func (j *JWTUtil) GenerateToken(user *domain.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    string(user.Role),                     // kullanıcının rolü
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // tokenin süresi 24 saat
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // token oluştururuz
	return token.SignedString([]byte(j.secretKey))             // tokeni döneriz
}
