package apiserver

import (
	"io"
	"net/http"
)

func (s *APIserver) health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "HELLO")
	}
}

func (s *APIserver) hundleDifferentFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "DifferentMsg")
	}
}

func WriteResponce()
