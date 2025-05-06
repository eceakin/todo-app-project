package main

import (
	"log"
	"net/http"
	"todo-app-project/internal/config"
	httpdelivery "todo-app-project/internal/delivery/http"
	"todo-app-project/internal/repository/mock"
	"todo-app-project/internal/usecase/auth"
	"todo-app-project/internal/usecase/todo"
	"todo-app-project/internal/utils"

	"github.com/gorilla/mux"
)

func main() {

	cfg := config.NewConfig()
	jwtUtil := utils.NewJWTUtil(cfg.JWTSecret)

	userRepo := mock.NewUserMockRepository()
	itemRepo := mock.NewTodoItemMockRepository()
	listRepo := mock.NewTodoListMockRepository(itemRepo)

	authUseCase := auth.NewAuthUseCase(userRepo, jwtUtil)
	listUseCase := todo.NewListUseCase(listRepo, itemRepo)
	itemUseCase := todo.NewTodoItemUseCase(itemRepo, listRepo)

	handler := httpdelivery.NewTodoHandler(listUseCase, itemUseCase)
	//authHandler := httpdelivery.NewAuthHandler(authUseCase)

	r := mux.NewRouter()

	// Login endpoint'i (kimlik doğrulama gerektirmez)
	r.HandleFunc("/login", httpdelivery.NewAuthHandler(authUseCase).Login).Methods("POST")

	// API subrouter'ı ve middleware
	api := r.PathPrefix("/api").Subrouter()
	api.Use(httpdelivery.AuthMiddleware(jwtUtil))

	// Korumalı endpoint'ler (AuthMiddleware uygulanır)
	api.HandleFunc("/lists", handler.CreateList).Methods("POST")
	api.HandleFunc("/lists/{id}", handler.UpdateList).Methods("PUT")
	api.HandleFunc("/lists/{id}", handler.DeleteList).Methods("DELETE")
	api.HandleFunc("/lists", handler.GetList).Methods("GET")
	api.HandleFunc("/lists/{id}/items", handler.GetItemsByListID).Methods("GET")
	api.HandleFunc("/lists/{id}/completion-rate", handler.CalculateCompletionRateHandler).Methods("GET")

	api.HandleFunc("/items", handler.AddItem).Methods("POST")
	api.HandleFunc("/items/{id}", handler.UpdateItem).Methods("PUT")
	api.HandleFunc("/items/{id}", handler.DeleteItem).Methods("DELETE")
	api.HandleFunc("/items/{id}", handler.CompleteItem).Methods("PATCH")
	api.HandleFunc("/items/{id}", handler.GetItemByID).Methods("GET")

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

/*"username": "admin",
  "password" : "admin"*/
