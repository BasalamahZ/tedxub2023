package http

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/tedxub2023/global/helper"
	"github.com/tedxub2023/internal/ticket"
)

type ticketsHandler struct {
	ticket ticket.Service
}

func (h *ticketsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.handleCreateTicket(w, r)
	case http.MethodPatch:
		h.handleUpdateTicket(w, r)
	default:
		helper.WriteErrorResponse(w, http.StatusMethodNotAllowed, []string{errMethodNotAllowed.Error()})
	}
}

func (h *ticketsHandler) handleCreateTicket(w http.ResponseWriter, r *http.Request) {
	// add timeout to context
	ctx, cancel := context.WithTimeout(r.Context(), 3000*time.Millisecond)
	defer cancel()

	var (
		err        error           // stores error in this handler
		source     string          // stores request source
		resBody    []byte          // stores response body to write
		statusCode = http.StatusOK // stores response status code
	)

	// write response
	defer func() {
		if err != nil {
			log.Printf("[Ticket HTTP][handleCreateTicket] Failed to create tickets. Source: %s, Err: %s\n", source, err.Error())
			helper.WriteErrorResponse(w, statusCode, []string{err.Error()})
			return
		}
		// success
		helper.WriteResponse(w, resBody, statusCode, helper.JSONContentTypeDecorator)
	}()

	// prepare channels for main go routine
	resChan := make(chan string, 1)
	errChan := make(chan error, 1)

	go func() {
		// read body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			statusCode = http.StatusBadRequest
			errChan <- errBadRequest
			return
		}

		// unmarshall body
		request := ticketHTTP{}
		err = json.Unmarshal(body, &request)
		if err != nil {
			statusCode = http.StatusBadRequest
			errChan <- errBadRequest
			return
		}

		// format HTTP request into service object
		reqTicket, err := parseTicketFromCreateRequest(request)
		if err != nil {
			statusCode = http.StatusBadRequest
			errChan <- err
			return
		}

		TicketNama, err := h.ticket.CreateTicket(ctx, reqTicket)
		if err != nil {
			// determine error and status code, by default its internal error
			parsedErr := errInternalServer
			statusCode = http.StatusInternalServerError
			if v, ok := mapHTTPError[err]; ok {
				parsedErr = v
				statusCode = http.StatusBadRequest
			}

			// log the actual error if its internal error
			if statusCode == http.StatusInternalServerError {
				log.Printf("[Ticket HTTP][handleCreateTicket] Internal error from CreateTicket. Err: %s\n", err.Error())
			}

			errChan <- parsedErr
			return
		}

		resChan <- TicketNama
	}()

	// wait and handle main go routine
	select {
	case <-ctx.Done():
		statusCode = http.StatusGatewayTimeout
		err = errRequestTimeout
	case err = <-errChan:
	case ticketNama := <-resChan:
		resBody, err = json.Marshal(helper.ResponseEnvelope{
			Status: "Success",
			Data: map[string]string{
				"nama": ticketNama,
			},
		})
	}
}

// parseTicketFromCreateRequest returns ticket
// from the given HTTP request object.
func parseTicketFromCreateRequest(th ticketHTTP) (ticket.Ticket, error) {
	result := ticket.Ticket{}

	if th.Nama != nil {
		result.Nama = *th.Nama
	}

	if th.NomorIdentitas != nil {
		result.NomorIdentitas = *th.NomorIdentitas
	}

	if th.AsalInstitusi != nil {
		result.AsalInstitusi = *th.AsalInstitusi
	}

	if th.Domisili != nil {
		result.Domisili = *th.Domisili
	}

	if th.Email != nil {
		result.Email = *th.Email
	}

	if th.NomorTelepon != nil {
		result.NomorTelepon = *th.NomorTelepon
	}

	if th.LineID != nil {
		result.LineID = *th.LineID
	}

	if th.Instagram != nil {
		result.Instagram = *th.Instagram
	}

	return result, nil
}

func (h *ticketsHandler) handleUpdateTicket(w http.ResponseWriter, r *http.Request) {
	// add timeout to context
	ctx, cancel := context.WithTimeout(r.Context(), 3000*time.Millisecond)
	defer cancel()

	// create variables to store error and response body
	var (
		err        error
		resBody    []byte
		statusCode = http.StatusOK
	)

	// write response
	defer func() {
		if err != nil {
			log.Printf("[Ticket HTTP][handleUpdateTicket] Failed to update tickets. Err: %s\n", err.Error())
			helper.WriteErrorResponse(w, statusCode, []string{err.Error()})
			return
		}
		resBody, err = json.Marshal(helper.ResponseEnvelope{
			Status: "Success",
		})
		helper.WriteResponse(w, resBody, statusCode, helper.JSONContentTypeDecorator)
	}()

	// prepare channels for main go routine
	errChan := make(chan error, 1)

	// create goroutine to handle main logic
	go func() {
		err := h.ticket.UpdateTicket(ctx)

		if err != nil {
			parsedErr := errInternalServer
			statusCode = http.StatusInternalServerError
			if v, ok := mapHTTPError[err]; ok {
				parsedErr = v
				statusCode = http.StatusBadRequest
			}

			// log the actual error if its internal error
			if statusCode == http.StatusInternalServerError {
				log.Printf("[Ticket HTTP][handleUpdateTicket] Internal error from UpdateTicket. Err: %s\n", err.Error())
			}

			errChan <- parsedErr
			return
		}
	}()

	// wait and handle main go routine
	select {
	case <-ctx.Done():
		statusCode = http.StatusGatewayTimeout
		err = errRequestTimeout
	case err = <-errChan:
	default:
	}
}
