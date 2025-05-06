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

	// Buraya route'lar eklenecek

	r.HandleFunc("/login", httpdelivery.NewAuthHandler(authUseCase).Login).Methods("POST")
	api := r.PathPrefix("/api").Subrouter()       // api prefix'i alırız
	api.Use(httpdelivery.AuthMiddleware(jwtUtil)) // JWT middleware'i alırız

	r.HandleFunc("/lists", handler.CreateList).Methods("POST")
	r.HandleFunc("/lists/{id}", handler.UpdateList).Methods("PUT")
	r.HandleFunc("/lists/{id}", handler.DeleteList).Methods("DELETE")
	r.HandleFunc("/lists", handler.GetList).Methods("GET")
	r.HandleFunc("/lists/{id}/items", handler.GetItemsByListID).Methods("GET")

	r.HandleFunc("/lists/{id}/completion-rate", handler.CalculateCompletionRateHandler).Methods("GET")

	r.HandleFunc("/items", handler.AddItem).Methods("POST")
	r.HandleFunc("/items/{id}", handler.UpdateItem).Methods("PUT")
	r.HandleFunc("/items/{id}", handler.DeleteItem).Methods("DELETE")
	r.HandleFunc("/items/{id}", handler.CompleteItem).Methods("PATCH")
	r.HandleFunc("/items/{id}", handler.GetItemByID).Methods("GET")

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
