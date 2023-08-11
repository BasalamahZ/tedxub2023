package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/tedxub2023/internal/ticket"
	tickethttphandler "github.com/tedxub2023/internal/ticket/handler/http"
	ticketservice "github.com/tedxub2023/internal/ticket/service"
	ticketpgstore "github.com/tedxub2023/internal/ticket/store/postgresql"

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
	db, err := sqlx.Connect("postgres", config.BaseConfig())
	if err != nil {
		log.Printf("[tedxub2023-api-http] failed to connect database: %s\n", err.Error())
		return nil, fmt.Errorf("failed to connect database: %s", err.Error())
	}

	// initialize ticket service
	var ticketSvc ticket.Service
	{
		pgStore, err := ticketpgstore.New(db)
		if err != nil {
			log.Printf("[twitter-api-http] failed to initialize ticket postgresql store: %s\n", err.Error())
			return nil, fmt.Errorf("failed to initialize ticket postgresql store: %s", err.Error())
		}

		ticketSvc, err = ticketservice.New(pgStore)
		if err != nil {
			log.Printf("[twitter-api-http] failed to initialize ticket service: %s\n", err.Error())
			return nil, fmt.Errorf("failed to initialize ticket service: %s", err.Error())
		}
	}

	// initialize ticket HTTP handler
	{
		identities := []tickethttphandler.HandlerIdentity{
			tickethttphandler.HandlerTickets,
		}

		ticketHTTP, err := tickethttphandler.New(ticketSvc, identities)
		if err != nil {
			log.Printf("[twitter-api-http] failed to initialize ticket http handlers: %s\n", err.Error())
			return nil, fmt.Errorf("failed to initialize ticket http handlers: %s", err.Error())
		}

		s.handlers = append(s.handlers, ticketHTTP)
	}

	return s, nil
}

// start starts the given server.
func (s *server) start() int {
	log.Println("[tedxub2023-api-http] starting server...")

	// create multiplexer object
	rootMux := mux.NewRouter()
	appMux := rootMux.PathPrefix("/api/v1").Subrouter()

	// starts handlers
	for _, h := range s.handlers {
		if err := h.Start(appMux); err != nil {
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
