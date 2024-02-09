package chain

import (
	"fmt"
	"github.com/RevittConsulting/cdk-envs/pkg/util"
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
	fmt.Println("setting up routes for orders")
	router.Group(func(r chi.Router) {
		r.Get("/chain", h.GetChains)
		r.Get("/chain/block", h.GetHighestBlock)
	})
}

func (h *Handler) GetHighestBlock(w http.ResponseWriter, r *http.Request) {
	block, err := h.s.GetHighestBlock(r.Context())
	if err != nil {
		util.WriteErr(w, err, http.StatusInternalServerError)
		return
	}

	util.WriteJSON(w, block)
}

func (h *Handler) GetChains(w http.ResponseWriter, r *http.Request) {
	chain, err := h.s.GetChains(r.Context())
	if err != nil {
		util.WriteErr(w, err, http.StatusInternalServerError)
		return
	}

	util.WriteJSON(w, chain)
}
