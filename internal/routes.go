package internal

import (
	"net/http"
)

func (s *Server) Routes() {
	s.mux.HandleFunc("GET /{$}", s.IndexHandler)
	s.mux.HandleFunc("GET /loket/ws", s.WSHandler)
}

func (s *Server) IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
