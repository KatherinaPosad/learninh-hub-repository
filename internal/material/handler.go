package material

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

func (h *Handler) GetMaterials(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ctx := r.Context()

		materials, err := h.service.GetMaterials(ctx)
		if err != nil {
			http.Error(w, "Failed to get materials", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(materials)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

	}
}

func (h *Handler) CreateMaterial(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var material Material
		ctx := r.Context()

		err := json.NewDecoder(r.Body).Decode(&material)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		err = h.service.CreateMaterial(ctx, &material)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(&material)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
