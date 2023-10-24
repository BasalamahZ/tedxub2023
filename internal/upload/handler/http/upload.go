package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"

	"github.com/tedxub2023/global/helper"
)

var (
	maxFileSize int64 = 1024 * 1024 // 1 MB
)

type uploadHandler struct {
}

func (h *uploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.handleUpload(w, r)
	default:
		helper.WriteErrorResponse(w, http.StatusMethodNotAllowed, []string{errMethodNotAllowed.Error()})
	}

}

func (h *uploadHandler) handleUpload(w http.ResponseWriter, r *http.Request) {
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
		// error
		if err != nil {
			log.Printf("[Upload HTTP][handleUpload] Failed to upload file. Source: %s, Err: %s\n", source, err.Error())
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
		// parse request body as multipart/form-data
		err = r.ParseMultipartForm(maxFileSize)
		if err != nil {
			statusCode = http.StatusBadRequest
			errChan <- errBadRequest
			return
		}

		// get file from form-data
		uploaded, uploadedHeader, err := r.FormFile("file")
		if err != nil {
			statusCode = http.StatusBadRequest
			errChan <- errBadRequest
			return
		}
		defer uploaded.Close()

		// get and validates file size
		uploadedSize := uploadedHeader.Size
		if uploadedSize > maxFileSize {
			statusCode = http.StatusBadRequest
			errChan <- errFileTooLarge
			return
		}

		cld, err := cloudinary.NewFromParams(os.Getenv("CLOUDINARY_API_NAME"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_API_SECRET"))
		if err != nil {
			statusCode = http.StatusBadRequest
			errChan <- errBadRequest
			return
		}

		// create file
		res, err := cld.Upload.Upload(ctx, uploaded, uploader.UploadParams{Folder: "Bukti-Transfer-TEDX"})
		if err != nil {
			// determine error and status code, by default its internal error
			parsedErr := errInternalServerError
			statusCode = http.StatusInternalServerError
			if v, ok := mapHTTPError[err]; ok {
				parsedErr = v
				statusCode = http.StatusBadRequest
			}

			// log the actual error if its internal error
			if statusCode == http.StatusInternalServerError {
				log.Printf("[Upload HTTP][handleUpload] Internal error from CreateFile. Err: %s\n", err.Error())
			}

			errChan <- parsedErr
			return
		}

		resChan <- res.SecureURL
	}()

	// wait and handle main go routine
	select {
	case <-ctx.Done():
		statusCode = http.StatusGatewayTimeout
		err = errRequestTimeout
	case err = <-errChan:
	case data := <-resChan:
		resBody, err = json.Marshal(helper.ResponseEnvelope{
			Data: data,
		})
	}
}
