package http

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/tedxub2023/global/helper"
	"github.com/tedxub2023/internal/mainevent"
)

type maineventsHandler struct {
	mainevent mainevent.Service
}

func (h *maineventsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.handleGetAllMainEvents(w, r)
	case http.MethodPost:
		h.handleReplaceMainEventByEmail(w, r)
	default:
		helper.WriteErrorResponse(w, http.StatusMethodNotAllowed, []string{errMethodNotAllowed.Error()})
	}
}

func (h *maineventsHandler) handleGetAllMainEvents(w http.ResponseWriter, r *http.Request) {
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
		// error
		if err != nil {
			log.Printf("[Transaction HTTP][handleGetAllMainEvents] Failed to get all mainevent. Source: %s, Err: %s\n", source, err.Error())
			helper.WriteErrorResponse(w, statusCode, []string{err.Error()})
			return
		}
		// success
		helper.WriteResponse(w, resBody, statusCode, helper.JSONContentTypeDecorator)
	}()

	// prepare channels for main go routine
	resChan := make(chan []mainevent.MainEvent, 1)
	errChan := make(chan error, 1)

	go func() {
		filter, err := parseGetMainEventsFilters(r.URL.Query())
		if err != nil {
			statusCode = http.StatusBadRequest
			errChan <- err
			return
		}
		res, err := h.mainevent.GetAllMainEvents(ctx, filter)
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
				log.Printf("[Match HTTP][handleGetAllGames] Internal error from GetAllMatchs. Err: %s\n", err.Error())
			}

			errChan <- parsedErr
			return
		}

		resChan <- res
	}()

	// wait and handle main go routine
	select {
	case <-ctx.Done():
		statusCode = http.StatusGatewayTimeout
		err = errRequestTimeout
	case err = <-errChan:
	case res := <-resChan:
		// format each transactions
		mainevents := make([]mainEventHTTP, 0)
		for _, r := range res {
			var m mainEventHTTP
			m, err = formatMainEvent(r)
			if err != nil {
				return
			}
			mainevents = append(mainevents, m)
		}

		// construct response data
		resBody, err = json.Marshal(helper.ResponseEnvelope{
			Data: mainevents,
		})
	}
}

func (h *maineventsHandler) handleReplaceMainEventByEmail(w http.ResponseWriter, r *http.Request) {
	// add timeout to context
	ctx, cancel := context.WithTimeout(r.Context(), 5000*time.Millisecond)
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
			log.Printf("[MainEvent HTTP][handleReplaceMainEventByEmail] Failed to replace MainEvents by email. Source: %s, Err: %s\n", source, err.Error())
			helper.WriteErrorResponse(w, statusCode, []string{err.Error()})
			return
		}
		// success
		helper.WriteResponse(w, resBody, statusCode, helper.JSONContentTypeDecorator)
	}()

	// prepare channels for main go routine
	resChan := make(chan int64, 1)
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
		request := mainEventHTTP{}
		err = json.Unmarshal(body, &request)
		if err != nil {
			statusCode = http.StatusBadRequest
			errChan <- errBadRequest
			return
		}

		// format HTTP request into service object
		reqMainEvent, err := parseMainEventFromCreateRequest(request)
		if err != nil {
			statusCode = http.StatusBadRequest
			errChan <- err
			return
		}

		ticketID, err := h.mainevent.ReplaceMainEventByEmail(ctx, reqMainEvent)
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
				log.Printf("[MainEvent HTTP][handleReplaceMainEventByEmail] Internal error from ReplaceMainEventByEmail. Err: %s\n", err.Error())
			}

			errChan <- parsedErr
			return
		}

		resChan <- ticketID
	}()

	// wait and handle main go routine
	select {
	case <-ctx.Done():
		statusCode = http.StatusGatewayTimeout
		err = errRequestTimeout
	case err = <-errChan:
	case ticketID := <-resChan:
		resBody, err = json.Marshal(helper.ResponseEnvelope{
			Status: "Success",
			Data:   ticketID,
		})
	}
}

// parseMainEventFromCreateRequest returns MainEvent
// from the given HTTP request object.
func parseMainEventFromCreateRequest(meh mainEventHTTP) (mainevent.MainEvent, error) {
	result := mainevent.MainEvent{
		Status: mainevent.StatusUnpaid,
		Type:   mainevent.TypeNormalSale,
	}

	if meh.Disabilitas == nil || *meh.Disabilitas == "" {
		result.Disabilitas = mainevent.NoneDisability
	} else {
		disability, err := parseDisability(*meh.Disabilitas)
		if err != nil {
			return mainevent.MainEvent{}, err
		}
		result.Disabilitas = disability
	}

	if meh.Nama != nil {
		result.Nama = *meh.Nama
	}

	if meh.NomorIdentitas != nil {
		result.NomorIdentitas = *meh.NomorIdentitas
	}

	if meh.AsalInstitusi != nil {
		result.AsalInstitusi = *meh.AsalInstitusi
	}

	if meh.Email != nil {
		result.Email = *meh.Email
	}

	if meh.NomorTelepon != nil {
		result.NomorTelepon = *meh.NomorTelepon
	}

	if meh.JumlahTiket != nil {
		result.JumlahTiket = *meh.JumlahTiket
	}

	return result, nil
}

func parseGetMainEventsFilters(request url.Values) (mainevent.GetAllMainEventsFilter, error) {
	result := mainevent.GetAllMainEventsFilter{}

	var types mainevent.Type
	if TypeStr := request.Get("type"); TypeStr != "" {
		parsedType, err := parseType(TypeStr)
		if err != nil {
			return result, errInvalidMainEventType
		}
		types = parsedType
	}

	return mainevent.GetAllMainEventsFilter{
		Type: types,
	}, nil
}
