package service

import (
	"Web_App/common"
	"Web_App/model"
	"Web_App/repository/mysql"
	"errors"
	"fmt"
	"go.uber.org/zap"
)

var (
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorUserExist       = errors.New("用户存在")
	ErrorCreateNewUser   = errors.New("创建用户失败")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
)

// Register 用户注册业务逻辑
func Register(request model.ParamRegister) error {

	user := model.User{
		// 通过雪花算法自动生成用户ID
		UserID:   common.GenID(),
		UserName: request.UserName,
		Password: request.Password,
	}

	// 查询是否用户存在
	if _, ok := mysql.CheckUserExistByName(user.UserName); ok {
		return ErrorUserExist
	}

	// 插入用户至数据库
	if err := mysql.InsertNewUser(user); err != nil {
		return ErrorCreateNewUser
	}
	zap.L().Info(fmt.Sprintf("user %s insert to users table", user.UserName))
	return nil
}

func Login(request model.ParamLogin) (string, error) {
	// 查询用户是否存在
	userQuery, ok := mysql.CheckUserExistByName(request.UserName)
	if !ok {
		return "", ErrorUserNotExist
	}

	// 密码匹配
	user := userQuery.(model.User)
	hashedPassword, _ := common.EncryptPassword(request.Password)
	if user.Password != hashedPassword {
		return "", ErrorInvalidPassword
	}

	tokenString, _ := common.GenToken(user.UserID, user.UserName)
	return tokenString, nil
}
