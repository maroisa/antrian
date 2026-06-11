package handler

import (
	"antrian/internal/db"
	"antrian/internal/utils"
	"log"
	"net/http"
	"sync"

	"github.com/coder/websocket"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	mux     *chi.Mux
	db      *db.Queries
	mu      sync.Mutex
	clients map[*websocket.Conn]bool
}

func NewServer(db *db.Queries) *Server {
	return &Server{
		mux:     chi.NewRouter(),
		db:      db,
		clients: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) Start() {
	port := utils.GetPort()

	s.Routes()

	log.Println("Listening on http://127.0.0.1" + port)
	http.ListenAndServe(port, s.mux)
}

func (s *Server) Use(middleware func(http.Handler) http.Handler) {
	s.mux.Use(middleware)
}
