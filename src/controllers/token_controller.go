package controllers

import (
	"go-social-api/src/resources"
	"go-social-api/src/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TokenController struct {
}

func NewTokenController() *TokenController {
	return &TokenController{}
}

func (ctrl *TokenController) GetToken(c *gin.Context) {
	cursor := c.DefaultQuery("cursor", "")
	limitStr := c.DefaultQuery("limit", "10")
	limit, _ := strconv.Atoi(limitStr)

	tokens, hasMore, err := services.GetTokensCursor(cursor, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}

	tokenResources := resources.ToTokenCollection(tokens)

	var nextCursor string = ""
	if len(tokens) > 0 {
		nextCursor = tokens[len(tokens)-1].ID
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": resources.CursorResource{
			NextCursor: nextCursor,
			HasMore:    hasMore,
			Count:      len(tokenResources),
			Data:       tokenResources,
		},
	})
}
