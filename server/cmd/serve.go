package cmd

import (
	"fmt"
	"net/http"
	"shohag/github.com/middleware"
)

func Serve() {

	manager := middleware.NewManager()

	manager.Use(
		middleware.Cors,
		middleware.Preflight,
		middleware.Logger,
	)
	mux := http.NewServeMux() // router

	wrappedMux := manager.WrappedMux(mux)

	initRoutes(mux, manager)

	fmt.Println("Server is running on port: 8080")

	err := http.ListenAndServe(":8080", wrappedMux)

	if err != nil {
		fmt.Println("Error starting the server, ", err)
	}
}
