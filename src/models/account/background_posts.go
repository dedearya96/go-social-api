package account

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BackgroundPosts struct {
	ID               string         `gorm:"type:char(36);primaryKey" json:"id"`
	ImagesBackground string         `json:"images_background"`
	ImagesThumbnail  string         `json:"images_thumbnail"`
	Types            string         `json:"types"`
	TextColor        string         `json:"text_color"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}

func (BackgroundPosts) TableName() string {
	return "background_posts"
}

func (m *BackgroundPosts) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	return
}
