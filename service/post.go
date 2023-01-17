package service

import (
	"Web_App/common"
	"Web_App/model"
	"Web_App/repository/mysql"
	"Web_App/repository/redis"
	"go.uber.org/zap"
	"time"
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

	// 向Redis存入创建时间
	err = redis.SavePostCreateTime(post.PostID, time.Now().Unix())

	// 成功
	return err

}

// GetPostByID 查询帖子的业务逻辑
func GetPostByID(postID int64) (model.PostDetail, error) {
	// 查询帖子详情
	post, err := mysql.FindPostByPostID(postID)
	if err != nil {
		return model.PostDetail{}, err
	}

	// 查询用户信息
	userFromAuthorID, err := mysql.FindUserByID(post.AuthorID)
	if err != nil {
		return model.PostDetail{}, err
	}

	// 查询社区信息
	community, err := mysql.FindCommunityByID(post.CommunityID)
	if err != nil {
		return model.PostDetail{}, err
	}

	// 成功查询到所有事务信息
	postDetail := model.PostDetail{
		Title:     post.Title,
		Content:   post.Content,
		Status:    post.Status,
		User:      &userFromAuthorID,
		Community: &community,
	}

	return postDetail, nil
}

func GetPostByPage(page model.Page) (response model.PostPageDetail, err error) {
	// 分页查询
	response, err = mysql.QueryPostByPage(page)
	if err != nil {
		return model.PostPageDetail{}, err
	}

	return response, nil
}
