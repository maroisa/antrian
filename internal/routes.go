package internal

import (
	"net/http"
	"path"
)

func (s *Server) Routes() {
	s.mux.HandleFunc("GET /{$}", s.IndexHandler)
	s.mux.HandleFunc("GET /loket/{kode}", s.LoketItemHandler)
	s.mux.HandleFunc("GET /loket", s.LoketHandler)
	s.mux.HandleFunc("GET /loket/ws", s.WSHandler)
	s.mux.Handle("/", assetHandler())
}

func (s *Server) IndexHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, "index.html", nil)
}

func (s *Server) LoketHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, "loket.html", string(s.data.Marshal()))
}

func (s *Server) LoketItemHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, "loket_item.html", map[string]interface{}{
		"loket": r.PathValue("kode"),
		"data":  string(s.data.Marshal()),
	})
}

func assetHandler() http.Handler {
	assetDir := path.Join("web", "assets")
	return http.FileServer(http.Dir(assetDir))
}
