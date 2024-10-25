package router

import (
	"bookstore/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(h handler.Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.DefaultLogger)

	r.Route("/user", func(r chi.Router) {
		r.Post("/", h.CreateUser)
	})

	r.Route("/book", func(r chi.Router) {
		r.Get("/", h.ListBooks)
	})

	r.Route("/order", func(r chi.Router) {
		r.Post("/", h.CreateOrder)
		r.Get("/history", h.GetOrderHistory)
	})

	return r
}
