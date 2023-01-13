package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

func Init() error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.pool_size"),
	})

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		zap.L().Fatal("cannot to connect to Redis Server")
		return err
	}

	return nil
}
