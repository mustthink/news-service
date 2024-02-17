package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/mustthink/news-service/internal/models"
)

type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) GetNewsByID(id uint) (*models.News, error) {
	args := m.Called(id)
	return args.Get(0).(*models.News), args.Error(1)
}

func (m *MockStorage) GetNews() ([]models.News, error) {
	args := m.Called()
	return args.Get(0).([]models.News), args.Error(1)
}

func (m *MockStorage) CreateNews(news *models.News) (uint, error) {
	args := m.Called(news)
	return args.Get(0).(uint), args.Error(1)
}

func (m *MockStorage) UpdateNews(news *models.News, id uint) (bool, error) {
	args := m.Called(news, id)
	return args.Get(0).(bool), args.Error(1)
}

func (m *MockStorage) IsNewsExist(news *models.News) (bool, error) {
	args := m.Called(news)
	return args.Get(0).(bool), args.Error(1)
}

func (m *MockStorage) DeleteNews(id uint) (bool, error) {
	args := m.Called(id)
	return args.Get(0).(bool), args.Error(1)
}

func TestNewsHandler(t *testing.T) {
	t.Run("get news by id", func(t *testing.T) {
		storage := new(MockStorage)
		storage.On("GetNewsByID", uint(1)).Return(&models.News{Model: gorm.Model{ID: 1}}, nil)

		req, _ := http.NewRequest("GET", "/news/1", nil)
		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/news/{id}", News(storage, nil))
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		storage.AssertExpectations(t)
	})

	t.Run("get news error", func(t *testing.T) {
		storage := new(MockStorage)
		storage.On("GetNews").Return([]models.News{{}}, errors.New("error"))

		req, _ := http.NewRequest("GET", "/news", nil)
		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/news", News(storage, logrus.New().WithField("want error", "true")))
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		storage.AssertExpectations(t)
	})

	t.Run("post news", func(t *testing.T) {
		storage := new(MockStorage)
		storage.On("IsNewsExist", mock.Anything).Return(false, nil)
		storage.On("CreateNews", mock.Anything).Return(uint(1), nil)

		req, _ := http.NewRequest("POST", "/news", strings.NewReader(`{
					"title": "test title",
					"content": "test content 2",
					"author_id": 1,
					"topic_id": 1
				}`))
		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/news", News(storage, nil)).Methods("POST")
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		storage.AssertExpectations(t)
	})

	t.Run("delete news", func(t *testing.T) {
		storage := new(MockStorage)
		storage.On("DeleteNews", uint(1)).Return(true, nil)

		req, _ := http.NewRequest("DELETE", "/news/1", nil)
		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/news/{id}", News(storage, nil)).Methods("DELETE")
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		storage.AssertExpectations(t)
	})
}
