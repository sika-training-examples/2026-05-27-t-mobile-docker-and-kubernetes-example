package main

import (
	"os"
	"github.com/ondrejsika/counter/pkg/server"
)

func main() {
	os.Setenv("EXTRA_TEXT", "Hello T-Mobile")
	os.Setenv("API_ONLY", "1")
	server.Server(server.ServerOptions{})
}
