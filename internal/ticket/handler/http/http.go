package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tedxub2023/internal/ticket"
)

var (
	errUnknownConfig = errors.New("unknown config name")
)

// Handler contains ticket HTTP-handlers.
type Handler struct {
	handlers map[string]*handler
	ticket   ticket.Service
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
var (
	// HandlerTickets denotes HTTP handler to interact
	// with a ticket
	HandlerTickets = HandlerIdentity{
		Name: "tickets",
		URL:  "/tickets",
	}
)

// New creates a new Handler.
func New(ticket ticket.Service, identities []HandlerIdentity) (*Handler, error) {
	h := &Handler{
		handlers: make(map[string]*handler),
		ticket:   ticket,
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
	case HandlerTickets.Name:
		httpHandler = &ticketsHandler{
			ticket: h.ticket,
		}
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

type ticketHTTP struct {
	ID             *int64  `json:"id"`
	Nama           *string `json:"nama"`
	NomorIdentitas *string `json:"nomor_identitas"`
	AsalInstitusi  *string `json:"asal_institusi"`
	Domisili       *string `json:"domisili"`
	Email          *string `json:"email"`
	NomorTelepon   *string `json:"nomor_telepon"`
	LineID         *string `json:"line_id"`
	Instagram      *string `json:"instagram"`
}
