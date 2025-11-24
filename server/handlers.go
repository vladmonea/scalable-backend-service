package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"scalable-backend-service/users"

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

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	if name != "" || age != "" {
		age, err := strconv.Atoi(age)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Undefined input param"))
		}
		for _, user := range users.GetUsers() {
			if user.Name == name || user.Age == age {
				w.WriteHeader(http.StatusOK)
				io.WriteString(w, fmt.Sprintf("Found user with the name %s and age %d\n", user.Name, user.Age))
				break
			}
		}
	} else {
		response := ListUserResponse{Users: users.GetUsers()}
		data, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var request Request
	err = json.Unmarshal(data, &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	users.AddUser(request.Name, request.Age)
	w.WriteHeader(http.StatusOK)
}
