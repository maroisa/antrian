package main

import (
	"antrian/internal/db"
	"antrian/internal/handler"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	db := db.NewConnection("rin@tcp(127.0.0.1:3306)/antrian")

	srv := handler.NewServer(db)

	srv.Use(middleware.Logger)
	srv.Use(middleware.Recoverer)
	srv.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://localhost:5173"},
		AllowedMethods: []string{"GET"},
	}))

	srv.Start()
}
