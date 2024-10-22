package api

import (
	"github.com/gin-gonic/gin"

	"github.com/aSel1x/Gin_Template/usecases"
)

func WithJWTAuth(userUsecase *usecases.UserUsecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatus(401)
			return
		}
		accessToken := authHeader[len("Bearer "):]
		user, err := userUsecase.RetrieveByToken(accessToken)
		if err != nil {
			ctx.AbortWithStatus(401)
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}
