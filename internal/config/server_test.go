package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerAddr(t *testing.T) {
	t.Run("default server configuration", func(t *testing.T) {
		server := Server{
			Host: "localhost",
			Port: 8080,
		}
		expectedAddr := "localhost:8080"
		assert.Equal(t, expectedAddr, server.Addr())
	})

	t.Run("custom server configuration", func(t *testing.T) {
		server := Server{
			Host: "192.168.1.1",
			Port: 3000,
		}
		expectedAddr := "192.168.1.1:3000"
		assert.Equal(t, expectedAddr, server.Addr())
	})
}
