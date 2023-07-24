package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// ResponseEnvelope is Sisva standard JSON object for HTTP
// response.
type ResponseEnvelope struct {
	Data   interface{} `json:"data,omitempty"`
	Errors []string    `json:"errors,omitempty"`
	Status string      `json:"status,omitempty"`
}

// ResponseDecorator is a HTTP respose decorator.
type ResponseDecorator interface {
	// Decorate use the given response write to decorate HTTP
	// response.
	Decorate(w http.ResponseWriter)
}

// Followings are the predefined ResponseDecorator
var (
	JSONContentTypeDecorator = NewContentTypeDecorator("application/json")

	// ErrSourceNotFound is returned when there is no source
	// in the HTTP request header.
	ErrSourceNotFound = errors.New("source not found")
)

// contentTypeDecorator implements responseDecorator to set
// Content-Type in the HTTP response.
type contentTypeDecorator string

// NewContentTypeDecorator returns a ResponseDecorator to
// updates Content-Type in the HTTP response.
func NewContentTypeDecorator(contentType string) ResponseDecorator {
	return contentTypeDecorator(contentType)
}

// Decorate updates Content-Type in the HTTP response.
func (d contentTypeDecorator) Decorate(w http.ResponseWriter) {
	w.Header().Set("Content-Type", string(d))
}

// WriteResponse writes HTTP response based on the given
// arguments:
//   - w: Response writer object.
//   - body: Data to write in response. It is recommended to
//     use marshalled ResponseEnvelope.
//   - statusCode: HTTP status code.
//   - decorators: List of ResponseDecorator to updates
//     response writer as wished.
func WriteResponse(w http.ResponseWriter, body []byte, statusCode int, decorators ...ResponseDecorator) (int, error) {
	// apply decorators
	for _, decorator := range decorators {
		decorator.Decorate(w)
	}

	// write response
	w.WriteHeader(statusCode)
	return w.Write(body)
}

// WriteErrorResponse writes HTTP response based on the given
// arguments:
//   - w: Response writer object.
//   - statusCode: HTTP status code.
//   - errs: List of errror messages.
func WriteErrorResponse(w http.ResponseWriter, statusCode int, errs []string) {
	// construct response object
	response := ResponseEnvelope{
		Errors: errs,
		Status: http.StatusText(statusCode),
	}

	// marshal json
	json, err := json.Marshal(response)
	if err != nil {
		// use default error body if failed to marshall response
		json = []byte(fmt.Sprintf(`{"errors":["%s"],"status":"%s"}`, err.Error(), http.StatusText(http.StatusInternalServerError)))
	}

	// create decorators
	decorators := []ResponseDecorator{
		JSONContentTypeDecorator,
	}

	// write constructed respons
	WriteResponse(w, json, statusCode, decorators...)
}
