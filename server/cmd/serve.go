package cmd

import (
	"fmt"
	"net/http"
	"shohag/github.com/global_router"
	"shohag/github.com/handlers"
)

func Serve() {
	mux := http.NewServeMux() // router

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))

	mux.Handle("POST /create-products", http.HandlerFunc(handlers.CreateProduct))

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server is running on port: 8080")

	err := http.ListenAndServe(":8080", globalRouter)

	if err != nil {
		fmt.Println("Error starting the server, ", err)
	}
}
