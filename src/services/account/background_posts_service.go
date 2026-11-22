package account

import (
	"go-social-api/src/config"
	"go-social-api/src/models/account"
)

type BackgroundPostsService struct{}

func NewBackgroundPostsService() *BackgroundPostsService {
	return &BackgroundPostsService{}
}

func (s *BackgroundPostsService) GetAllBgPosts() ([]account.BackgroundPosts, error) {
	var bgPosts []account.BackgroundPosts
	err := config.DBAccount.Find(&bgPosts).Error
	return bgPosts, err
}

func (s *BackgroundPostsService) GetById(id string) (*account.BackgroundPosts, error) {
	var post account.BackgroundPosts
	err := config.DBAccount.First(&post, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}
