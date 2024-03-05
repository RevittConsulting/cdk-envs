package buckets

import (
	"fmt"
	"github.com/RevittConsulting/cdk-envs/pkg/utils"
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
	fmt.Println("setting up routes for buckets...")
	router.Group(func(r chi.Router) {
		r.Post("/buckets", h.ChangeDB)
		r.Get("/buckets/data", h.listDataSource)
		r.Get("/buckets", h.listBuckets)
		r.Get("/buckets/{bucketName}/count", h.keysCount)
		r.Get("/buckets/{bucketName}/pages/{pageNum}/{pageLen}", h.getPage)
		r.Get("/buckets/{bucketName}/count/{length}", h.countLength)
		r.Get("/buckets/{bucketName}/count/{length}/keys", h.keysCountLength)
		r.Get("/buckets/{bucketName}/keys/{key}", h.lookupByKey)
		r.Get("/buckets/{bucketName}/values/{value}", h.searchByValue)
	})
}

func (h *HttpHandler) ChangeDB(w http.ResponseWriter, r *http.Request) {
	req := &DBRequest{}
	err := utils.ReadJSON(r, req)
	if err != nil {
		utils.WriteErr(w, err, http.StatusBadRequest)
		return
	}
	if err := h.s.ChangeDB(req.Path); err != nil {
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

func (h *HttpHandler) listBuckets(w http.ResponseWriter, r *http.Request) {
	buckets, err := h.s.ListBuckets()
	if err != nil {
		utils.WriteErr(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, buckets)
}

func (h *HttpHandler) keysCount(w http.ResponseWriter, r *http.Request) {
	bucketName := chi.URLParam(r, "bucketName")

	count, err := h.s.KeysCount(bucketName)
	if err != nil {
		utils.WriteErr(w, err, http.StatusInternalServerError)
		return
	}

	response := map[string]uint64{"count": count}

	utils.WriteJSON(w, response)
}

func (h *HttpHandler) getPage(w http.ResponseWriter, r *http.Request) {
	bucketName := chi.URLParam(r, "bucketName")
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

	pages, err := h.s.GetPage(bucketName, pageNum, pageLen)
	if err != nil {
		utils.WriteErr(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, pages)
}

func (h *HttpHandler) countLength(w http.ResponseWriter, r *http.Request) {
	bucketName := chi.URLParam(r, "bucketName")
	length, err := strconv.ParseUint(chi.URLParam(r, "length"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	count, _, err := h.s.KeysCountLength(bucketName, length)
	if err != nil {
		utils.WriteErr(w, err, http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{"count": count}

	utils.WriteJSON(w, response)
}

func (h *HttpHandler) keysCountLength(w http.ResponseWriter, r *http.Request) {
	bucketName := chi.URLParam(r, "bucketName")
	length, err := strconv.ParseUint(chi.URLParam(r, "length"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	count, keys, err := h.s.KeysCountLength(bucketName, length)
	if err != nil {
		utils.WriteErr(w, err, http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{"count": count, "keys": keys}

	utils.WriteJSON(w, response)
}

func (h *HttpHandler) lookupByKey(w http.ResponseWriter, r *http.Request) {
	bucketName := chi.URLParam(r, "bucketName")
	searchKey := chi.URLParam(r, "key")

	foundValue, err := h.s.LookupByKey(bucketName, searchKey)
	if err != nil {
		utils.WriteErr(w, err, http.StatusInternalServerError)
		return
	}

	if foundValue == nil {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}

	response := map[string]string{"value": fmt.Sprintf("%x", foundValue)}

	utils.WriteJSON(w, response)
}

func (h *HttpHandler) searchByValue(w http.ResponseWriter, r *http.Request) {
	bucketName := chi.URLParam(r, "bucketName")
	searchValue := chi.URLParam(r, "value")

	num, err := strconv.ParseUint(searchValue, 16, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	foundKeys, err := h.s.SearchByValue(bucketName, num)
	if err != nil {
		utils.WriteErr(w, err, http.StatusInternalServerError)
		return
	}

	response := map[string][]string{"keys": foundKeys}

	utils.WriteJSON(w, response)
}
