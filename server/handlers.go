package server

import (
	"io"
	"net/http"
)

var (
	welcomeMsg = "Welcome to the graceful server! 💃🏼\n"
	helloMsg   = "hallo fraulain!\n"
)

func SimpleHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, helloMsg)
}

func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, welcomeMsg)
}
