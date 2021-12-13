package apiserver

import (
	"net/http"
)

func (s *APIserver) health(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("health handler")
	resp := CommonResponse{Result: nil, Success: true}
	s.WriteHTTPResponse(resp, w)
}

func (s *APIserver) hundleCalculateTo(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("To handler")
	req, err := parseRequest(r)
	if err != nil {
		s.WriteHTTPResponse(ErrorResponce(err.Error()), w)
		return
	}

	key := makeCacheKey(0, req.EndIndex)
	seq, err := s.cache.GetValue(key)
	if err != nil {
		s.logger.Infof("failed get value from cache on key %v,err= %v\n", key, err)
	}

	if seq != nil {
		s.logger.Infof("success get value from cache on key %v\n", key)
		s.WriteHTTPResponse(SuccessResponce(seq), w)
		return
	}

	err = s.fibonaciProvider.Calculate(0, req.EndIndex)
	if err != nil {
		s.WriteHTTPResponse(ErrorResponce(err.Error()), w)
		return
	}

	seq = s.fibonaciProvider.FibonacciSequence
	err = s.cache.SetValue(key, seq)
	if err != nil {
		s.logger.Infof("failed set value from cache on key %v,err= %v\n", key, err)
	}

	s.WriteHTTPResponse(SuccessResponce(seq), w)
}

func (s *APIserver) hundleCalculateFromTo(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("fromTo handler")
	req, err := parseRequest(r)
	if err != nil {
		s.WriteHTTPResponse(ErrorResponce(err.Error()), w)
		return
	}

	key := makeCacheKey(req.StartIndex, req.EndIndex)
	seq, err := s.cache.GetValue(key)
	if err != nil {
		s.logger.Infof("failed get value from cache on key %v,err= %v\n", key, err)
	}
	if seq != nil {
		s.logger.Infof("success get value from cache on key %v\n", key)
		s.WriteHTTPResponse(SuccessResponce(seq), w)
		return
	}

	err = s.fibonaciProvider.Calculate(req.StartIndex, req.EndIndex)
	if err != nil {
		s.WriteHTTPResponse(ErrorResponce(err.Error()), w)
		return
	}

	seq = s.fibonaciProvider.FibonacciSequence
	err = s.cache.SetValue(key, seq)
	if err != nil {
		s.logger.Infof("failed set value from cache on key %v,err= %v\n", key, err)
	}

	s.logger.Infof("success set value on cache with key %v,err= %v\n", key, err)
	s.WriteHTTPResponse(SuccessResponce(seq), w)
}
