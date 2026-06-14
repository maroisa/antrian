package handler

import (
	"antrian/internal/db"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (s *Server) Register(w http.ResponseWriter, r *http.Request) {
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
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	nama := r.FormValue("nama")
	password := r.FormValue("password")

	res, err := s.db.GetUser(r.Context(), nama)
	if err != nil {
		log.Println(err)
		http.Error(w, "gagal mendapatkan pengguna", http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(password))
	if err != nil {
		log.Println(err)
		http.Error(w, "nama atau password salah", http.StatusBadRequest)
		return
	}

	_, tokenString, err := s.token.Encode(map[string]interface{}{
		"user_id": res.ID,
	})

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 168),
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

}
