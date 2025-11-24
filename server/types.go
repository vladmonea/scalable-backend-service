package server

import (
	"net"
	"net/http"

	"scalable-backend-service/users"
)

type Request struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ListUserResponse struct {
	Users []users.User `json:"users,omitempty"`
}

type GracefulServer struct {
	httpServer *http.Server
	listener   net.Listener
}
