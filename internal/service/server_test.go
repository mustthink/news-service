package service

import (
	"testing"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/mustthink/news-service/internal/config"
	"github.com/mustthink/news-service/internal/storage"
)

func TestNewServer(t *testing.T) {
	t.Run("creates server with correct address", func(t *testing.T) {
		cfg := config.Server{
			Host: "localhost",
			Port: 8080,
		}
		s := new(storage.Storage)
		log := logrus.New()

		server := newServer(cfg, s, log)

		assert.Equal(t, "localhost:8080", server.Addr)
	})

	t.Run("creates server with correct handler", func(t *testing.T) {
		cfg := config.Server{
			Host: "localhost",
			Port: 8080,
		}
		s := new(storage.Storage)
		log := logrus.New()

		server := newServer(cfg, s, log)

		_, ok := server.Handler.(*mux.Router)
		assert.True(t, ok)
	})
}
