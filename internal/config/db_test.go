package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseUri(t *testing.T) {
	t.Run("valid database configuration", func(t *testing.T) {
		db := Database{
			Host:     "localhost",
			Port:     5432,
			User:     "user",
			Password: "password",
			DBName:   "database",
		}
		expectedUri := "host=localhost port=5432 user=user dbname=database password=password sslmode=disable"
		assert.Equal(t, expectedUri, db.Uri())
	})

	t.Run("empty database configuration", func(t *testing.T) {
		db := Database{}
		expectedUri := "host= port=0 user= dbname= password= sslmode=disable"
		assert.Equal(t, expectedUri, db.Uri())
	})
}
