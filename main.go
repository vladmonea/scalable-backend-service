package main

import (
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type gracefulServer struct {
	httpServer *http.Server
	listener   net.Listener
}

func handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming request: %v", r)
	io.WriteString(w, "hallo fraulain!\n")
}

func newServer(port string) *gracefulServer {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handle)
	httpServer := &http.Server{Addr: port, Handler: mux}
	return &gracefulServer{httpServer: httpServer}
}

func (server *gracefulServer) start() error {
	listener, err := net.Listen("tcp", server.httpServer.Addr)
	if err != nil {
		return err
	}
	server.listener = listener
	go server.httpServer.Serve(server.listener)
	log.Default().Printf("Server now listening on %s\n", server.httpServer.Addr)
	return nil
}

func (server *gracefulServer) shutdown() error {
	if server.listener != nil {
		err := server.listener.Close()
		if err != nil {
			return err
		}
	}
	log.Default().Println("Server shutting down...")
	return nil
}

func main() {
	port := "8080"

	flag.StringVar(&port, "port", "8080", "./server -port 8080")
	flag.Parse()

	done := make(chan bool, 1)
	interrupts := make(chan os.Signal, 1)
	signal.Notify(interrupts, syscall.SIGTERM, syscall.SIGINT)

	server := newServer(port)

	err := server.start()
	if err != nil {
		log.Fatalf("Error starting server - %v\n", err)
	}

	go func() {
		sig := <-interrupts
		log.Default().Printf("Received shutdown signal - %v\n", sig)
		server.shutdown()
		done <- true
	}()

	<-done
}
