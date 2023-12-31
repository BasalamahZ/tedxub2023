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
	"github.com/tedxub2023/internal/transaction"
)

type transactionHandler struct {
	transaction transaction.Service
}

func (h *transactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	transactionID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Printf("[Transaction HTTP][transactionHandler] Failed to parse transaction ID. ID: %s. Err: %s\n", vars["id"], err.Error())
		helper.WriteErrorResponse(w, http.StatusBadRequest, []string{errInvalidTransactionID.Error()})
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.handleGetTransactionByID(w, r, int64(transactionID))
	case http.MethodPatch:
		h.handleUpdatePaymentStatus(w, r, int64(transactionID))
	default:
		helper.WriteErrorResponse(w, http.StatusMethodNotAllowed, []string{errMethodNotAllowed.Error()})
	}
}

func (h *transactionHandler) handleGetTransactionByID(w http.ResponseWriter, r *http.Request, transactionID int64) {
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
			log.Printf("[Transaction HTTP][handleGetTransactionByID] Failed to get transaction. transactionID: %v. Source: %s, Err: %s\n", transactionID, source, err.Error())
			helper.WriteErrorResponse(w, statusCode, []string{err.Error()})
			return
		}
		// success
		helper.WriteResponse(w, resBody, statusCode, helper.JSONContentTypeDecorator)
	}()

	// prepare channels for main go routine
	resChan := make(chan transaction.Transaction, 1)
	errChan := make(chan error, 1)

	go func() {
		// parsed filter
		nomorTiket, err := parseGetTransactionFilter(r.URL.Query())
		if err != nil {
			statusCode = http.StatusBadRequest
			errChan <- err
			return
		}

		// get result transaction
		res, err := h.transaction.GetTransactionByID(ctx, transactionID, nomorTiket)
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
				log.Printf("[Transaction HTTP][handleGetTransactionByID] Internal error from GetTransactionByID. transactionID: %v. Err: %s\n", transactionID, err.Error())
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
		var t transactionHTTP
		t, err = formatTransaction(res)
		if err != nil {
			return
		}

		// construct response data
		resBody, err = json.Marshal(helper.ResponseEnvelope{
			Data: t,
		})
	}
}

func parseGetTransactionFilter(request url.Values) (string, error) {
	var nomor_tiket string
	if queryNomorTIket := request.Get("nomor_tiket"); queryNomorTIket != "" {
		nomor_tiket = queryNomorTIket
	}

	return nomor_tiket, nil
}

func (h *transactionHandler) handleUpdatePaymentStatus(w http.ResponseWriter, r *http.Request, transactionID int64) {
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
			log.Printf("[Transaction HTTP][handleUpdatePaymentStatus] Failed update transaction. transactionID: %v. Source: %s, Err: %s\n", transactionID, source, err.Error())
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
		request := transactionHTTP{}
		err = json.Unmarshal(body, &request)
		if err != nil {
			statusCode = http.StatusBadRequest
			errChan <- errBadRequest
			return
		}

		// get current result transaction
		current, err := h.transaction.GetTransactionByID(ctx, transactionID, "")
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
				log.Printf("[Transaction HTTP][handleGetTransactionByID] Internal error from GetTransactionByID. transactionID: %v. Err: %s\n", transactionID, err.Error())
			}

			errChan <- parsedErr
			return
		}

		// format HTTP request into service object
		reqTransaction, err := parseTransactionFromUpdateRequest(request, current)
		if err != nil {
			statusCode = http.StatusBadRequest
			errChan <- err
			return
		}

		err = h.transaction.UpdatePaymentStatus(ctx, reqTransaction)
		if err != nil {
			parsedErr := errInternalServer
			statusCode = http.StatusInternalServerError
			if v, ok := mapHTTPError[err]; ok {
				parsedErr = v
				statusCode = http.StatusBadRequest
			}
			if statusCode == http.StatusInternalServerError {
				log.Printf("[Transaction HTTP][handleUpdatePaymentStatus] Internal error from UpdatePaymentStatus. transactionID: %v. Err: %s\n", transactionID, err.Error())
			}
			errChan <- parsedErr
			return
		}
		resChan <- transactionID
	}()

	select {
	case <-ctx.Done():
		err = errRequestTimeout
		statusCode = http.StatusGatewayTimeout
	case err = <-errChan:
	case transactionID := <-resChan:
		resBody, err = json.Marshal(helper.ResponseEnvelope{
			Status: "Success",
			Data:   transactionID,
		})
	}
}

// parseTransactionFromUpdateRequest returns transaction
// from the given HTTP request object.
func parseTransactionFromUpdateRequest(th transactionHTTP, current transaction.Transaction) (transaction.Transaction, error) {
	result := current

	if th.ImageURI != nil {
		result.ImageURI = *th.ImageURI
	}

	if th.StatusPayment != nil {
		result.StatusPayment = *th.StatusPayment
	}

	return result, nil
}
