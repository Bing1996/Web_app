package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	PostID      int64  `json:"post_id,omitempty" gorm:"column:post_id;type:bigint(20);not null"`
	Title       string `json:"title,omitempty" gorm:"column:title;type:varchar(128);not null"`
	Content     string `json:"content,omitempty" gorm:"column:content;type:varchar(8192);not null"`
	AuthorID    int64  `json:"author_id,omitempty" gorm:"column:author_id;type:bigint(20);not null"`
	CommunityID int    `json:"community_id,omitempty" gorm:"column:community_id;type:int(10);not null"`
	Status      bool   `json:"status,omitempty" gorm:"column:status;type:tinyint(4);not null"`
}
