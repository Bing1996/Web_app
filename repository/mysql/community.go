package mysql

import (
	"Web_App/model"
	"errors"
	"go.uber.org/zap"
)

func GetAllCommunities() (CommunityList []*model.Community, err error) {
	db.Find(&CommunityList)
	if len(CommunityList) == 0 {
		zap.L().Warn("community table is empty")
		return nil, errors.New("no community found")
	}

	return CommunityList, nil
}
