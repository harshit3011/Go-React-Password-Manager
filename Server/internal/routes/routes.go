package routes

import "github.com/gin-gonic/gin"

func BackendRoutes(router *gin.Engine){
	router.GET("/",func(ctx *gin.Context) {
		ctx.String(200,"Hi, this is my password-manager project")
	})
	
}