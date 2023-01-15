package service

import (
	"Web_App/common"
	"Web_App/model"
	"Web_App/repository/mysql"
	"go.uber.org/zap"
)

func CreatePost(username string, request model.ParamCreatePost) error {
	var post model.Post

	// 成功绑定, 生产帖子ID
	post.PostID = common.GenID()

	// 从Context拿到作者用户信息
	user, err := mysql.FindUserByName(username)
	post.AuthorID = user.UserID

	post.Content = request.Content
	post.Title = request.Title
	post.CommunityID = 1
	post.Status = true

	err = mysql.InsertNewPost(post)
	if err != nil {
		zap.L().Fatal("cannot create post", zap.Error(err))
		return err
	}
	// 成功
	return nil

}
