package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (s *Server) Routes() {
	s.mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	s.mux.Get("/ws", s.WSHandler)

	s.mux.Get("/loket/{id:[0-9]+}/baru", s.NewAntrian)
	s.mux.Get("/loket/{id:[0-9]+}", s.GetAntrian)
	s.mux.Get("/antrian/{id:[0-9]+}/selesai", s.AntrianSelesai)
	s.mux.Get("/antrian/{id:[0-9]+}/panggil", s.AntrianDipanggil)
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

func (s *Server) AntrianDipanggil(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	paramID, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	antrian, err := s.db.GetAntrian(r.Context(), int32(paramID))
	if err != nil {
		panic(err)
	}

	allLoket, err := s.db.GetAllLoket(r.Context())

	s.Broadcast(r.Context(), map[string]interface{}{
		"data":       allLoket,
		"terpanggil": antrian,
	})
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
