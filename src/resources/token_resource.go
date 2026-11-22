package resources

import (
	"go-social-api/src/models/auth"
)

type TokenResource struct {
	ID        string `json:"id"`
	TokenName string `json:"token_name"`
	Token     string `json:"token"`
	CreatedAt string `json:"created_at"`
}

func ToTokenResource(token auth.TokenApp) TokenResource {
	return TokenResource{
		ID:        token.ID,
		TokenName: token.AppName,
		Token:     token.Token,
		CreatedAt: token.CreatedAt,
	}
}

func ToTokenCollection(tokens []auth.TokenApp) []TokenResource {
	var resources []TokenResource
	for _, token := range tokens {
		resources = append(resources, ToTokenResource(token))
	}
	return resources
}
