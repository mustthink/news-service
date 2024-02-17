package storage

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"

	"github.com/mustthink/news-service/internal/config"
)

// Storage is a storage layer for the news service.
type Storage struct {
	db *gorm.DB
}

// New creates a new storage layer.
func New(cfg config.Database, logger *logrus.Entry) (*Storage, error) {
	const op = "storage.New"

	db, err := gorm.Open("postgres", cfg.Uri())
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	db.SetLogger(logger)

	return &Storage{db: db}, nil
}

// MustNew creates a new storage layer and panics if an error occurs.
func MustNew(cfg config.Database, logger *logrus.Entry) *Storage {
	st, err := New(cfg, logger)
	if err != nil {
		panic(err)
	}
	return st
}

// Close closes the storage layer.
func (s *Storage) Close() error {
	return s.db.Close()
}
