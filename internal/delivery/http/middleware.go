/*
	package http

import (

	"context"
	"net/http"
	"strings"
	"todo-app-project/internal/utils"

	"github.com/golang-jwt/jwt/v4"

)

type contextKey string // context anahtarı türü

const (

	ContextUserID = "user_id" // context anahtarı
	ContextRole   = "role"    // context anahtarı

)

	func AuthMiddleware(jwtUtil *utils.JWTUtil) func(http.Handler) http.Handler {
		return func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				authHeader := r.Header.Get("Authorization") // Authorization başlığını alırız
				if !strings.HasPrefix(authHeader, "Bearer ") {
					http.Error(w, "invalid token", http.StatusUnauthorized) // hata varsa döneriz
					return
				}
				tokenString := strings.TrimPrefix(authHeader, "Bearer ") // tokeni alırız

				claims := &jwt.MapClaims{} // claims'i alırız
				_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
					return []byte(jwtUtil.SecretKey()), nil // secret key'i alırız
				},
				)
				if err != nil {
					http.Error(w, "invalid token", http.StatusUnauthorized) // hata varsa döneriz
					return
				}
				userID := int((*claims)["user_id"].(float64)) // userID'yi alırız
				role := (*claims)["role"].(string)            // rolü alırız

				ctx := context.WithValue(r.Context(), ContextUserID, userID) // context'e userID'yi ekleriz
				ctx = context.WithValue(ctx, ContextRole, role)              // context'e rolü ekleriz

				next.ServeHTTP(w, r.WithContext(ctx)) // sonraki handler'ı çağırırız
			})
		}
	} // AuthMiddleware fonksiyonu, JWT token doğrulama işlemini yapar ve kullanıcı bilgilerini context'e ekler
*/
package http

import (
	"context"
	"log"
	"net/http"
	"strings"
	"todo-app-project/internal/utils"

	"github.com/golang-jwt/jwt/v4"
)

type contextKey string

const (
	ContextUserID = "user_id"
	ContextRole   = "role"
)

func AuthMiddleware(jwtUtil *utils.JWTUtil) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("AuthMiddleware: Çalıştı")
			authHeader := r.Header.Get("Authorization")
			log.Printf("AuthMiddleware: Gelen Authorization Header: %s", authHeader) // Loglama eklendi

			if !strings.HasPrefix(authHeader, "Bearer ") {
				log.Println("AuthMiddleware: Geçersiz token formatı") // Loglama eklendi
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			claims := &jwt.MapClaims{}
			_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(jwtUtil.SecretKey()), nil
			})
			if err != nil {
				log.Printf("AuthMiddleware: Token ayrıştırma hatası: %v", err) // Loglama eklendi
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			userIDFloat, ok := (*claims)["user_id"].(float64)
			if !ok {
				log.Println("AuthMiddleware: user_id claim'i bulunamadı veya yanlış tipte") // Loglama eklendi
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}
			userID := int(userIDFloat)
			log.Printf("AuthMiddleware: Çözülen UserID: %d", userID) // Loglama eklendi

			role, ok := (*claims)["role"].(string)
			if !ok {
				log.Println("AuthMiddleware: role claim'i bulunamadı veya yanlış tipte") // Loglama eklendi
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}
			log.Printf("AuthMiddleware: Çözülen Rol: %s", role) // Loglama eklendi

			ctx := context.WithValue(r.Context(), ContextUserID, userID)
			ctx = context.WithValue(ctx, ContextRole, role)

			log.Println("AuthMiddleware: UserID ve Rol context'e eklendi") // Loglama eklendi
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
