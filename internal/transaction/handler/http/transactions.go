package http

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/tedxub2023/global/helper"
	"github.com/tedxub2023/internal/transaction"
)

type transactionsHandler struct {
	transaction transaction.Service
}

func (h *transactionsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.handleReplaceTransactionByEmail(w, r)
	default:
		helper.WriteErrorResponse(w, http.StatusMethodNotAllowed, []string{errMethodNotAllowed.Error()})
	}
}

func (h *transactionsHandler) handleReplaceTransactionByEmail(w http.ResponseWriter, r *http.Request) {
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
			log.Printf("[Transaction HTTP][handleReplaceTransactionByEmail] Failed to replace transactions by email. Source: %s, Err: %s\n", source, err.Error())
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
		request := transactionHTTP{}
		err = json.Unmarshal(body, &request)
		if err != nil {
			statusCode = http.StatusBadRequest
			errChan <- errBadRequest
			return
		}

		// format HTTP request into service object
		reqTransaction, err := parseTransactionFromCreateRequest(request)
		if err != nil {
			statusCode = http.StatusBadRequest
			errChan <- err
			return
		}

		ticketID, err := h.transaction.ReplaceTransactionByEmail(ctx, reqTransaction)
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
				log.Printf("[Transaction HTTP][handleReplaceTransactionByEmail] Internal error from ReplaceTransactionByEmail. Err: %s\n", err.Error())
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

// parseTransactionFromCreateRequest returns transaction
// from the given HTTP request object.
func parseTransactionFromCreateRequest(th transactionHTTP) (transaction.Transaction, error) {
	result := transaction.Transaction{}

	if th.Nama != nil {
		result.Nama = *th.Nama
	}

	if th.JenisKelamin != nil {
		result.JenisKelamin = *th.JenisKelamin
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

	if th.JumlahTiket != nil {
		result.JumlahTiket = *th.JumlahTiket
	}

	if th.Tanggal != nil && *th.Tanggal != "" {
		tanggal, err := time.Parse(dateFormat, *th.Tanggal)
		if err != nil {
			return transaction.Transaction{}, errInvalidDateFormat
		}

		result.Tanggal = tanggal
	}

	return result, nil
}
