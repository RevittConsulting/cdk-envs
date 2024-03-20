package ws

import (
	"encoding/json"
	"github.com/RevittConsulting/logger"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

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
	logger.Log().Info("setting up routes for websockets...")
	router.Group(func(r chi.Router) {
		r.Get("/ws", h.handleWebSockets)
	})
}

func (h *Handler) handleWebSockets(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	for {
		chainData, err := h.s.PollChainData()
		if err != nil {
			log.Println("error polling chain data:", err)
			continue
		}

		bytes, err := json.Marshal(chainData)
		if err != nil {
			log.Println("error marshalling chain data:", err)
			continue
		}

		err = ws.WriteMessage(websocket.TextMessage, bytes)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: client disconnected unexpectedly: %v", err)
			} else {
				log.Println("error writing message:", err)
			}
			break
		}
		time.Sleep(1 * time.Second)
	}
}
