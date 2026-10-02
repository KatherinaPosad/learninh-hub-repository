package reviews

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

func (h *Handler) GetReviews(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ctx := r.Context()
		reviews, err := h.service.GetReviews(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(reviews)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) CreateReview(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var review Review

		err := json.NewDecoder(r.Body).Decode(&review)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		ctx := r.Context()

		err = h.service.CreateReview(ctx, &review)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(&review)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) DeleteReview(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		var review Review
		ctx := r.Context()

		err := json.NewDecoder(r.Body).Decode(&review)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		rowsAffected, err := h.service.DeleteReview(ctx, review)
		if err != nil {
			http.Error(w, "failed to delete review", http.StatusInternalServerError)
			return
		}
		if rowsAffected == 0 {
			http.Error(w, "Review not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
