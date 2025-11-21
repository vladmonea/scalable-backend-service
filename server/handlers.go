package server

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

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

func UserHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	if name != "" || age != "" {
		age, err := strconv.Atoi(age)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Undefined input param"))
		}
		for _, user := range getUsers() {
			if user.name == name || user.age == age {
				w.WriteHeader(http.StatusOK)
				io.WriteString(w, fmt.Sprintf("Found user with the name %s and age %d\n", user.name, user.age))
				break
			}
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No request params passed\n"))
	}
}
