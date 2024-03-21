package tx

import (
	"github.com/RevittConsulting/chain-dev-utils/pkg/utils"
	"github.com/RevittConsulting/logger"
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
	logger.Log().Info("setting up routes for tx...")
	router.Group(func(r chi.Router) {
		r.Post("/tx", h.createTx)
	})
}

func (h *HttpHandler) createTx(w http.ResponseWriter, r *http.Request) {
	req := &Request{}
	err := utils.ReadJSON(r, req)
	if err != nil {
		utils.WriteErr(w, err, http.StatusBadRequest)
		return
	}

	tx, err := h.s.CreateTx(r.Context(), req)
	if err != nil {
		utils.WriteErr(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, tx)
}
