package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 响应体
/*
{
	"code": "业务错误码",
	"msg": "提示信息",
	"data": {} // 数据
}
*/
type Response struct {
	Code ResponseCode `json:"code,omitempty"`
	Msg  interface{}  `json:"msg,omitempty"`
	Data interface{}  `json:"data,omitempty"`
}

func ResponseError(ctx *gin.Context, code ResponseCode) {
	rd := &Response{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}

	ctx.JSON(http.StatusOK, rd)
}

func ResponseErrorWithMsg(ctx *gin.Context, code ResponseCode, msg interface{}) {
	rd := &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}

	ctx.JSON(http.StatusOK, rd)
}

func ResponseSuccessful(ctx *gin.Context, data interface{}) {
	rd := &Response{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	}

	ctx.JSON(http.StatusOK, rd)
}
