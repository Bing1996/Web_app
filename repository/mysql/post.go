package mysql

import "Web_App/model"

func InsertNewPost(p model.Post) error {
	err := db.Model(model.Post{}).Create(&p).Error
	return err
}
