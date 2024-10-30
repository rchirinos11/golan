package main

import (
	"os"

	"github.com/rchirinos11/golan/serve"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else {
		port = ":" + port
	}
	serve.Run(port)
}
