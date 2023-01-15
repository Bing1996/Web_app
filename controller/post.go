package controller

import (
	"Web_App/model"
	"Web_App/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePost() gin.HandlerFunc {
	return func(context *gin.Context) {
		var request model.ParamCreatePost

		err := context.ShouldBindJSON(&request)
		if err != nil {
			fmt.Println("bind fail")
			ResponseError(context, CodeInvalidParam)
			return
		}

		// 拿到Context中的username
		username, err := GetCurrentUser(context)
		if err != nil {
			zap.L().Warn("user in context not found", zap.Error(err))
			ResponseError(context, CodeServerBusy)
			return
		}

		// 创建帖子
		err = service.CreatePost(username, request)
		if err != nil {
			zap.L().Warn("cannot create post", zap.Error(err))
			ResponseError(context, CodeServerBusy)
			return
		}
		// 创建成功
		ResponseSuccessful(context, gin.H{
			"author": username,
			"title":  request.Title,
		})

	}
}
