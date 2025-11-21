package server

import (
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"scalable-backend-service/logger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"go.uber.org/zap"
)

type GracefulServer struct {
	httpServer *http.Server
	listener   net.Listener
}

func (server *GracefulServer) PreStart() error {
	logger := logger.InitLogger()
	if logger == nil {
		errMsg := "failed to initialize logger"
		log.Println(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func NewServer(port string) *GracefulServer {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", SimpleHandler)
	router.Get("/greeting/{name}", GreetingHandler)

	router.Get("/greeting", GreetingHandler)

	router.Get("/users", UserHandler)

	httpServer := &http.Server{Addr: ":" + port, Handler: router}
	return &GracefulServer{httpServer: httpServer}
}

func (server *GracefulServer) Start() (chan bool, error) {
	listener, err := net.Listen("tcp", server.httpServer.Addr)
	if err != nil {
		return nil, err
	}
	server.listener = listener
	go server.httpServer.Serve(server.listener)
	logger.GetLoggerInstance().Info("Server now listening!", zap.String("addresss", server.httpServer.Addr))

	done := make(chan bool, 1)
	interrupts := make(chan os.Signal, 1)
	signal.Notify(interrupts, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-interrupts
		logger.GetLoggerInstance().Warn("Signal intercepted", zap.String("signal", sig.String()))
		server.Shutdown()
		done <- true
	}()

	return done, nil
}

func (server *GracefulServer) Shutdown() error {
	logger.Close()
	if server.listener != nil {
		err := server.listener.Close()
		if err != nil {
			return err
		}
	}
	log.Default().Println("Server shutting down...")
	return nil
}
