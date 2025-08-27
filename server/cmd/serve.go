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

	// mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	manager := middleware.NewManager()

	mux.Handle("GET /test", manager.With(middleware.Hudai, middleware.Logger)(http.HandlerFunc(handlers.Test)))
	
	mux.Handle("GET /products", manager.With(
		middleware.Hudai, 
		middleware.Logger,
		)(http.HandlerFunc(handlers.GetProducts)))

	// mux.Handle("GET /test", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.Test))))

	// mux.Handle("GET /products", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.GetProducts))))

	mux.Handle("POST /products", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.CreateProduct))))

	mux.Handle("GET /products/{id}", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.GetProductById))))

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server is running on port: 8080")

	err := http.ListenAndServe(":8080", globalRouter)

	if err != nil {
		fmt.Println("Error starting the server, ", err)
	}
}
