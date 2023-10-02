package handler

import (
	"net/http"

	"github.com/ivanovpetr/tx-parser/service/parser"
)

// Handler handles http requests
type Handler struct {
	storage parser.Storage
}

// New returns new Handler
func New(storage parser.Storage) *Handler {
	return &Handler{storage: storage}
}

// InternalServerErrorHandler writes code 500 to the response body
func InternalServerErrorHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error"))
}

// NotFoundHandler writes code 404 to the response body
func NotFoundHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found"))
}
