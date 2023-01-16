package service

import (
	"Web_App/model"
	"Web_App/repository/mysql"
)

func ShowAllCommunityList() (communityList []*model.Community, err error) {
	communityList, err = mysql.GetAllCommunities()
	return communityList, err
}

func CreateNewCommunity(request model.ParamCreateCommunity) error {
	var c model.Community

	c.CommunityID = 1
	c.CommunityName = request.CommunityName
	c.Introduction = request.Introduction

	err := mysql.InsertNewCommunityRecord(c)
	if err != nil {
		return err
	}
	return nil
}
