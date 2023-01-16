package controller

import (
	"Web_App/model"
	"Web_App/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

// CommunityList 查询到所有的社区model.Community{} 以列表形式返回
func CommunityList() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 查询到所有的社区model.Community{} 以列表形式返回
		communityList, err := service.ShowAllCommunityList()
		if err != nil {
			ResponseError(context, CodeServerBusy)
			return
		}
		ResponseSuccessful(context, communityList)
	}
}

// CreateCommunity 创建社区
func CreateCommunity() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 数据绑定
		var request model.ParamCreateCommunity
		if err := context.ShouldBindJSON(&request); err != nil {
			ResponseError(context, CodeInvalidParam)
			return
		}

		// 创建
		if err := service.CreateNewCommunity(request); err != nil {
			ResponseError(context, CodeInvalidParam)
			return
		}

		ResponseSuccessful(context, fmt.Sprintf("社区 %s 创建成功", request.CommunityName))
		return
	}
}
