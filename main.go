package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming request: %v", r)
	io.WriteString(w, "hallo fraulain!\n")
}

func main() {
	port := 8080

	http.HandleFunc("/", handle)

	log.Printf("Starting server on localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
