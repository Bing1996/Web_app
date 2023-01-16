package router

import (
	"Web_App/asset/settings"
	"Web_App/controller"
	"Web_App/controller/auth"
	"Web_App/controller/hello"
	middleware "Web_App/middleware/auth"
	"Web_App/middleware/logger"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()

	if settings.Conf.GinConfig.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 登录服务
	authGroup := r.Group("/api/auth")
	{
		authGroup.POST("/register", auth.Register())
		authGroup.POST("/login", auth.Login())
	}

	// 测试服务
	testGroup := r.Group("/api/test", middleware.HMACAuthMiddleware())
	{
		testGroup.GET("/hello", hello.Hello())
	}

	// 社区服务
	communityGroup := r.Group("/api/community", middleware.HMACAuthMiddleware())
	{
		communityGroup.GET("/getAll", controller.CommunityList())
		communityGroup.POST("/create", controller.CreateCommunity())
	}

	// 帖子服务
	postGroup := r.Group("/api/post", middleware.HMACAuthMiddleware())
	{
		postGroup.POST("create", controller.CreatePost())
		postGroup.GET("/:post_id", controller.GetPostDetail())
		// TODO 帖子的分页列表查询功能
	}

	return r
}
