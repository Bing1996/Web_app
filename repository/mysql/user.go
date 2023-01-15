package mysql

import (
	"Web_App/common"
	"Web_App/model"
	"errors"
)

// 把每一步数据库操作封装成函数
// 等待Service曾根据业务逻辑调用

var (
	ErrorQueryUserNotFound = errors.New("user not found")
)

func CheckUserExistByName(username string) (any, bool) {
	var user model.User
	db.First(&user, "username = ?", username)

	if user.UserName == "" {
		return nil, false
	}
	return user, true
}

func InsertNewUser(u model.User) error {
	// 密码md5加密
	HashedPassword, err := common.EncryptPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = HashedPassword

	(DBMulti["cloud"]).Create(&u)
	return nil
}

func FindUserByName(username string) (user model.User, err error) {
	db.First(&user, "username = ?", username)

	if user.UserName == "" {
		return user, ErrorQueryUserNotFound
	}

	return user, nil
}
