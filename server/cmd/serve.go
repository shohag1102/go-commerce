package cmd

import (
	"shohag/github.com/config"
	"shohag/github.com/rest"
)

func Serve() {
	cnf := config.GetConfig()

	rest.Start(cnf)
}
