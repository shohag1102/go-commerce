package cmd

import (
	"net/http"
	"shohag/github.com/handlers"
	"shohag/github.com/middleware"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"GET /test",
		manager.With(
			http.HandlerFunc(handlers.Test),
			middleware.Test,
		),
	)

	mux.Handle(
		"GET /products",
		manager.With(
			http.HandlerFunc(handlers.GetProducts),
		),
	)

	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(handlers.CreateProduct),
		),
	)

	mux.Handle(
		"GET /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.GetProductById),
		),
	)
}
