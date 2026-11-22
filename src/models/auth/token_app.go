package auth

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TokenApp struct {
	ID        string         `json:"id" gorm:"type:char(36);primaryKey"`
	AppName   string         `json:"app_name" gorm:"type:varchar(255);not null"`
	Token     string         `json:"token" gorm:"type:varchar(255);unique;not null"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (TokenApp) TableName() string {
	return "token_apps"
}

func (t *TokenApp) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == "" {
		t.ID = uuid.New().String()
	}
	return
}
