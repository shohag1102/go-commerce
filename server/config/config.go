package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
}

var configurations Config

func loadConfig() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Env not imported")
		os.Exit(1)
	}
	version := os.Getenv("VERSION")

	if version == "" {
		fmt.Println("Provide valid version")
		os.Exit(1)
	}

	service := os.Getenv("SERVICE_NAME")

	if service == "" {
		fmt.Println("Provide valid service name")
		os.Exit(1)
	}

	portStr := os.Getenv("HTTP_PORT")

	if portStr == "" {
		fmt.Println("Provide valid port")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(portStr, 10, 64)

	if err != nil {
		fmt.Println("Provide valid port number")
		os.Exit(1)
	}

	configurations = Config{
		Version:     version,
		ServiceName: service,
		HttpPort:    int(port),
	}
}

func GetConfig() Config {
	loadConfig()
	return configurations
}
