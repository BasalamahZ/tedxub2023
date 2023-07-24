package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/tedxub2023/cmd/tedxub2023-api-http/config"
)

// Following constants are the possible exit code returned
// when running a server.
const (
	CodeSuccess = iota
	CodeBadConfig
	CodeFailServeHTTP
)

// Run creates a server and starts the server.
//
// Run returns a status code suitable for os.Exit() argument.
func Run() int {
	s, err := new()
	if err != nil {
		return CodeBadConfig
	}

	return s.start()
}

// server is the long-runnning application.
type server struct {
	srv      *http.Server
	handlers []handler
}

// handler provides mechanism to start HTTP handler. All HTTP
// handlers must implements this interface.
type handler interface {
	Start(multiplexer *mux.Router) error
}

// new creates and returns a new server.
func new() (*server, error) {
	s := &server{
		srv: &http.Server{
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}

	// connect to dabatabase
	_, err := sqlx.Connect("postgres", config.BaseConfig())
	if err != nil {
		log.Printf("[tedxub2023-api-http] failed to connect database: %s\n", err.Error())
		return nil, fmt.Errorf("failed to connect database: %s", err.Error())
	}

	return s, nil
}

// start starts the given server.
func (s *server) start() int {
	log.Println("[tedxub2023-api-http] starting server...")

	// create multiplexer object
	rootMux := mux.NewRouter()

	// starts handlers
	for _, h := range s.handlers {
		if err := h.Start(rootMux); err != nil {
			log.Printf("[tedxub2023-api-http] failed to start handler: %s\n", err.Error())
			return CodeFailServeHTTP
		}
	}

	// assign multiplexer as server handler
	s.srv.Handler = rootMux

	// listen and serve
	log.Printf("[tedxub2023-api-http] Server is running at %s:%s", os.Getenv("ADDRESS"), os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("ADDRESS"), os.Getenv("PORT")), rootMux))

	return CodeSuccess
}
