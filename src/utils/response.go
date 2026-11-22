package utils

import (
	"go-social-api/src/models/auth"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SendLoginResponse(c *gin.Context, user *auth.User, accessToken, refreshToken string, expiresAt time.Time) {
	maxAge := int(time.Until(expiresAt).Seconds())
	c.SetCookie("refresh_token", refreshToken, maxAge, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"success":       true,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"user": gin.H{
			"id":             user.ID,
			"signup_with":    user.SignupWith,
			"role":           user.Role,
			"name":           user.Name,
			"username":       user.Username,
			"email_verified": user.EmailVerifiedAt != nil,
		},
	})
}
