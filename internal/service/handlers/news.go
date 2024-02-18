package handlers

import (
	"encoding/json"
	errorsBasic "errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/mustthink/news-service/internal/errors"
	"github.com/mustthink/news-service/internal/models"
)

// NewsProvider provides methods for retrieving news.
type NewsProvider interface {
	// GetNewsByID retrieves a news item by its ID.
	// Parameters:
	//   id: The ID of the news item.
	// Returns:
	//   The news item, or an error if the news item could not be retrieved.
	GetNewsByID(id uint) (*models.News, error)

	// GetNews retrieves all news items.
	// Returns:
	//   A list of news items, or an error if the news items could not be retrieved.
	GetNews() ([]models.News, error)
}

func getNews(storage NewsProvider, id uint) (interface{}, error) {
	var (
		news interface{}
		err  error
	)

	if id != 0 {
		news, err = storage.GetNewsByID(id)
	} else {
		news, err = storage.GetNews()
	}

	if err != nil {
		return nil, errors.NewErrWithStatusf(http.StatusInternalServerError, "couldn't get news from storage w err: %w", err)
	}
	return news, nil
}

// NewsSaver provides methods for saving news.
type NewsSaver interface {
	// CreateNews creates a new news item.
	// Parameters:
	//   news: The news item to create.
	// Returns:
	//   The ID of the created news item, or an error if the news item could not be created.
	CreateNews(news *models.News) (uint, error)

	// UpdateNews updates a news item.
	// Parameters:
	//   news: The news item to update.
	//   id: The ID of the news item to update.
	// Returns:
	//   A boolean indicating whether the update was successful, or an error if the news item could not be updated.
	UpdateNews(news *models.News, id uint) (bool, error)

	// IsNewsExist checks if a news item exists.
	// Parameters:
	//   news: The news item to check.
	// Returns:
	//   A boolean indicating whether the news item exists, or an error if the check could not be performed.
	IsNewsExist(news *models.News) (bool, error)
}

func postNews(storage NewsSaver, news *models.News) (interface{}, error) {
	if err := news.Validate(); err != nil {
		return nil, errors.NewErrWithStatusf(http.StatusBadRequest, "news validation failed: %w", err)
	}

	exist, err := storage.IsNewsExist(news)
	if err != nil {
		return nil, errors.NewErrWithStatusf(http.StatusInternalServerError, "couldn't check if news exist w err: %w", err)
	}
	if exist {
		return nil, errors.NewErrWithStatusf(http.StatusConflict, "news already exist")
	}

	id, err := storage.CreateNews(news)
	if err != nil {
		return nil, errors.NewErrWithStatusf(http.StatusInternalServerError, "couldn't create news w err: %w", err)
	}

	return struct {
		ID uint `json:"id"`
	}{
		ID: id,
	}, nil
}

func putNews(storage NewsSaver, news *models.News, id uint) (interface{}, error) {
	success, err := storage.UpdateNews(news, id)
	if err != nil {
		return nil, errors.NewErrWithStatusf(http.StatusInternalServerError, "couldn't update news w err: %w", err)
	}
	if !success {
		return nil, errors.NewErrWithStatusf(http.StatusNotFound, "news w id: %d not found", id)
	}

	return struct {
		Success bool `json:"success"`
	}{
		Success: success,
	}, nil
}

// NewsRemover provides methods for removing news.
type NewsRemover interface {
	// DeleteNews deletes a news item.
	// Parameters:
	//   id: The ID of the news item to delete.
	// Returns:
	//   A boolean indicating whether the deletion was successful, or an error if the news item could not be deleted.
	DeleteNews(id uint) (bool, error)
}

func deleteNews(storage NewsRemover, id uint) (interface{}, error) {
	success, err := storage.DeleteNews(id)
	if err != nil {
		return nil, errors.NewErrWithStatusf(http.StatusInternalServerError, "couldn't delete news w err: %w", err)
	}
	if !success {
		return nil, errors.NewErrWithStatusf(http.StatusNotFound, "news with id: %d not found", id)
	}

	return struct {
		Success bool `json:"success"`
	}{
		Success: success,
	}, nil
}

// Storage provides methods for retrieving, saving, and removing news.
type Storage interface {
	NewsProvider
	NewsSaver
	NewsRemover
}

func News(storage Storage, log *logrus.Entry) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		rawID := mux.Vars(request)["id"]
		id, _ := strconv.Atoi(rawID)
		if id < 0 {
			log.Errorf("handling request with invalid id: %d", id)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		var (
			response interface{}
			err      error
		)
		switch request.Method {
		case http.MethodGet:
			log.Debugf("handling GET request for news with id: %d", id)
			response, err = getNews(storage, uint(id))

		case http.MethodDelete:
			log.Debugf("handling DELETE request for news with id: %d", id)
			response, err = deleteNews(storage, uint(id))

		case http.MethodPost, http.MethodPut:
			body, err := io.ReadAll(request.Body)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}

			var news models.News
			if err := json.Unmarshal(body, &news); err != nil {
				writer.WriteHeader(http.StatusBadRequest)
				return
			}

			if request.Method == http.MethodPost {
				log.Debugf("handling POST request for news: %+v", news)
				response, err = postNews(storage, &news)
			} else {
				log.Debugf("handling PUT request for news with id: %d, news: %+v", id, news)
				response, err = putNews(storage, &news, uint(id))
			}
		}

		if err != nil {
			log.Errorf("error handling request: %s", err.Error())

			var errStatus errors.ErrWithStatus
			if errorsBasic.As(err, &errStatus) {
				writer.WriteHeader(errStatus.Status())
			} else {
				writer.WriteHeader(http.StatusInternalServerError)
			}

			return
		}

		data, err := json.Marshal(response)
		if err != nil {
			log.Errorf("error marshaling response: %s", err.Error())
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.Write(data)
	}
}
