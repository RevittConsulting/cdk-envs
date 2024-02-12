package chains

import (
	"fmt"
	"github.com/RevittConsulting/cdk-envs/pkg/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Handler struct {
	s *Service
}

func NewHandler(r chi.Router, s *Service) *Handler {
	h := &Handler{
		s: s,
	}
	h.SetupRoutes(r)
	return h
}

func (h *Handler) SetupRoutes(router chi.Router) {
	fmt.Println("setting up routes for chains...")
	router.Group(func(r chi.Router) {
		r.Get("/chains", h.GetChains)
	})
}

func (h *Handler) GetChains(w http.ResponseWriter, r *http.Request) {
	chain, err := h.s.GetChains(r.Context())
	if err != nil {
		utils.WriteErr(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, chain)
}
