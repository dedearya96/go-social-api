package auth

import (
	"go-social-api/src/requests/auth"
	"go-social-api/src/services"
	"go-social-api/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		authService: services.NewAuthService(),
	}
}

func (ctrl *AuthController) LoginWithEmail(c *gin.Context) {
	var req auth.LoginRequest

	if !utils.ValidateRequest(c, &req) {
		return
	}

	user, accessToken, refreshToken, expiresAt, err := ctrl.authService.LoginWithEmail(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": err.Error()})
		return
	}

	utils.SendLoginResponse(c, user, accessToken, refreshToken, expiresAt)
}
