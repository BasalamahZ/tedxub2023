package http

import (
	"net/http"

	"github.com/tedxub2023/global/helper"
	mainTicket "github.com/tedxub2023/internal/main-ticket"
)

type mainTicketHandler struct {
	mainTicket mainTicket.Service
}

func (h *mainTicketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	default:
		helper.WriteErrorResponse(w, http.StatusMethodNotAllowed, []string{errMethodNotAllowed.Error()})
	}
}
