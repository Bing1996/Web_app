package controller

import (
	"Web_App/model"
	"Web_App/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CreatePost 创建帖子句柄
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

// GetPostDetail 查询帖子详情
func GetPostDetail() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 获取参数
		param := context.Param("post_id")
		postId, err := strconv.ParseInt(param, 10, 64)
		if err != nil {
			ResponseError(context, CodeInvalidParam)
			return
		}

		// 基于帖子ID查询详情
		postDetail, err := service.GetPostByID(postId)
		if err != nil {
			ResponseErrorWithMsg(context, CodeServerBusy, err.Error())
			return
		}

		// 成功查询返回
		ResponseSuccessful(context, postDetail)
	}
}

// ShowPostList 分页查询帖子列表
func ShowPostList() gin.HandlerFunc {
	return func(context *gin.Context) {
		var page model.Page
		if context.Bind(&page) != nil {
			ResponseError(context, CodeInvalidParam)
			return
		}

		// 分页查询
		response, err := service.GetPostByPage(page)
		if err != nil {
			ResponseError(context, CodeServerBusy)
			return
		}

		ResponseSuccessful(context, response)
	}
}
