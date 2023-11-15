package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/tedxub2023/global/helper"
	"github.com/tedxub2023/internal/mainevent"
)

type checkInHandler struct {
	mainevent mainevent.Service
}

func (h *checkInHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Printf("[CheckIn HTTP][checkInHandler] Failed to parse mainevent ID. ID: %s. Err: %s\n", vars["id"], err.Error())
		helper.WriteErrorResponse(w, http.StatusBadRequest, []string{errInvalidMainEventID.Error()})
		return
	}

	switch r.Method {
	case http.MethodPatch:
		h.handleUpdateCheckInStatus(w, r, id)
	default:
		helper.WriteErrorResponse(w, http.StatusMethodNotAllowed, []string{errMethodNotAllowed.Error()})
	}
}

func (h *checkInHandler) handleUpdateCheckInStatus(w http.ResponseWriter, r *http.Request, id int64) {
	ctx, cancel := context.WithTimeout(r.Context(), 2000*time.Millisecond)
	defer cancel()

	var (
		err        error           // stores error in this handler
		source     string          // stores request source
		resBody    []byte          // stores response body to write
		statusCode = http.StatusOK // stores response status code

	)

	resChan := make(chan string, 1)
	errChan := make(chan error, 1)

	defer func() {
		if err != nil {
			log.Printf("[CheckIn HTTP][handleUpdateCheckInStatus] Failed to update check in status. maineventID: %v. Source: %s, Err: %s\n", id, source, err.Error())
			helper.WriteErrorResponse(w, statusCode, []string{err.Error()})
			return
		}
		helper.WriteResponse(w, resBody, statusCode, helper.JSONContentTypeDecorator)
	}()

	go func() {
		ticketNumber, err := parseUpdateCheckInFilter(r.URL.Query())
		if err != nil {
			statusCode = http.StatusBadRequest
			errChan <- err
			return
		}

		res, err := h.mainevent.UpdateCheckInStatus(ctx, id, ticketNumber)
		if err != nil {
			parsedErr := errInternalServer
			statusCode = http.StatusInternalServerError
			if v, ok := mapHTTPError[err]; ok {
				parsedErr = v
				statusCode = http.StatusBadRequest
			}
			errChan <- parsedErr
		}
		resChan <- res
	}()

	select {
	case <-ctx.Done():
		err = errRequestTimeout
		statusCode = http.StatusGatewayTimeout
		return
	case err = <-errChan:
	case res := <-resChan:
		resBody, err = json.Marshal(helper.ResponseEnvelope{
			Status: "Success",
			Data:   res,
		})

		if err != nil {
			err = errFailedJSONMarshal
			statusCode = http.StatusBadRequest
			return
		}
	}
}

func parseUpdateCheckInFilter(url url.Values) (string, error) {
	var ticketNumber string
	if ticketNumber = url.Get("ticket_number"); ticketNumber == "" {
		return "", errInvalidTicketNumber
	}
	return ticketNumber, nil
}
