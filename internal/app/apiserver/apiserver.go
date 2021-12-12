package apiserver

import (
	"net/http"

	"github.com/go/http-rest-api/cache"
	"github.com/go/http-rest-api/fibonaci"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIserver struct {
	config           *Config
	logger           *logrus.Logger
	router           *mux.Router
	fibonaciProvider fibonaci.FibonaciProvider
	cache            cache.Cache
}

func New(config *Config) *APIserver {
	cacheStore := cache.NewMemCacher("localhost:11211")
	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
		cache:  cacheStore,
	}
}

func (s *APIserver) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureHandlers()

	s.logger.Info("starting server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIserver) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *APIserver) configureHandlers() {
	s.router.HandleFunc("/health", s.health)
	s.router.HandleFunc("/to", s.hundleCalculateTo)
	s.router.HandleFunc("/fromTo", s.hundleCalculateFromTo)
}
