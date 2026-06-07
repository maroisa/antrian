package internal

import (
	"log"
	"net/http"
	"sync"

	"github.com/coder/websocket"
)

type Server struct {
	mux     *http.ServeMux
	data    *LoketData
	mu      sync.Mutex
	clients map[*websocket.Conn]bool
}

func NewServer() *Server {
	mux := http.NewServeMux()
	data := &LoketData{}

	s := &Server{
		mux:     mux,
		data:    data,
		clients: make(map[*websocket.Conn]bool),
	}
	s.Routes()
	return s
}

func (s *Server) Start() {
	const PORT = ":3000"
	log.Println("Listening on port " + PORT)
	http.ListenAndServe(PORT, s.mux)
}
