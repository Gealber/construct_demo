package main

import (
	"github.com/Gealber/construct_demo/api"
	"os"
)

func main() {
	port := os.Getenv("SERVER_PORT")
	if len(port) == 0 {
		port = "3000"
	}
	server := api.NewServer()
	server.Run(":" + port)
}
