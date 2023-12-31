package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/manninen-pythin/bookings/pkg/config"
	"github.com/manninen-pythin/bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	//adding middleware
	mux.Use(middleware.Recoverer)

	//implementing routing for each page
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
