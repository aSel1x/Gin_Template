package api

import (
	"github.com/gin-gonic/gin"

	"github.com/aSel1x/Gin_Template/core"
)

func SetupRouter(container *core.AppProvider) *gin.Engine {
	r := gin.Default()

	uc := UserController{AppProvider: container}

	user := r.Group("/user")
	{
		user.POST("", uc.Create)
		user.POST("/auth", uc.Auth)
		user.GET("", WithJWTAuth(container.UserUsecase), uc.Retrieve)
	}

	return r
}
