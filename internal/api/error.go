package api

import (
	"net/http"
	"encoding/json"
)

// Internal HttpError
type HttpError interface {
	error
	getCode() int
	sendError(w http.ResponseWriter)
}

//*********    httpBadRequest  *********//
type httpBadRequest struct {
	msg string
}

func NewHttpBadRequest(msg string) httpBadRequest {
	return httpBadRequest{msg}
}

func (badReq httpBadRequest) Error() string {
	return string(http.StatusBadRequest)
}

func (badReq httpBadRequest) getCode() int {
	return http.StatusBadRequest
}
func (badReq httpBadRequest) sendError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)
	if err := json.NewEncoder(w).Encode(badReq); err != nil {
		panic(err)
	}
}

//*********    httpBadRequest  *********//
type httpInternalServerError struct {
	msg string
}

func NewHttpInternalServerError(msg string) httpInternalServerError {
	return httpInternalServerError{msg}
}

func (intError httpInternalServerError) Error() string {
	return string(http.StatusInternalServerError)
}

func (intError httpInternalServerError) getCode() int {
	return http.StatusInternalServerError
}
func (intError httpInternalServerError) sendError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)
	if err := json.NewEncoder(w).Encode(intError); err != nil {
		panic(err)
	}
}