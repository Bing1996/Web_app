package model

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	UserID   int64  `json:"user_id,omitempty" `
	UserName string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Gender   bool   `json:"gender,omitempty"`
	Email    string `json:"email,omitempty"`
}
