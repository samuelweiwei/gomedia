package v1

import "github.com/gin-gonic/gin"

func RegisterUserRoutesV1(r *gin.RouterGroup) {
	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "User list",
		})
	})
	r.POST("/users", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "User registered",
		})
	})
}
