package handler

import (
	"antrian/internal/db"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) Routes() {
	s.mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	s.mux.Get("/ws", s.WSHandler)
	s.mux.Post("/login", s.Login)
	s.mux.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		nama := r.FormValue("nama")
		password := r.FormValue("password")

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err)
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}

		err = s.db.CreateUser(r.Context(), db.CreateUserParams{
			Nama:     nama,
			Password: string(hashedPassword),
		})

		if err != nil {
			panic(err)
		}
	})

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
