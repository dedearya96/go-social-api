package account

import (
	resAccount "go-social-api/src/resources/account"
	"go-social-api/src/services/account"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BackgroundPostsController struct {
	bgService *account.BackgroundPostsService
}

func NewBackgroundPostsController() *BackgroundPostsController {
	return &BackgroundPostsController{
		bgService: account.NewBackgroundPostsService(),
	}
}

func (ctrl *BackgroundPostsController) GetAll(c *gin.Context) {
	posts, err := ctrl.bgService.GetAllBgPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed fetch data"})
		return
	}
	var resource []resAccount.BackgroundPostsResource
	for _, post := range posts {
		resource = append(resource, resAccount.ToResource(post))
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": resource})
}

func (ctrl *BackgroundPostsController) GetById(c *gin.Context) {
	id := c.Param("id")

	post, err := ctrl.bgService.GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Data Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    resAccount.ToResource(*post),
	})
}
