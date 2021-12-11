package apiserver

import (
	"net/http"
)

func (s *APIserver) health() http.HandlerFunc {
	resp := CommonResponse{Result: nil, Success: true}

	return func(w http.ResponseWriter, r *http.Request) {
		s.WriteHTTPResponse(resp, w)
	}
}

func (s *APIserver) hundleCalculateTo() http.HandlerFunc {
	resp := CommonResponse{}

	endIndex := 10

	err := s.fibonaciProvider.Calculate(0, endIndex)
	if err != nil {
		resp.Success = false
		resp.ErrMsg = err.Error()
	}
	return func(w http.ResponseWriter, r *http.Request) {
		resp.Success = true
		resp.Result = s.fibonaciProvider.FibonaciSequence
		s.WriteHTTPResponse(resp, w)
	}
}

func (s *APIserver) hundleCalculateFromTo() http.HandlerFunc {
	resp := CommonResponse{}

	return func(w http.ResponseWriter, r *http.Request) {
		s.WriteHTTPResponse(resp, w)
	}
}
