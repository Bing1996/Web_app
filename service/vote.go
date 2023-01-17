package service

import (
	"Web_App/repository/redis"
)

func VoteForPost(postID, userID string, value float64) error {
	if err := redis.VoteForPost(postID, userID, value); err != nil {
		return err
	}
	return nil
}
