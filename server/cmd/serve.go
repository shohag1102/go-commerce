package cmd

import (
	"fmt"
	"net/http"
	"shohag/github.com/global_router"
	"shohag/github.com/middleware"
)

func Serve() {
	mux := http.NewServeMux() // router

	manager := middleware.NewManager()

	manager.Use(
		middleware.Logger,
		middleware.Hudai,
	)

	initRoutes(mux, manager)

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server is running on port: 8080")

	err := http.ListenAndServe(":8080", globalRouter)

	if err != nil {
		fmt.Println("Error starting the server, ", err)
	}
}
