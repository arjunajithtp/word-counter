package routers

import (
	"github.com/go-chi/chi"
	"github.com/arjunajithtp/word-counter/internal/handlers"
)

func GetRoutes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Get("/home", handlers.HomeHandler)
	mux.Post("/home", handlers.HomeHandler)

	return mux
}
