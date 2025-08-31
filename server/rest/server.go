package rest

import (
	"fmt"
	"net/http"
	"shohag/github.com/config"
	"shohag/github.com/rest/middleware"
	"strconv"
)

func Start(cnf config.Config) {

	manager := middleware.NewManager()

	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)
	mux := http.NewServeMux() // router

	wrappedMux := manager.WrappedMux(mux)

	initRoutes(mux, manager)

	addr := ":" + strconv.Itoa(cnf.HttpPort)

	fmt.Println("Server is running on port: ", addr)

	err := http.ListenAndServe(addr, wrappedMux)

	if err != nil {
		fmt.Println("Error starting the server, ", err)
	}
}
