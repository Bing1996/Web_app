package router

import (
	"Web_App/middleware/logger"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()
	r.Use(logger.GinLogger())

	return r
}
