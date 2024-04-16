package api

import (
	"encoding/json"
	"net/http"

	"integration-crud/db"

	"github.com/gorilla/mux"
)

type Handler struct {
	store db.Store
}

func NewHandler(store db.Store) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.store.GetItems()
	if err != nil {
		http.Error(w, "Failed to get items", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items)
}

func (h *Handler) GetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	item, err := h.store.GetItem(id)
	if err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(item)
}

func (h *Handler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item db.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "Invalid item data", http.StatusBadRequest)
		return
	}

	err = h.store.CreateItem(&item)
	if err != nil {
		http.Error(w, "Failed to create item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var item db.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "Invalid item data", http.StatusBadRequest)
		return
	}

	err = h.store.UpdateItem(id, &item)
	if err != nil {
		http.Error(w, "Failed to update item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.store.DeleteItem(id)
	if err != nil {
		http.Error(w, "Failed to delete item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
