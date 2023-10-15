package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tedxub2023/internal/transaction"
)

var (
	errUnknownConfig = errors.New("unknown config name")
)

// Handler contains transaction HTTP-handlers.
type Handler struct {
	handlers    map[string]*handler
	transaction transaction.Service
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
	// HandlerTransaction denotes HTTP handler to interact
	// with a transaction
	HandlerTransaction = HandlerIdentity{
		Name: "transaction",
		URL:  "/transactions/{id}",
	}

	// HandlerTransactions denotes HTTP handler to interact
	// with transactions
	HandlerTransactions = HandlerIdentity{
		Name: "transactions",
		URL:  "/transactions",
	}
)

// New creates a new Handler.
func New(transaction transaction.Service, identities []HandlerIdentity) (*Handler, error) {
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
	case HandlerTransaction.Name:
		httpHandler = &transactionHandler{
			transaction: h.transaction,
		}
	case HandlerTransactions.Name:
		httpHandler = &transactionsHandler{
			transaction: h.transaction,
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

type transactionHTTP struct {
	ID               *int64    `json:"id"`
	Nama             *string   `json:"nama"`
	JenisKelamin     *string   `json:"jenis_kelamin"`
	NomorIdentitas   *string   `json:"nomor_identitas"`
	AsalInstitusi    *string   `json:"asal_institusi"`
	Domisili         *string   `json:"domisili"`
	Email            *string   `json:"email"`
	NomorTelepon     *string   `json:"nomor_telepon"`
	LineID           *string   `json:"line_id"`
	Instagram        *string   `json:"instagram"`
	JumlahTiket      *int      `json:"jumlah_tiket"`
	Harga            *int64    `json:"harga"`
	NomorTiket       *[]string `json:"nomor_tiket"`
	ResponseMidtrans *string   `json:"response_midtrans"`
	CheckInStatus    *bool     `json:"checkin_status"`
	CheckInCounter   *int      `json:"checkin_counter"`
}
