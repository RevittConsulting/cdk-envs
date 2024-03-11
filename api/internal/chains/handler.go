package chains

import (
	"fmt"
	"github.com/RevittConsulting/cdk-envs/pkg/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type HttpHandler struct {
	s *HttpService
}

func NewHandler(r chi.Router, s *HttpService) *HttpHandler {
	h := &HttpHandler{
		s: s,
	}
	h.SetupRoutes(r)
	return h
}

func (h *HttpHandler) SetupRoutes(router chi.Router) {
	fmt.Println("setting up routes for chains...")
	router.Group(func(r chi.Router) {
		r.Get("/chains", h.GetChains)
		r.Post("/chains", h.ChangeChainService)
		r.Get("/chains/stop", h.StopServices)
	})
}

func (h *HttpHandler) GetChains(w http.ResponseWriter, r *http.Request) {
	chain, err := h.s.GetChains(r.Context())
	if err != nil {
		utils.WriteErr(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, chain)
}

func (h *HttpHandler) ChangeChainService(w http.ResponseWriter, r *http.Request) {
	req := &ChainRequest{}
	err := utils.ReadJSON(r, req)
	if err != nil {
		utils.WriteErr(w, err, http.StatusBadRequest)
		return
	}
	chain, err := h.s.ChangeChainService(r.Context(), req.ServiceName)
	if err != nil {
		utils.WriteErr(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, chain)
}

func (h *HttpHandler) StopServices(w http.ResponseWriter, r *http.Request) {
	err := h.s.StopServices(r.Context())
	if err != nil {
		utils.WriteErr(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, "services stopped")
}
