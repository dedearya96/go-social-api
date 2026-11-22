package auth

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID              string          `gorm:"type:char(36);primaryKey" json:"id"`
	CoverProfiles   []CoverProfile  `gorm:"foreignKey:UsersID" json:"cover_profiles,omitempty"`
	PhotosProfiles  []PhotosProfile `gorm:"foreignKey:UsersID" json:"photos_profiles,omitempty"`
	SignupWith      string          `gorm:"type:enum('email', 'google')" json:"signup_with"`
	Role            string          `gorm:"type:enum(''accommodation', 'restos', 'services', 'stores', 'ticket', 'users');default:users" json:"role"`
	Name            string          `json:"name"`
	Email           string          `gorm:"unique" json:"email"`
	EmailVerifiedAt *time.Time      `json:"email_verified_at"`
	Password        string          `json:"-"`
	GoogleID        *string         `gorm:"unique" json:"google_id"`
	Birthday        time.Time       `json:"birthday"`
	RememberToken   string          `json:"remember_token,omitempty"`
	Username        string          `gorm:"unique" json:"username"`
	Phone           *string         `json:"phone"`
	Gender          *string         `gorm:"type:enum('male', 'female')" json:"gender"`
	AccountLevel    string          `gorm:"type:enum('Free', 'VIP');default:Free" json:"account_level"`
	Bio             *string         `json:"bio"`
	Website         *string         `json:"website"`
	Hobby           *string         `json:"hobby"`
	Address         *string         `json:"address"`
	Verification    *time.Time      `json:"verification"`
	Bussiness       *time.Time      `json:"bussiness"`
	LastSeen        *time.Time      `json:"last_seen"`
	Blocked         *time.Time      `json:"blocked"`
	TokenFcm        *string         `json:"token_fcm"`
	TokenReferal    *string         `json:"token_referal"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	DeletedAt       gorm.DeletedAt  `gorm:"index" json:"deleted_at"`

	Follower  interface{} `gorm:"-" json:"-"`
	Following interface{} `gorm:"-" json:"-"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return
}
