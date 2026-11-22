package services

import (
	"go-social-api/src/config"
	models "go-social-api/src/models/auth"
)

func GetTokensCursor(lastID string, limit int) ([]models.TokenApp, bool, error) {
	var tokens []models.TokenApp

	query := config.DBAuth.Limit(limit + 1).Order("id asc")

	if lastID != "" {
		query = query.Where("id > ?", lastID)
	}

	if err := query.Find(&tokens).Error; err != nil {
		return nil, false, err
	}

	hasMore := false
	if len(tokens) > limit {
		hasMore = true
		tokens = tokens[:limit]
	}

	return tokens, hasMore, nil
}
