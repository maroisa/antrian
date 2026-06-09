package main

import (
	"antrian/internal/db"
	"antrian/internal/handler"

	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	db := db.NewConnection("rin@tcp(127.0.0.1:3306)/antrian")

	srv := handler.NewServer(db)

	srv.Use(middleware.Logger)
	srv.Use(middleware.Recoverer)

	srv.Start()
}
