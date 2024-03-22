package buckets

import (
	"github.com/RevittConsulting/chain-dev-utils/internal/types"
	"github.com/RevittConsulting/chain-dev-utils/pkg/utils"
	"github.com/RevittConsulting/logger"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
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
	logger.Log().Info("setting up routes for buckets...")
	router.Group(func(r chi.Router) {
		r.Post("/buckets/data/change", h.ChangeDB)
		r.Get("/buckets/data/list", h.listDataSource)
		r.Get("/buckets", h.getBuckets)

		r.Route("/buckets/{bucketName}", func(r chi.Router) {
			r.Use(h.TryGetBucket)

			r.Get("/count", h.getKeysCount)
			r.Get("/pages/{pageNum}/{pageLen}", h.getPages)
			r.Get("/count/{length}", h.countLength)
			r.Get("/count/{length}/keys", h.keysCountLength)
			r.Get("/count/{length}/values", h.valuesCountLength)
			r.Get("/keys/{key}", h.searchByKey)
			r.Get("/values/{value}", h.searchByValue)
		})
	})
}

func (h *HttpHandler) ChangeDB(w http.ResponseWriter, r *http.Request) {
	req := &DBRequest{}
	err := utils.ReadJSON(r, req)
	if err != nil {
		utils.WriteErr(w, err, http.StatusBadRequest)
		return
	}
	if err = h.s.ChangeDB(req.Path); err != nil {
		utils.WriteErr(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, map[string]string{"message": "Database changed"})
}

func (h *HttpHandler) listDataSource(w http.ResponseWriter, r *http.Request) {
	dataSource, err := h.s.ListDataSource()
	if err != nil {
		utils.WriteErr(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, dataSource)
}

func (h *HttpHandler) getBuckets(w http.ResponseWriter, r *http.Request) {
	buckets, err := h.s.ListBuckets()
	if err != nil {
		utils.WriteErr(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, buckets)
}

func (h *HttpHandler) getKeysCount(w http.ResponseWriter, r *http.Request) {
	count, err := h.s.KeysCount(r.Context())
	if err != nil {
		utils.WriteErr(w, err, http.StatusInternalServerError)
		return
	}

	response := map[string]uint64{"count": count}

	utils.WriteJSON(w, response)
}

func (h *HttpHandler) getPages(w http.ResponseWriter, r *http.Request) {
	pageNum, err := strconv.Atoi(chi.URLParam(r, "pageNum"))
	if err != nil {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}
	pageLen, err := strconv.Atoi(chi.URLParam(r, "pageLen"))
	if err != nil {
		http.Error(w, "Invalid page length", http.StatusBadRequest)
		return
	}

	maxPageLen := 200

	if pageLen > maxPageLen {
		pageLen = maxPageLen
	}

	pages, err := h.s.GetPage(r.Context(), pageNum, pageLen)
	if err != nil {
		utils.WriteErr(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, pages)
}

func (h *HttpHandler) countLength(w http.ResponseWriter, r *http.Request) {
	length, err := strconv.ParseUint(chi.URLParam(r, "length"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	count, _, err := h.s.KeysCountLength(r.Context(), length)
	if err != nil {
		utils.WriteErr(w, err, http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{"count": count}

	utils.WriteJSON(w, response)
}

func (h *HttpHandler) keysCountLength(w http.ResponseWriter, r *http.Request) {
	length, err := strconv.ParseUint(chi.URLParam(r, "length"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	count, keys, err := h.s.KeysCountLength(r.Context(), length)
	if err != nil {
		utils.WriteErr(w, err, http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{"count": count, "kv": keys}

	utils.WriteJSON(w, response)
}

func (h *HttpHandler) valuesCountLength(w http.ResponseWriter, r *http.Request) {
	length, err := strconv.ParseUint(chi.URLParam(r, "length"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	count, values, err := h.s.ValuesCountLength(r.Context(), length)
	if err != nil {
		utils.WriteErr(w, err, http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{"count": count, "kv": values}

	utils.WriteJSON(w, response)
}

func (h *HttpHandler) searchByKey(w http.ResponseWriter, r *http.Request) {
	searchKey := chi.URLParam(r, "key")

	foundValues, err := h.s.SearchByKey(r.Context(), searchKey)
	if err != nil {
		utils.WriteErr(w, err, http.StatusInternalServerError)
		return
	}

	if foundValues == nil {
		http.Error(w, "no keys found", http.StatusNotFound)
		return
	}

	var res []types.KeyValuePairString
	for _, value := range foundValues {
		res = append(res, types.KeyValuePairString{Key: searchKey, Value: value})
	}

	utils.WriteJSON(w, res)
}

func (h *HttpHandler) searchByValue(w http.ResponseWriter, r *http.Request) {
	searchValue := chi.URLParam(r, "value")

	foundValues, err := h.s.SearchByValue(r.Context(), searchValue)
	if err != nil {
		utils.WriteErr(w, err, http.StatusInternalServerError)
		return
	}

	if foundValues == nil {
		http.Error(w, "no values found", http.StatusNotFound)
		return
	}

	var res []types.KeyValuePairString
	for _, key := range foundValues {
		res = append(res, types.KeyValuePairString{Key: key, Value: searchValue})
	}

	utils.WriteJSON(w, res)
}
