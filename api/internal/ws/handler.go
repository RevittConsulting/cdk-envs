package ws

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
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
	fmt.Println("setting up routes for websockets...")
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
		block, err := h.s.PollChainData()
		if err != nil {
			log.Println("error polling chain data:", err)
			break
		}

		err = ws.WriteMessage(websocket.TextMessage, []byte(strconv.FormatUint(block, 10)))
		if err != nil {
			// handle error
			break
		}
		time.Sleep(1 * time.Second)
	}

	//for {
	//	_, message, err := ws.ReadMessage()
	//	if err != nil {
	//		log.Println("read:", err)
	//		break
	//	}
	//	log.Printf("received: %s", message)
	//
	//	response, err := h.s.ProcessData(message)
	//	if err != nil {
	//		log.Println("error processing data:", err)
	//		break
	//	}
	//
	//	err = ws.WriteMessage(websocket.TextMessage, response)
	//	if err != nil {
	//		log.Println("write:", err)
	//		break
	//	}
	//}
}
