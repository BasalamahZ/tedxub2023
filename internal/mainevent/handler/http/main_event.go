package http

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/tedxub2023/global/helper"
	"github.com/tedxub2023/internal/mainevent"
)

type maineventHandler struct {
	mainevent mainevent.Service
}

func (h *maineventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	maineventID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Printf("[MainEvent HTTP][maineventHandler] Failed to parse mainevent ID. ID: %s. Err: %s\n", vars["id"], err.Error())
		helper.WriteErrorResponse(w, http.StatusBadRequest, []string{errInvalidMainEventID.Error()})
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.handleGetMainEventByID(w, r, int64(maineventID))
	case http.MethodPatch:
		h.handleUpdatePaymentStatus(w, r, int64(maineventID))
	default:
		helper.WriteErrorResponse(w, http.StatusMethodNotAllowed, []string{errMethodNotAllowed.Error()})
	}
}

func (h *maineventHandler) handleGetMainEventByID(w http.ResponseWriter, r *http.Request, maineventID int64) {
	// add timeout to context
	ctx, cancel := context.WithTimeout(r.Context(), 2000*time.Millisecond)
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
			log.Printf("[MainEvent HTTP][handleGetMainEventByID] Failed to get mainevent. maineventID: %v. Source: %s, Err: %s\n", maineventID, source, err.Error())
			helper.WriteErrorResponse(w, statusCode, []string{err.Error()})
			return
		}
		// success
		helper.WriteResponse(w, resBody, statusCode, helper.JSONContentTypeDecorator)
	}()

	// prepare channels for main go routine
	resChan := make(chan mainevent.MainEvent, 1)
	errChan := make(chan error, 1)

	go func() {
		// parsed filter
		nomorTiket, err := parseGetMainEventFilter(r.URL.Query())
		if err != nil {
			statusCode = http.StatusBadRequest
			errChan <- err
			return
		}

		// get result MainEvent
		res, err := h.mainevent.GetMainEventByID(ctx, maineventID, nomorTiket)
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
				log.Printf("[MainEvent HTTP][handleGetMainEventByID] Internal error from GetMainEventByID. maineventID: %v. Err: %s\n", maineventID, err.Error())
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
		// format user
		var m mainEventHTTP
		m, err = formatMainEvent(res)
		if err != nil {
			return
		}

		// construct response data
		resBody, err = json.Marshal(helper.ResponseEnvelope{
			Data: m,
		})
	}
}

func parseGetMainEventFilter(request url.Values) (string, error) {
	var nomor_tiket string
	if queryNomorTIket := request.Get("nomor_tiket"); queryNomorTIket != "" {
		nomor_tiket = queryNomorTIket
	}

	return nomor_tiket, nil
}

func (h *maineventHandler) handleUpdatePaymentStatus(w http.ResponseWriter, r *http.Request, maineventID int64) {
	ctx, cancel := context.WithTimeout(r.Context(), 10000*time.Millisecond)
	defer cancel()

	var (
		err        error
		source     string
		resBody    []byte
		statusCode = http.StatusOK
	)

	defer func() {
		if err != nil {
			log.Printf("[MainEvent HTTP][handleUpdatePaymentStatus] Failed update mainevent. maineventID: %v. Source: %s, Err: %s\n", maineventID, source, err.Error())
			helper.WriteErrorResponse(w, statusCode, []string{err.Error()})
			return
		}
		helper.WriteResponse(w, resBody, statusCode, helper.JSONContentTypeDecorator)
	}()

	errChan := make(chan error, 1)
	resChan := make(chan int64, 1)

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

		// get current result MainEvent
		current, err := h.mainevent.GetMainEventByID(ctx, maineventID, "")
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
				log.Printf("[MainEvent HTTP][handleGetMainEventByID] Internal error from GetMainEventByID. maineventID: %v. Err: %s\n", maineventID, err.Error())
			}

			errChan <- parsedErr
			return
		}

		// format HTTP request into service object
		reqMainEvent, err := parseMainEventFromUpdateRequest(request, current)
		if err != nil {
			statusCode = http.StatusBadRequest
			errChan <- err
			return
		}

		err = h.mainevent.UpdatePaymentStatus(ctx, reqMainEvent)
		if err != nil {
			parsedErr := errInternalServer
			statusCode = http.StatusInternalServerError
			if v, ok := mapHTTPError[err]; ok {
				parsedErr = v
				statusCode = http.StatusBadRequest
			}
			if statusCode == http.StatusInternalServerError {
				log.Printf("[MainEvent HTTP][handleUpdatePaymentStatus] Internal error from UpdatePaymentStatus. maineventID: %v. Err: %s\n", maineventID, err.Error())
			}
			errChan <- parsedErr
			return
		}
		resChan <- maineventID
	}()

	select {
	case <-ctx.Done():
		err = errRequestTimeout
		statusCode = http.StatusGatewayTimeout
	case err = <-errChan:
	case maineventID := <-resChan:
		resBody, err = json.Marshal(helper.ResponseEnvelope{
			Status: "Success",
			Data:   maineventID,
		})
	}
}

// parseMainEventFromUpdateRequest returns MainEvent
// from the given HTTP request object.
func parseMainEventFromUpdateRequest(meh mainEventHTTP, current mainevent.MainEvent) (mainevent.MainEvent, error) {
	result := current

	if meh.ImageURI != nil {
		result.ImageURI = *meh.ImageURI
	}

	if meh.Status != nil {
		status, err := parseStatus(*meh.Status)
		if err != nil {
			return mainevent.MainEvent{}, err
		}

		result.Status = status
	}

	if meh.FileURI != nil {
		result.FileURI = *meh.FileURI
	}

	return result, nil
}
