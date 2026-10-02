package category

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetCategories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ctx := r.Context()

		categories, err := h.service.GetCategories(ctx)
		if err != nil {
			http.Error(w, "Failed to get categories", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(categories)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var category Category
		ctx := r.Context()

		err := json.NewDecoder(r.Body).Decode(&category)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		err = h.service.CreateCategory(ctx, &category)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(&category)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
