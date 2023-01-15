package auth

import (
	"Web_App/common"
	"Web_App/controller"
	"Web_App/model"
	"Web_App/service"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Register 用户注册接口
func Register() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 从前端POST请求拿到请求体中的JSON数据
		var request model.ParamRegister
		if err := context.ShouldBindJSON(&request); err != nil {
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				controller.ResponseError(context, controller.CodeInvalidParam)
				return
			}

			// validator.ValidatorErrors类型错误进行翻译
			controller.ResponseErrorWithMsg(
				context, controller.CodeInvalidParam, common.RemoveTopStruct(errs.Translate(common.Trans)))
			return
		}

		// 业务逻辑层
		if err := service.Register(request); errors.Is(err, service.ErrorUserExist) {
			controller.ResponseError(context, controller.CodeUserExist)
		} else if errors.Is(err, service.ErrorCreateNewUser) {
			controller.ResponseError(context, controller.CodeServerBusy)
		} else {
			controller.ResponseSuccessful(context, nil)
		}
	}
}

// Login 用户登录接口
func Login() gin.HandlerFunc {
	return func(context *gin.Context) {
		var request model.ParamLogin
		if err := context.ShouldBindJSON(&request); err != nil {
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				controller.ResponseError(context, controller.CodeInvalidParam)
				return
			}

			// validator.ValidatorErrors类型错误进行翻译
			controller.ResponseErrorWithMsg(
				context, controller.CodeInvalidParam, common.RemoveTopStruct(errs.Translate(common.Trans)))
			return
		}

		// 业务层
		tokenString, err := service.Login(request)
		if errors.Is(err, service.ErrorInvalidPassword) || errors.Is(err, service.ErrorUserNotExist) {
			controller.ResponseError(context, controller.CodeInvalidPassword)
		} else {
			// 登录成功，颁发Token
			controller.ResponseSuccessful(context, gin.H{
				"token": tokenString,
			})
		}
	}
}
