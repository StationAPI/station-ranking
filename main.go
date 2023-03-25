package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/stationapi/station-ranking/db"
	"github.com/stationapi/station-ranking/routes"
)

func main() {
	db, err := db.Connect()

	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)

	r.Get("/ranking/new", func(w http.ResponseWriter, r *http.Request) {
    err := routes.New(w, r, db)

    if err != nil {
      fmt.Println(err)
    }
	})

	r.Get("/ranking/recently-bumped", func(w http.ResponseWriter, r *http.Request) {
		routes.RecentlyBumped(w, r, db)
	})

	r.Get("/ranking/hot", func(w http.ResponseWriter, r *http.Request) {
		routes.Hot(w, r, db)
	})
}
