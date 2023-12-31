package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	errUnknownConfig = errors.New("unknown config name")
)

// Handler contains our team HTTP-handlers.
type Handler struct {
	handlers map[string]*handler
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
	// HandlerOutTeams denotes HTTP handler to interact
	// with a our team
	HandlerOutTeams = HandlerIdentity{
		Name: "ourteams",
		URL:  "/ourteams",
	}
)

// New creates a new Handler.
func New(identities []HandlerIdentity) (*Handler, error) {
	h := &Handler{
		handlers: make(map[string]*handler),
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
	case HandlerOutTeams.Name:
		httpHandler = &ourteamsHandler{}
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

type OurTeam struct {
	ID        int        `json:"id"`
	Divisi    string     `json:"divisi"`
	CoreTeam  []CoreTeam `json:"coreteam"`
	Volunteer []string   `json:"volunteer"`
}

type CoreTeam struct {
	Nama      string `json:"nama"`
	Fakultas  string `json:"fakultas"`
	ImageURL  string `json:"image_url"`
	Instagram string `json:"instagram"`
	LinkedIn  string `json:"linkedIn"`
}
