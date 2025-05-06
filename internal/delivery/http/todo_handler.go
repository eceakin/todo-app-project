package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-app-project/internal/usecase/todo"

	"github.com/gorilla/mux"
)

type TodoHandler struct {
	itemUsecase *todo.TodoItemUseCase // TodoUseCase arayüzü
	listUsecase *todo.ListUseCase     // ItemUseCase arayüzü
}

func NewTodoHandler(listUC *todo.ListUseCase, itemUC *todo.TodoItemUseCase) *TodoHandler {
	return &TodoHandler{
		listUsecase: listUC,
		itemUsecase: itemUC,
	}
}

// -- list endpointleri -- //

func (h *TodoHandler) CreateList(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(ContextUserID).(int) // context'ten userID'yi alırız

	var body struct {
		Name string `json:"name"` // isim alanı
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest) // hata varsa döneriz
		return
	}
	if err := h.listUsecase.Create(userID, body.Name); err != nil {
		http.Error(w, "failed to create list", http.StatusInternalServerError) // hata varsa döneriz
		return
	}
	w.WriteHeader(http.StatusCreated) // 201 Created döneriz
}

func (h *TodoHandler) GetList(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(ContextUserID).(int) // context'ten userID'yi alırız
	role := r.Context().Value(ContextRole).(string)  // context'ten rolü alırız
	isAdmin := role == "admin"                       // admin olup olmadığını kontrol ederiz

	lists, err := h.listUsecase.GetAll(userID, isAdmin) // tüm listeleri alırız
	if err != nil {
		http.Error(w, "failed to get lists", http.StatusInternalServerError) // hata varsa döneriz
		return
	}
	json.NewEncoder(w).Encode(lists) // listeleri döneriz

}
func (h *TodoHandler) UpdateList(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(ContextUserID).(int)
	listID, _ := strconv.Atoi(mux.Vars(r)["id"])

	var body struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if err := h.listUsecase.Update(userID, listID, body.Name); err != nil {
		http.Error(w, "failed to update list", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (h *TodoHandler) DeleteList(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(ContextUserID).(int)
	listID, _ := strconv.Atoi(mux.Vars(r)["id"])

	if err := h.listUsecase.SoftDelete(userID, listID); err != nil {
		http.Error(w, "failed to delete list", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// madde endpointleri //

func (h *TodoHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(ContextUserID).(int) // context'ten userID'yi alırız
	var body struct {
		ListID  int    `json:"list_id"` // liste ID'si alanı
		Content string `json:"content"` // içerik alanı
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest) // hata varsa döneriz
		return
	}
	if err := h.itemUsecase.Create(userID, body.ListID, body.Content); err != nil {
		http.Error(w, "failed to create item", http.StatusInternalServerError) // hata varsa döneriz
		return
	}
	w.WriteHeader(http.StatusCreated) // 201 Created döneriz

}
func (h *TodoHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(ContextUserID).(int)
	itemID, _ := strconv.Atoi(mux.Vars(r)["id"])

	var body struct {
		Content string `json:"content"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if err := h.itemUsecase.Update(userID, itemID, body.Content); err != nil {
		http.Error(w, "failed to update item", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (h *TodoHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(ContextUserID).(int)
	itemID, _ := strconv.Atoi(mux.Vars(r)["id"])

	if err := h.itemUsecase.SoftDelete(userID, itemID); err != nil {
		http.Error(w, "failed to delete item", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *TodoHandler) CompleteItem(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(ContextUserID).(int) // context'ten userID'yi alırız
	idStr := mux.Vars(r)["id"]                       // URL'den ID'yi alırız
	id, _ := strconv.Atoi(idStr)                     // ID'yi int'e çeviririz

	if err := h.itemUsecase.CompleteItem(userID, id); err != nil {
		http.Error(w, "failed to mark item as done", http.StatusInternalServerError) // hata varsa döneriz
		return
	}
	w.WriteHeader(http.StatusOK) // 200 OK döneriz

}
func (h *TodoHandler) GetItemsByListID(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(ContextUserID).(int)
	listID, _ := strconv.Atoi(mux.Vars(r)["id"])

	items, err := h.itemUsecase.GetByListID(listID, userID)
	if err != nil {
		http.Error(w, "failed to get items", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(items)
}
func (h *TodoHandler) GetItemByID(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(ContextUserID).(int)
	itemID, _ := strconv.Atoi(mux.Vars(r)["id"])

	item, err := h.itemUsecase.GetByID(userID, itemID)
	if err != nil {
		http.Error(w, "failed to get item", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(item)
}
func (h *TodoHandler) CalculateCompletionRateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	listIDStr := vars["id"]

	listID, err := strconv.Atoi(listIDStr)
	if err != nil {
		http.Error(w, "Invalid list ID", http.StatusBadRequest)
		return
	}

	rate, err := h.listUsecase.CalculateCompletionRate(listID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := map[string]float64{
		"completion_rate": rate,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
