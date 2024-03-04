package main

import (
	"log"
	"stepp-backend/src/server"
)

func main() {
	if err := server.RunServer(); err != nil {
		log.Fatal("Error cannot run server")
	}
}
