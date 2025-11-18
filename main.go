package main

import (
	"log"

	"scalable-backend-service/server"
)

func main() {
	port := "8080"

	server := server.NewServer(port)

	server.PreStart()
	done, err := server.Start()
	if err != nil {
		server.Shutdown()
		log.Fatalf("Error starting server - %v\n", err)
	}
	<-done
}
