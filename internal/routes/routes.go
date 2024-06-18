package routes

import "github.com/go-chi/chi/v5"

func Master() *chi.Mux {
	r := chi.NewRouter()
	return r
}
