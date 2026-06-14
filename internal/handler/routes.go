package handler

import (
	"antrian/internal/db"
	"antrian/web"
	"io"
	"io/fs"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
)

func (s *Server) Routes() {
	s.mux.Get("/ws", s.WSHandler)
	s.mux.Post("/login", s.Login)
	s.mux.Post("/register", s.Register)

	s.mux.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(s.token))
		r.Use(jwtauth.Authenticator(s.token))

		r.Get("/auth", func(w http.ResponseWriter, r *http.Request) {})
		r.Get("/loket/baru", s.NewAntrian)
		r.Get("/antrian/minta", s.MintaAntrian)
		r.Get("/antrian/ambil/{loketID}/{antrianID}", s.AmbilAntrian)
		r.Get("/loket/{id:[0-9]+}", s.GetAntrian)
		r.Get("/antrian/{id:[0-9]+}/selesai", s.AntrianSelesai)
	})

	distFS, err := fs.Sub(web.EmbeddedFiles, "dist")
	if err != nil {
		panic("Gagal memuat frontend")
	}

	fileServer := http.FileServer(http.FS(distFS))
	s.mux.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/")

		_, err := distFS.Open(path)
		if os.IsNotExist(err) && path != "" {
			indexFile, _ := distFS.Open("index.html")
			defer indexFile.Close()

			seeker, ok := indexFile.(io.ReadSeeker)
			if !ok {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			stat, _ := indexFile.Stat()

			http.ServeContent(w, r, "index.html", stat.ModTime(), seeker)
			return
		}

		fileServer.ServeHTTP(w, r)
	})
}

func (s *Server) GetAntrian(w http.ResponseWriter, r *http.Request) {
	loketID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		panic(err)
	}

	antrian, err := s.db.ListAntrian(r.Context(), int32(loketID))
	if err != nil {
		panic(err)
	}

	render.JSON(w, r, antrian)
}

func (s *Server) NewAntrian(w http.ResponseWriter, r *http.Request) {
	res, err := s.db.InsertAntrian(r.Context())
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	data, err := s.db.GetAntrian(r.Context(), int32(id))
	if err != nil {
		panic(err)
	}

	render.JSON(w, r, data)
}

func (s *Server) AntrianSelesai(w http.ResponseWriter, r *http.Request) {
	paramID := chi.URLParam(r, "id")
	antrianID, err := strconv.Atoi(paramID)
	if err != nil {
		panic(err)
	}

	s.db.SelesaiAntrian(r.Context(), int32(antrianID))

	render.JSON(w, r, map[string]string{
		"message": "antrian " + paramID + " telah selesai",
	})
}

func (s *Server) MintaAntrian(w http.ResponseWriter, r *http.Request) {
	antrian, err := s.db.MintaAntrian(r.Context())
	if err != nil {
		panic(err)
	}

	render.JSON(w, r, antrian)
}

func (s *Server) AmbilAntrian(w http.ResponseWriter, r *http.Request) {
	loketID, _ := strconv.Atoi(chi.URLParam(r, "loketID"))
	antrianID, _ := strconv.Atoi(chi.URLParam(r, "antrianID"))

	err := s.db.AmbilAntrian(r.Context(), db.AmbilAntrianParams{
		ID:    int32(antrianID),
		Loket: int32(loketID),
	})

	if err != nil {
		panic(err)
	}

	render.JSON(w, r, nil)
}
