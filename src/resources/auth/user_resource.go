package auth

import (
	"go-social-api/src/models/auth"
	"go-social-api/src/services"
	"os"
)

type MediaResource struct {
	FilesOri    string `json:"files_ori"`
	FilesLarge  string `json:"files_large"`
	FilesMedium string `json:"files_medium"`
	FilesSmall  string `json:"files_small"`
}

type UserResource struct {
	ID             string         `json:"id"`
	Photo          *MediaResource `json:"photo"`
	Cover          *MediaResource `json:"cover"`
	Name           string         `json:"name"`
	Username       string         `json:"username"`
	Role           string         `json:"role"`
	AccountLevel   string         `json:"account_level"`
	Bio            *string        `json:"bio"`
	Website        *string        `json:"website"`
	Phone          *string        `json:"phone"`
	LastSeen       string         `json:"last_seen"`
	CreatedAt      string         `json:"created_at"`
	FollowersCount int64          `json:"followers_count"`
	FollowingCount int64          `json:"following_count"`
}

func ToUserResource(user auth.User, followersCount int64, followingCount int64) UserResource {
	storageURL := os.Getenv("STORAGE_URL")

	transformMedia := func(filesOri, large, medium, small string) *MediaResource {
		if filesOri == "" {
			return nil
		}
		return &MediaResource{
			FilesOri:    storageURL + filesOri,
			FilesLarge:  storageURL + large,
			FilesMedium: storageURL + medium,
			FilesSmall:  storageURL + small,
		}
	}

	var photoRes *MediaResource
	if len(user.PhotosProfiles) > 0 {
		p := user.PhotosProfiles[0]
		photoRes = transformMedia(p.FilesOri, p.FilesLargeThumbnail, p.FilesMediumThumbnail, p.FilesSmallThumbnail)
	}

	var coverRes *MediaResource
	if len(user.CoverProfiles) > 0 {
		c := user.CoverProfiles[0]
		coverRes = transformMedia(c.FilesOri, c.FilesLargeThumbnail, c.FilesMediumThumbnail, c.FilesSmallThumbnail)
	}

	var lastSeenStr string
	if user.LastSeen != nil {
		lastSeenStr = user.LastSeen.Format("2006-01-02 15:04:05")
	} else {
		lastSeenStr = ""
	}

	var createdAtStr string
	if !user.CreatedAt.IsZero() {
		createdAtStr = user.CreatedAt.Format("2006-01-02 15:04:05")
	}
	return UserResource{
		ID:             user.ID,
		Cover:          coverRes,
		Photo:          photoRes,
		Name:           user.Name,
		Username:       user.Username,
		Role:           user.Role,
		Bio:            user.Bio,
		Website:        user.Website,
		Phone:          user.Phone,
		LastSeen:       lastSeenStr,
		CreatedAt:      createdAtStr,
		FollowersCount: followersCount,
		FollowingCount: followingCount,
	}
}

func ToUserCollection(users []auth.User, statsMap map[string]services.UserStats) []UserResource {
	var resources []UserResource
	for _, user := range users {
		stats := statsMap[user.ID]
		resources = append(resources, ToUserResource(user, stats.Followers, stats.Following))
	}
	return resources
}
