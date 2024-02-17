package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// News represents the news model
type News struct {
	gorm.Model

	AuthorID uint64 `json:"author_id"`
	TopicID  uint64 `json:"topic_id"`

	Title   string `gorm:"size:255" json:"title"`
	Content string `gorm:"type:text" json:"content"`
}

// Validate checks if the news is valid for saving into the database
func (n News) Validate() error {
	switch {
	case len(n.Title) < 0, len(n.Title) > 255:
		return fmt.Errorf("invalid title")
	case n.Content == "":
		return fmt.Errorf("content is empty")
	case n.AuthorID == 0:
		return fmt.Errorf("author id is not set")
	case n.TopicID == 0:
		return fmt.Errorf("topic id is not set")
	case n.ID != 0:
		return fmt.Errorf("id should not be set")
	case !n.CreatedAt.IsZero() || !n.UpdatedAt.IsZero() || n.DeletedAt != nil:
		return fmt.Errorf("timestamps should not be set")
	default:
		return nil
	}
}
