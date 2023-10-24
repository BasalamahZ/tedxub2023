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
	"github.com/tedxub2023/internal/transaction"
)

type transactionsHandler struct {
	transaction transaction.Service
}

func (h *transactionsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.handleGetAllTransactions(w, r)
	case http.MethodPost:
		h.handleReplaceTransactionByEmail(w, r)
	default:
		helper.WriteErrorResponse(w, http.StatusMethodNotAllowed, []string{errMethodNotAllowed.Error()})
	}
}

func (h *transactionsHandler) handleGetAllTransactions(w http.ResponseWriter, r *http.Request) {
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
			log.Printf("[Transaction HTTP][handleGetAllTransactions] Failed to get all transaction. Source: %s, Err: %s\n", source, err.Error())
			helper.WriteErrorResponse(w, statusCode, []string{err.Error()})
			return
		}
		// success
		helper.WriteResponse(w, resBody, statusCode, helper.JSONContentTypeDecorator)
	}()

	// prepare channels for main go routine
	resChan := make(chan []transaction.Transaction, 1)
	errChan := make(chan error, 1)

	go func() {
		// parsed filter
		statusPayment, tanggal, err := parseGetTransactionsFilter(r.URL.Query())
		if err != nil {
			statusCode = http.StatusBadRequest
			errChan <- err
			return
		}

		res, err := h.transaction.GetAllTransactions(ctx, statusPayment, tanggal)
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
		transactions := make([]transactionHTTP, 0)
		for _, r := range res {
			var t transactionHTTP
			t, err = formatTransaction(r)
			if err != nil {
				return
			}
			transactions = append(transactions, t)
		}

		// construct response data
		resBody, err = json.Marshal(helper.ResponseEnvelope{
			Data: transactions,
		})
	}
}

func parseGetTransactionsFilter(request url.Values) (string, time.Time, error) {
	var statusPayment string
	if query := request.Get("status_payment"); query != "" {
		statusPayment = query
	}

	var tanggal time.Time
	if queryDate := request.Get("tanggal"); queryDate != "" {
		beforeParsed, err := time.Parse(dateFormat, queryDate)
		if err != nil {
			return "", time.Time{}, err
		}
		tanggal = beforeParsed
	}

	return statusPayment, tanggal, nil
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
