package main

import (
	"os"

	"github.com/rchirinos11/golan/cmd"
)

func main() {
	args := os.Args
	if len(args) == 2 {
		cmd.Execute(args[1])
		return
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else {
		port = ":" + port
	}
	cmd.Run(port)
}
