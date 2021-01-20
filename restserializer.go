package restserializer

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

type ErrorSerializer interface {
	RenderError(err error, w http.ResponseWriter, statusCode int)
	WriteContentType(w http.ResponseWriter)
}

type ErrBody struct {
	Error Error `json:"error"`
}

type Error struct {
	Resource   string `json:"resource,omitempty"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
}

func HttpErrorRender(err error, w http.ResponseWriter, resource string) {
	var statusCode int

	switch errors.Cause(err) {
	case ErrNotFound:
		statusCode = http.StatusNotFound
	case ErrBadRequest:
		statusCode = http.StatusBadRequest
	case ErrForbidden:
		statusCode = http.StatusForbidden
	case ErrConflict:
		statusCode = http.StatusConflict
	case ErrUnknown:
		statusCode = http.StatusInternalServerError
	default:
		statusCode = http.StatusInternalServerError
	}

	RenderError(err, w, statusCode, resource)
}

func RenderError(err error, w http.ResponseWriter, statusCode int, resource string) {
	WriteContentType(w)
	w.WriteHeader(statusCode)
	restErr := ErrBody{
		Error{Message: err.Error(), StatusCode: statusCode, Status: http.StatusText(statusCode), Resource: resource},
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(&restErr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(
			fmt.Sprintf(`{"error":{"resource":"%s","message":Internal Server Error","statusCode":%d,"status":"%s"}}`,
				resource,
				http.StatusInternalServerError,
				http.StatusText(http.StatusInternalServerError))))
		return
	}
	return
}

func WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}
