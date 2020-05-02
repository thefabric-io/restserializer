package restserializer

import (
	"encoding/json"
	"net/http"
)

type Serializer interface {
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

func RenderError(err error, w http.ResponseWriter, statusCode int, resource string) {
	WriteContentType(w)
	w.WriteHeader(statusCode)
	restErr := ErrBody{
		Error{Message: err.Error(), StatusCode: statusCode, Status: http.StatusText(statusCode), Resource: resource},
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(&restErr)
	if err != nil {
		RenderError(err, w, http.StatusInternalServerError, resource)
		return
	}
	return
}

func WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}
