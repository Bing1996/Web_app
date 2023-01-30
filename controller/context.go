package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	CtxCurrentUser              = "username"
	CtxCurrentUserID            = "user_id"
	ErrorCtxCurrentUserNotFound = errors.New("context无信息")
)

func GetCurrentUser(ctx *gin.Context) (string, error) {
	value, exists := ctx.Get(CtxCurrentUser)
	if !exists {
		return "", ErrorCtxCurrentUserNotFound
	}
	username := value.(string)
	return username, nil
}

func GetCurrentUserID(ctx *gin.Context) (string, error) {
	value, exists := ctx.Get(CtxCurrentUserID)
	if !exists {
		return "nil", ErrorCtxCurrentUserNotFound
	}

	// 保存在Context的userID与claims.UserID的类型相同为int64，需要根据业务需要做类型的转换
	userID := value.(int64)
	return strconv.Itoa(int(userID)), nil
}
