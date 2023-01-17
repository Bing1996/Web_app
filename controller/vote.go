package controller

import (
	"Web_App/common"
	"Web_App/model"
	"Web_App/repository/redis"
	"Web_App/service"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func VoteForPost() gin.HandlerFunc {
	return func(context *gin.Context) {
		request := new(model.ParamNewVote)
		if err := context.ShouldBindJSON(request); err != nil {
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				ResponseError(context, CodeInvalidParam)
				return
			}
			ResponseErrorWithMsg(context, CodeInvalidParam, common.RemoveTopStruct(errs.Translate(common.Trans)))
			return
		}

		// 通过context中间件保存的用户信息拿到userID
		userID, err := GetCurrentUserID(context)
		if err != nil {
			// user不存在重新登录
			ResponseError(context, CodeUserNeedToReLogin)
			return
		}

		if err := service.VoteForPost(request.PostID, userID, float64(request.Direction)); err != nil {
			// 投票失败
			if errors.Is(err, redis.ErrorVoteRepeat) {
				ResponseError(context, CodeUserRepeatVote)
				return
			}
			zap.L().Error("user vote failure: ", zap.Error(err))
			ResponseError(context, CodeServerBusy)
			return
		}

		ResponseSuccessful(context, nil)
	}
}
