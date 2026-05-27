package main

import (
	"os"
	"github.com/ondrejsika/counter-frontend-go/pkg/server"
)

func main() {
	os.Setenv("FONT_COLOR", "white")
	os.Setenv("BACKGROUND_COLOR", "#EC018D")
	server.Server()
}
