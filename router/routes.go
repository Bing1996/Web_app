package router

import (
	"Web_App/asset/settings"
	"Web_App/middleware/logger"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()

	if settings.Conf.GinConfig.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	return r
}
