package controller

import (
	"Web_App/common"
	"Web_App/model"
	"Web_App/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func Register() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 从前端POST请求拿到请求体中的JSON数据
		var request model.ParamRegister
		if err := context.ShouldBindJSON(&request); err != nil {
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				context.JSON(http.StatusBadRequest, gin.H{
					"msg": "请求参数有误",
					"err": err.Error(),
				})
				return
			}

			// validator.ValidatorErrors类型错误进行翻译
			context.JSON(http.StatusBadRequest, gin.H{
				"msg": "参数校验错误",
				"err": common.RemoveTopStruct(errs.Translate(common.Trans)),
			})
			return
		}

		// 业务逻辑层
		if err := service.Register(request); err != nil {

		}
	}
}
