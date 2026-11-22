package auth

import (
	"go-social-api/src/resources/auth"
	"go-social-api/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (ctrl *UserController) GetAll(c *gin.Context) {
	users, statsMap, err := ctrl.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal"})
		return
	}

	var resource []auth.UserResource
	for _, user := range users {
		stats := statsMap[user.ID]
		resource = append(resource, auth.ToUserResource(user, stats.Followers, stats.Following))
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": resource})
}

func (ctrl *UserController) GetByUsername(c *gin.Context) {
	username := c.Param("username")
	user, followers, following, err := ctrl.userService.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   auth.ToUserResource(*user, followers, following),
	})
}
