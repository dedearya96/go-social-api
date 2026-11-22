package account

import (
	"go-social-api/src/models/account"
	"os"
)

type BackgroundPostsResource struct {
	ID               string `json:"id"`
	ImagesBackground string `json:"images_background"`
	ImagesThumbnail  string `json:"images_thumbnail"`
	Types            string `json:"types"`
	TextColor        string `json:"text_color"`
}

func ToResource(model account.BackgroundPosts) BackgroundPostsResource {
	storageURL := os.Getenv("STORAGE_URL")
	return BackgroundPostsResource{
		ID:               model.ID,
		ImagesBackground: storageURL + model.ImagesBackground,
		ImagesThumbnail:  storageURL + model.ImagesThumbnail,
		Types:            model.Types,
		TextColor:        model.TextColor,
	}
}
