package auth

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CoverProfile struct {
	ID                   string         `gorm:"type:char(36);primaryKey" json:"id"`
	UsersID              string         `gorm:"type:char(36);index" json:"users_id"`
	FilesOri             string         `json:"files_ori"`
	FilesLargeThumbnail  string         `json:"files_large_thumbnail"`
	FilesMediumThumbnail string         `json:"files_medium_thumbnail"`
	FilesSmallThumbnail  string         `json:"files_small_thumbnail"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
	DeletedAt            gorm.DeletedAt `gorm:"index" json:"-"`

	User User `gorm:"foreignKey:UsersID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
}

func (CoverProfile) TableName() string {
	return "cover_profiles"
}

func (m *CoverProfile) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	return
}
