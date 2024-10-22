package api

import (
	"github.com/gin-gonic/gin"

	"github.com/aSel1x/Gin_Template/core"
	"github.com/aSel1x/Gin_Template/models"
)

type UserController struct {
	*core.AppProvider
}

func (uc *UserController) Create(ctx *gin.Context) {
	var data models.UserCreate
	ctx.BindJSON(&data)
	user, _ := uc.UserUsecase.Create(data)
	ctx.JSON(200, user)
}

func (uc *UserController) Auth(ctx *gin.Context) {
	var data models.UserCreate
	ctx.BindJSON(&data)
	user, _ := uc.UserUsecase.Authenticate(data)
	auth, _ := uc.OAuth2(user)
	ctx.JSON(200, auth)
}

func (uc *UserController) Retrieve(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)
	ctx.JSON(200, user)
}
