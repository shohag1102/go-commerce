package cmd

import (
	"fmt"
	"net/http"
	"shohag/github.com/middleware"
)

func Serve() {
	mux := http.NewServeMux() // router

	manager := middleware.NewManager()

	wrappedMux := manager.WrappedMux(
		mux,
		middleware.Logger,
		middleware.Hudai,
		middleware.CorsWithPreflight,
	)

	initRoutes(mux, manager)

	fmt.Println("Server is running on port: 8080")

	err := http.ListenAndServe(":8080", wrappedMux)

	if err != nil {
		fmt.Println("Error starting the server, ", err)
	}
}
