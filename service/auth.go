package service

import (
	"Web_App/common"
	"Web_App/model"
	"fmt"
)

// Register 用户注册业务逻辑
func Register(request model.ParamRegister) error {

	user := model.User{
		UserID:   common.GenID(),
		UserName: request.UserName,
		Password: request.Password,
	}

	fmt.Println(user.UserName)

	return nil
}
