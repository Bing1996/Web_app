package mysql

import (
	"Web_App/model"
	"errors"
)

var (
	ErrorPostNotFound = errors.New("post not found")
)

func InsertNewPost(p model.Post) error {
	err := db.Model(model.Post{}).Create(&p).Error
	return err
}

// FindPostByPostID 根据帖子ID进行Mysql查询
func FindPostByPostID(id int64) (post model.Post, err error) {
	db.Find(&post, "post_id = ?", id)
	if post.PostID == 0 {
		return post, ErrorPostNotFound
	}
	return post, nil
}

// QueryPostByPage 分页查询
func QueryPostByPage(page model.Page) (model.PostPageDetail, error) {
	var total int64
	var post []model.Post
	if err := db.Limit(page.PageSize).Offset((page.Offset - 1) * page.PageSize).Find(&post).Error; err != nil {
		return model.PostPageDetail{}, err
	}

	db.Model(&post).Count(&total)

	var response = model.PostPageDetail{
		Total:       int(total),
		CurrentPage: page.Offset,
		MultiPost:   post,
	}

	return response, nil
}
