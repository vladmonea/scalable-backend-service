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
	"time"
)

var (
	welcomeMsg = "Welcome to the graceful server! 💃🏼\n"
	helloMsg   = "hallo fraulain!\n"
)

type gracefulServer struct {
	httpServer *http.Server
	listener   net.Listener
}

func withSimpleLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Default().Printf("Incoming traffic on route %s", r.URL.Path)
		log.Default().Printf("Request: %v", r)
		handler.ServeHTTP(w, r)
	})
}

func withTimer(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		log.Default().Printf("Timing traffic on route %s", r.URL.Path)
		handler.ServeHTTP(w, r)
		defer log.Default().Printf("Request took %d nanoseconds", time.Now().UnixNano()-t1.UnixNano())
	})
}

func (server *gracefulServer) preStart() {
	server.httpServer.Handler = withTimer(withSimpleLogger(server.httpServer.Handler))
}

func handle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, helloMsg)
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, welcomeMsg)
}

func newServer(port string) *gracefulServer {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handle)
	mux.HandleFunc("/greeting", greetingHandler)
	httpServer := &http.Server{Addr: ":" + port, Handler: mux}
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
	var port string

	flag.StringVar(&port, "port", "8080", "./server -port 8080")
	flag.Parse()

	done := make(chan bool, 1)
	interrupts := make(chan os.Signal, 1)
	signal.Notify(interrupts, syscall.SIGTERM, syscall.SIGINT)

	server := newServer(port)

	server.preStart()
	err := server.start()
	if err != nil {
		log.Fatalf("Error starting server - %v\n", err)
	}

	go func() {
		sig := <-interrupts
		log.Default().Printf("Signal intercepted - %v\n", sig)
		server.shutdown()
		done <- true
	}()

	<-done
}
