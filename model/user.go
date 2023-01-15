package model

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	UserID   int64  `json:"user_id,omitempty"  gorm:"column:user_id;type:bigint(20);not null"`
	UserName string `json:"username,omitempty" gorm:"column:username;type:varchar(64);not null"`
	Password string `json:"password,omitempty" gorm:"column:password;type:varchar(64);not null"`
	Gender   bool   `json:"gender,omitempty" gorm:"column:gender;type:tinyint(4)"`
	Email    string `json:"email,omitempty" gorm:"column:email;type:varchar(64)"`
}
