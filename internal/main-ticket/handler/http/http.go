package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	mainTicket "github.com/tedxub2023/internal/main-ticket"
)

var (
	errUnknownConfig = errors.New("unknown config name")
)

// dateFormat denotes the standard date format used in
// transaction HTTP request and response.
var dateFormat = "02-01-2006"

// Handler contains transaction HTTP-handlers.
type Handler struct {
	handlers    map[string]*handler
	transaction mainTicket.Service
}

// handler is the HTTP handler wrapper.
type handler struct {
	h        http.Handler
	identity HandlerIdentity
}

// HandlerIdentity denotes the identity of an HTTP hanlder.
type HandlerIdentity struct {
	Name string
	URL  string
}

// Followings are the known HTTP handler identities
var ()

// New creates a new Handler.
func New(transaction mainTicket.Service, identities []HandlerIdentity) (*Handler, error) {
	h := &Handler{
		handlers:    make(map[string]*handler),
		transaction: transaction,
	}

	// apply options
	for _, identity := range identities {
		if h.handlers == nil {
			h.handlers = map[string]*handler{}
		}

		h.handlers[identity.Name] = &handler{
			identity: identity,
		}

		handler, err := h.createHTTPHandler(identity.Name)
		if err != nil {
			return nil, err
		}

		h.handlers[identity.Name].h = handler
	}

	return h, nil
}

// createHTTPHandler creates a new HTTP handler that
// implements http.Handler.
func (h *Handler) createHTTPHandler(configName string) (http.Handler, error) {
	var httpHandler http.Handler
	switch configName {

	default:
		return httpHandler, errUnknownConfig
	}
	return httpHandler, nil
}

// Start starts all HTTP handlers.
func (h *Handler) Start(multiplexer *mux.Router) error {
	for _, handler := range h.handlers {
		multiplexer.Handle(handler.identity.URL, handler.h)
	}
	return nil
}

type mainTicketHTTP struct {
}
