package account

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Follow struct {
	ID                string    `gorm:"type:char(36);primaryKey" json:"id"`
	UsersIDFollowers  string    `json:"users_id_followers"`
	UsersIDFollowings string    `json:"users_id_followings"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`

	Followers interface{} `gorm:"-" json:"-"`
	Following interface{} `gorm:"-" json:"-"`
}

func (Follow) TableName() string { return "follows" }

func (m *Follow) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	return
}
