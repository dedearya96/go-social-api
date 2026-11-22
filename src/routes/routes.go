package routes

import (
	"go-social-api/src/controllers/account"
	"go-social-api/src/controllers/auth"
	"go-social-api/src/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	userCtrl := auth.NewUserController()
	authCtrl := auth.NewAuthController()
	bgCtrl := account.NewBackgroundPostsController()

	api := r.Group("/api")
	api.Use(middleware.TokenAppMiddleware())
	{
		api.POST("/auth/login", authCtrl.LoginWithEmail)
		api.GET("/background-posts", bgCtrl.GetAll)
		api.GET("/background-posts/:id", bgCtrl.GetById)

		api.Use(middleware.JWTMiddleware())
		{
			api.GET("/users", userCtrl.GetAll)
			api.GET("/users/:username", userCtrl.GetByUsername)
		}
	}

	return r
}
