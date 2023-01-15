package hello

import (
	"Web_App/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 从context拿到username
		username, err := controller.GetCurrentUser(context)
		if err != nil {
			controller.ResponseError(context, controller.CodeCtxCurrentUserNotFound)
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"msg": fmt.Sprintf("welcome %s!", username),
		})
		return
	}
}
