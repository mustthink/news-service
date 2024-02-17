package service

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/mustthink/news-service/internal/config"
	"github.com/mustthink/news-service/internal/storage"
)

type Service struct {
	config  *config.Config
	server  *http.Server
	storage *storage.Storage
	logger  *logrus.Logger
}

func New() *Service {
	cfg := config.MustLoad()

	log := logrus.New()
	log.SetLevel(cfg.LogLevel())
	log.Debugf("logger successfully initialized")

	st := storage.MustNew(cfg.DB, log.WithField("component", "storage"))
	log.Debugf("storage successfully initialized")

	srv := newServer(cfg.Srv, st, log)
	log.Debugf("server successfully initialized")

	return &Service{
		config:  cfg,
		storage: st,
		server:  srv,
		logger:  log,
	}
}

func (s *Service) Run() {
	err := s.server.ListenAndServe()
	s.logger.Fatalf("server stopped with error: %s", err.Error())
}

func (s *Service) Stop() {
	if err := s.server.Shutdown(context.TODO()); err != nil {
		s.logger.Errorf("server stopped with error: %s", err.Error())
	}
	s.logger.Debugf("server successfully stopped")

	if err := s.storage.Close(); err != nil {
		s.logger.Errorf("storage stopped with error: %s", err.Error())
	}
	s.logger.Debugf("storage connection successfully closed")
}
