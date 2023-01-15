package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var (
	CtxCurrentUser              = "username"
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
