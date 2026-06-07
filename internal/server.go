package internal

import (
	"log"
	"net/http"
	"sync"

	"github.com/coder/websocket"
	"github.com/rs/cors"
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

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{
			"https://antrian.maroisa.com",
			"http://localhost:3000",
		},
		AllowedMethods:   []string{"GET", "POST"},
		AllowCredentials: true,
	}).Handler(s.mux)

	http.ListenAndServe(PORT, handler)
}
