package main

import (
	"net/http"

	// "github.com/bmizerany/pat"
	"github.com/ernazar151020/go-packages/internal/config"
	"github.com/ernazar151020/go-packages/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	// return mux

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteLog)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/majors-suit", handlers.Repo.MajorsSuit)
	mux.Get("/general-quarters", handlers.Repo.Generals)

	mux.Get("/search-availabilty", handlers.Repo.Search)
	mux.Post("/search-availabilty", handlers.Repo.PosetSearch)
	mux.Post("/search-availabilty-json", handlers.Repo.AvailabilityJson)

	mux.Get("/reservation", handlers.Repo.Reservation)
	mux.Post("/reservation", handlers.Repo.PostReservation)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux

}
