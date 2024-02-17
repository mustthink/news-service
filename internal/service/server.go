package service

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/mustthink/news-service/internal/config"
	"github.com/mustthink/news-service/internal/service/handlers"
	"github.com/mustthink/news-service/internal/storage"
)

func newServer(cfg config.Server, s *storage.Storage, log *logrus.Logger) *http.Server {
	return &http.Server{
		Addr:    cfg.Addr(),
		Handler: getRouter(s, log),
	}
}

func getRouter(s *storage.Storage, log *logrus.Logger) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc(
		"/posts",
		handlers.News(s, log.WithField("handler", "/posts")),
	).Methods(http.MethodPost, http.MethodGet)

	router.HandleFunc(
		"/posts/{id}",
		handlers.News(s, log.WithField("handler", "/posts/{id}")),
	).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)

	return router
}
