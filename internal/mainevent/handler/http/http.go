package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tedxub2023/internal/mainevent"
)

var (
	errUnknownConfig = errors.New("unknown config name")
)

// Handler contains mainevent HTTP-handlers.
type Handler struct {
	handlers  map[string]*handler
	mainevent mainevent.Service
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
	// HandlerMainEvent denotes HTTP handler to interact
	// with a mainevent
	HandlerMainEvent = HandlerIdentity{
		Name: "mainevent",
		URL:  "/mainevents/{id}",
	}

	// HandlerMainEvents denotes HTTP handler to interact
	// with mainevents
	HandlerMainEvents = HandlerIdentity{
		Name: "mainevents",
		URL:  "/mainevents",
	}

	// HandlerCounter denotes HTTP handler to interact
	// with Counter
	HandlerCounter = HandlerIdentity{
		Name: "counter",
		URL:  "/counter",
	}

	// HandlerCheckIn update checkin status
	// after participant scan ticket (enter event)
	HandlerCheckIn = HandlerIdentity{
		Name: "checkin",
		URL:  "/checkin/mainevent/{id}",
	}
)

// New creates a new Handler.
func New(mainevent mainevent.Service, identities []HandlerIdentity) (*Handler, error) {
	h := &Handler{
		handlers:  make(map[string]*handler),
		mainevent: mainevent,
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
	case HandlerMainEvent.Name:
		httpHandler = &maineventHandler{
			mainevent: h.mainevent,
		}
	case HandlerMainEvents.Name:
		httpHandler = &maineventsHandler{
			mainevent: h.mainevent,
		}
	case HandlerCounter.Name:
		httpHandler = &counterHandler{
			mainevent: h.mainevent,
		}
	case HandlerCheckIn.Name:
		httpHandler = &checkInHandler{
			mainevent: h.mainevent,
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

type mainEventHTTP struct {
	ID                *int64    `json:"id"`
	Nama              *string   `json:"nama"`
	Disabilitas       *string   `json:"disabilitas"`
	NomorIdentitas    *string   `json:"nomor_identitas"`
	AsalInstitusi     *string   `json:"asal_institusi"`
	Email             *string   `json:"email"`
	NomorTelepon      *string   `json:"nomor_telepon"`
	JumlahTiket       *int      `json:"jumlah_tiket"`
	TotalHarga        *int64    `json:"total_harga"`
	OrderID           *string   `json:"order_id"`
	Type              *string   `json:"type"`
	Status            *string   `json:"status"`
	ImageURI          *string   `json:"image_uri"`
	NomorTiket        *[]string `json:"nomor_tiket"`
	CheckInStatus     *bool     `json:"checkin_status"`
	CheckInNomorTiket *[]string `json:"checkin_nomor_tiket"`
}
