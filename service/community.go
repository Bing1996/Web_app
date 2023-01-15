package service

import (
	"Web_App/model"
	"Web_App/repository/mysql"
)

func ShowAllCommunityList() (communityList []*model.Community, err error) {
	communityList, err = mysql.GetAllCommunities()
	return communityList, err
}
