package controller

import (
	"Web_App/service"
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
