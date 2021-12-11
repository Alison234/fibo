package apiserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go/http-rest-api/fibonaci"
)

type CommonResponse struct {
	Result  []fibonaci.Fibonaci `json:"result,omitempty"`
	ErrMsg  string              `json:"error_msg,omitempty"`
	Success bool                `json:"success"`
}

func (s *APIserver) WriteHTTPResponse(resp CommonResponse, w http.ResponseWriter) {
	body, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.logger.Info("Failed to marshal json response")

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(body)))

	w.WriteHeader(200)
	_, err = w.Write(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.logger.Info("Failed to write response")

		return
	}
	s.logger.Info("Marshaling data and write body")
}