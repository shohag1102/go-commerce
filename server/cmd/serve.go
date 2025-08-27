package cmd

import (
	"fmt"
	"net/http"
	"shohag/github.com/global_router"
	"shohag/github.com/handlers"
	"shohag/github.com/middleware"
)

func Serve() {
	mux := http.NewServeMux() // router

	manager := middleware.NewManager()

	mux.Handle("GET /test", manager.With((http.HandlerFunc(handlers.Test)),
		middleware.Logger,
		middleware.Hudai,
	))

	mux.Handle("GET /products", manager.With((http.HandlerFunc(handlers.GetProducts)),
		middleware.Logger,
		middleware.Hudai,
	))
	mux.Handle("POST /products", manager.With((http.HandlerFunc(handlers.CreateProduct)),
		middleware.Logger,
		middleware.Hudai,
	))
	mux.Handle("GET /products/{id}", manager.With((http.HandlerFunc(handlers.GetProductById)),
		middleware.Logger,
		middleware.Hudai,
	))

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server is running on port: 8080")

	err := http.ListenAndServe(":8080", globalRouter)

	if err != nil {
		fmt.Println("Error starting the server, ", err)
	}
}
