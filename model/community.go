package model

import "gorm.io/gorm"

type Community struct {
	gorm.Model
	CommunityID   int    `json:"community_id,omitempty" gorm:"column:community_id;type:int(10);not null"`
	CommunityName string `json:"community_name,omitempty" gorm:"column:community_name;type:varchar(128);not null"`
	Introduction  string `json:"introduction,omitempty" gorm:"column:introduction;type:varchar(256);not null"`
}
