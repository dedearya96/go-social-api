package auth

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshToken struct {
	ID        string    `gorm:"primaryKey"`
	UsersID   string    `gorm:"type:char(36);index"`
	Token     string    `gorm:"type:varchar(255);unique"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (RefreshToken) TableName() string {
	return "refresh_tokens"
}

func (m *RefreshToken) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	return
}
