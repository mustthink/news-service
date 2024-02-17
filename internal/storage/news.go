package storage

import (
	"fmt"

	"github.com/mustthink/news-service/internal/models"
)

// GetNewsByID returns news with the given id
func (s *Storage) GetNewsByID(id uint) (*models.News, error) {
	var news models.News
	result := s.db.First(&news, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &news, nil
}

// GetNews returns all news
func (s *Storage) GetNews() ([]models.News, error) {
	var news []models.News
	result := s.db.Find(&news)
	if result.Error != nil {
		return nil, result.Error
	}

	return news, nil
}

// CreateNews creates a new news and returns its id
func (s *Storage) CreateNews(news *models.News) (uint, error) {
	result := s.db.Create(news)
	if result.Error != nil {
		return 0, result.Error
	}

	return news.ID, nil
}

// UpdateNews updates news with the given id, returns true if news was updated, false if not found
func (s *Storage) UpdateNews(news *models.News, id uint) (bool, error) {
	result := s.db.Model(&models.News{}).Where("id = ?", id).Update(news)
	if result.Error != nil {
		return false, result.Error
	} else if result.RowsAffected == 0 {
		return false, fmt.Errorf("news with id %d not found", id)
	}

	return true, nil
}

// IsNewsExist checks if news with the same title and author_id already exists
func (s *Storage) IsNewsExist(news *models.News) (bool, error) {
	var count int64
	result := s.db.Model(&models.News{}).Where("title = ? AND author_id = ?", news.Title, news.AuthorID).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}

	return count > 0, nil
}

// DeleteNews deletes news with the given id, returns true if news was deleted, false if not found
func (s *Storage) DeleteNews(id uint) (bool, error) {
	result := s.db.Where("id = ?", id).Delete(&models.News{})
	if result.Error != nil {
		return false, result.Error
	} else if result.RowsAffected == 0 {
		return false, fmt.Errorf("news with id %d not found", id)
	}

	return true, nil
}
