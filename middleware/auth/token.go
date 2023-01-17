package middleware

import (
	"Web_App/common"
	"Web_App/controller"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	authorizationMethod = "Bearer"
)

func HMACAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 客户端携带Token 并存储在请求头的Authoritarian里面
		// 并使用Bearer开头
		authHeader := context.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseError(context, controller.CodeEmptyAuthorizationHeader)
			context.Abort()
			return
		}

		// Bearer: Token
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == authorizationMethod) {
			controller.ResponseError(context, controller.CodeInvalidTokenFormat)
			context.Abort()
			return
		}

		// 解析Token
		claims, err := common.ParseToken(parts[1])
		if err != nil {
			controller.ResponseError(context, controller.CodeInvalidToken)
			context.Abort()
			return
		}

		// 保存至context中
		context.Set(controller.CtxCurrentUser, claims.UserName)

		// 保存为int64的类型，注意拿去的any类型断言
		context.Set(controller.CtxCurrentUserID, claims.UserID)
		// 后续通过Get拿到username名字
		context.Next()
	}
}
