package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/http"
)

type ErrResponse struct {
	Error string `json:"error"`
}

type Response struct {
	Data      interface{} `json:"data"`
	FromCache bool        `json:"from_cache"`
}

func WriteJSON(w http.ResponseWriter, v interface{}) {
	if v == nil {
		WriteErr(w, fmt.Errorf("not found"), http.StatusNotFound)
		return
	}

	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetIndent("", "\t")

	if err := encoder.Encode(v); err != nil {
		WriteErr(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	_, err := w.Write(buffer.Bytes())
	if err != nil {
		log.Fatal("error writing JSON to response:", zap.Error(err))
	}
}

func WriteErr(w http.ResponseWriter, err error, code int) {
	if err == nil {
		WriteErr(w, fmt.Errorf(http.StatusText(code)), code)
		return
	}

	switch err.Error() {
	case "unauthorised":
		w.WriteHeader(http.StatusUnauthorized)
		WriteJSON(w, ErrResponse{Error: err.Error()})
	default:
		w.WriteHeader(code)
		WriteJSON(w, ErrResponse{Error: err.Error()})
	}
}
