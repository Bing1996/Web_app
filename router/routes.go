package router

import (
	controller "Web_App/controller/auth"
	"Web_App/middleware/logger"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()

	r.Use(logger.GinLogger())

	authGroup := r.Group("api/auth")
	{
		authGroup.POST("/register", controller.Register())
	}

	return r
}
