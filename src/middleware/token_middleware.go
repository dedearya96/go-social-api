package middleware

import (
	"errors"
	"net/http"

	"go-social-api/src/config"
	auth "go-social-api/src/models/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TokenAppMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		appKey := c.GetHeader("App_Key")

		if appKey == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "App_Key header is required",
			})
			c.Abort()
			return
		}

		var token auth.TokenApp
		err := config.DBAuth.Where("token = ?", appKey).First(&token).Error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusUnauthorized, gin.H{
					"status":  false,
					"message": "Your app_key is invalid!",
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  false,
					"message": "Internal server error during authentication",
				})
			}
			c.Abort()
			return
		}

		c.Set("app_token", token)

		c.Next()
	}
}
