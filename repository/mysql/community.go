package mysql

import (
	"Web_App/model"
	"errors"
	"go.uber.org/zap"
)

var (
	ErrorCommunityNotFound      = errors.New("community not found")
	ErrorCommunityCreateFailure = errors.New("cannot insert new record to communities table")
)

func GetAllCommunities() (CommunityList []*model.Community, err error) {
	db.Find(&CommunityList)
	if len(CommunityList) == 0 {
		zap.L().Warn("community table is empty")
		return nil, errors.New("no community found")
	}

	return CommunityList, nil
}

func FindCommunityByID(communityID int) (c model.Community, err error) {
	db.Find(&c, "community_id = ?", communityID)
	if c.CommunityID == 0 {
		return model.Community{}, ErrorCommunityNotFound
	}

	return c, nil
}

func InsertNewCommunityRecord(c model.Community) error {
	err := db.Model(model.Community{}).Create(&c).Error
	if err != nil {
		zap.L().Warn("cannot insert new record to communities table: ", zap.Error(err))
		return err
	}
	return nil
}
