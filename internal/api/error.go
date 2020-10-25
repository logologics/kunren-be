package api

import (
	"encoding/json"
	"net/http"
)

// HTTPError encodes a http error
type HTTPError struct {
	error
	Message string `json:"msg"`
	Code    int    `json:"code"`
	Context string 
}

// NewHTTPBadRequest creates a 400
func NewHTTPBadRequest(err error, msg string, ctx string) HTTPError {
	return HTTPError{err, msg, http.StatusBadRequest, ctx}
}

// NewNotFound creates a 404
func NewNotFound(err error, msg string, ctx string) HTTPError {
	return HTTPError{err, msg, http.StatusNotFound, ctx}
}

// NewHTTPInternalServerError creates a 500
func NewHTTPInternalServerError(err error, msg string, ctx string) HTTPError {
	return HTTPError{err, msg, http.StatusInternalServerError, ctx}
}

// SendError sends the error ti the writer
func (httpErr HTTPError) SendError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(httpErr.Code)
	if err := json.NewEncoder(w).Encode(httpErr); err != nil {
		panic(err)
	}
}
