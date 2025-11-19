package server

import (
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var (
	welcomeMsg = "Welcome to the graceful server! 💃🏼\n"
	helloMsg   = "hallo fraulain!\n"
)

func SimpleHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, helloMsg)
}

func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	greeting := welcomeMsg
	if name != "" {
		greeting = "Hello " + name + "!\n" + welcomeMsg
	} else {
		name = r.URL.Query().Get("name")
		if name != "" {
			greeting = "Hello " + name + "!\n" + welcomeMsg
		}
	}
	io.WriteString(w, greeting)
}
