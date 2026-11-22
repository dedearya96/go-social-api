package services

import (
	"errors"
	"go-social-api/src/config"
	"go-social-api/src/models/account"
	"go-social-api/src/models/auth"

	"gorm.io/gorm"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

type UserStats struct {
	Followers int64
	Following int64
}

func (s *UserService) GetFollowCounts(userID string) (int64, int64, error) {
	var followersCount, followingCount int64

	err := config.DBAccount.Model(&account.Follow{}).Where("users_id_followings = ?", userID).Count(&followersCount).Error
	if err != nil {
		return 0, 0, err
	}

	err = config.DBAccount.Model(&account.Follow{}).Where("users_id_followers = ?", userID).Count(&followingCount).Error
	if err != nil {
		return 0, 0, err
	}

	return followersCount, followingCount, nil
}

func (s *UserService) GetAllUsers() ([]auth.User, map[string]UserStats, error) {
	var users []auth.User

	err := config.DBAuth.Preload("PhotosProfiles").Preload("CoverProfiles").Omit("password").Find(&users).Error
	if err != nil {
		return nil, nil, err
	}

	statsMap := make(map[string]UserStats)
	for _, user := range users {
		followers, following, _ := s.GetFollowCounts(user.ID)
		statsMap[user.ID] = UserStats{
			Followers: followers,
			Following: following,
		}
	}

	return users, statsMap, nil
}
func (s *UserService) GetUserByUsername(username string) (*auth.User, int64, int64, error) {
	var user auth.User

	err := config.DBAuth.Preload("PhotosProfiles").Preload("CoverProfiles").Where("username = ?", username).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, 0, errors.New("User Not Found")
		}
		return nil, 0, 0, err
	}

	followers, following, err := s.GetFollowCounts(user.ID)

	return &user, followers, following, nil
}
