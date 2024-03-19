package datastream

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"github.com/RevittConsulting/cdk-envs/pkg/utils"
)

type Handler struct {
	s *Service
}

func NewHandler(r chi.Router, s *Service) *Handler {
	h := &Handler{s: s}
	h.SetupRoutes(r)
	return h
}

func (h *Handler) SetupRoutes(router chi.Router) {
	log.Println("setting up routes for datastream...")
	router.Group(func(r chi.Router) {
		r.Get("/datastream", h.GetDatastream)
	})
}

func (h *Handler) GetDatastream(w http.ResponseWriter, r *http.Request) {
	total := h.s.GetTotalEntries()
	utils.WriteJSON(w, total)
}
