package redis

import (
	"context"
	"github.com/go-redis/redis/v9"
)

func SavePostCreateTime(postID, createTime int64) error {
	var ctx = context.Background()
	_, err := rdb.ZAdd(ctx, addRedisKeyPrefix(KeyPostTimeZSet), redis.Z{
		Score:  float64(createTime),
		Member: postID,
	}).Result()

	return err
}
