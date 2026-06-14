package main

import (
	"antrian/internal/db"
	"antrian/internal/handler"
	"antrian/internal/utils"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
)

func main() {
	db := db.NewConnection("maronn@tcp(127.0.0.1:3306)/antrian")

	secret := utils.GetSecret()
	tokenAuth := jwtauth.New("HS256", []byte(secret), nil)

	srv := handler.NewServer(db, tokenAuth)

	srv.Use(middleware.Logger)
	srv.Use(middleware.Recoverer)
	srv.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowCredentials: true,
	}))

	srv.Start()
}
