package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/tedxub2023/global/helper"
	"github.com/tedxub2023/internal/mainevent"
)

type counterHandler struct {
	mainevent mainevent.Service
}

func (h *counterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.handleGetMainEventCounter(w, r)
	default:
		helper.WriteErrorResponse(w, http.StatusMethodNotAllowed, []string{errMethodNotAllowed.Error()})
	}
}

func (h *counterHandler) handleGetMainEventCounter(w http.ResponseWriter, r *http.Request) {
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
			log.Printf("[Main Event HTTP][handleGetMainEventCounter] Failed to get all mainevent. Source: %s, Err: %s\n", source, err.Error())
			helper.WriteErrorResponse(w, statusCode, []string{err.Error()})
			return
		}
		// success
		helper.WriteResponse(w, resBody, statusCode, helper.JSONContentTypeDecorator)
	}()

	// prepare channels for main go routine
	resChan := make(chan int, 1)
	errChan := make(chan error, 1)

	go func() {
		filter, err := ParseGetMainEventsFilter(r.URL.Query())
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

		counter := 0

		for _, t := range res {
			counter += t.JumlahTiket
		}

		resChan <- counter
	}()

	// wait and handle main go routine
	select {
	case <-ctx.Done():
		statusCode = http.StatusGatewayTimeout
		err = errRequestTimeout
	case err = <-errChan:
	case res := <-resChan:
		// construct response data
		resBody, err = json.Marshal(helper.ResponseEnvelope{
			Data: res,
		})
	}
}

func ParseGetMainEventsFilter(request url.Values) (mainevent.GetAllMainEventsFilter, error) {
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
