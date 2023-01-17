package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v9"
	"math"
	"time"
)

const (
	ScorePerVote   = 432
	oneWeekSeconds = 7 * 24 * 3600
)

var (
	ErrorVoteTimeExpire = errors.New("post create time exceeds one week")
	ErrorVoteRepeat     = errors.New("have voted")
)

func VoteForPost(postId, userId string, value float64) error {
	// 判断投票限制
	ctx = context.Background()

	// 此时间戳需要通过Unix()转化后才能做判断
	postCreateTime := rdb.ZScore(ctx, addRedisKeyPrefix(KeyPostTimeZSet), postId).Val()
	if float64(time.Now().Unix())-postCreateTime > oneWeekSeconds {
		// 投票超时
		return ErrorVoteTimeExpire
	}

	// 开始逻辑判断, 查询用户历史投票记录
	// 更新帖子的分数记录
	historyScoreFromUser := rdb.ZScore(ctx, addRedisKeyPrefix(KeyPostVotedZSetPF+postId), userId).Val()

	pipline := rdb.TxPipeline()
	var voteUp float64
	if value == historyScoreFromUser {
		// 投票与上次一致
		return ErrorVoteRepeat
	} else if value > historyScoreFromUser {
		voteUp = 1
	} else {
		voteUp = -1
	}

	/*
		Redis更新帖子的分数
		加入用户从+1变成-1需要扣除-432-432两次
		voteUp 控制矢量方向，绝对值控制分数的总值
	*/

	pipline.ZIncrBy(ctx,
		addRedisKeyPrefix(KeyPostScoreZSet+postId),
		voteUp*math.Abs(historyScoreFromUser-value)*ScorePerVote, postId)

	/*
		更新用户投票记录
		加入投票为0意味着用户取消投票，则从Redis删除该用户的投票纪录
		反之需要增加Redis的用户投票纪录
	*/

	if value == 0 {
		pipline.ZRem(ctx, addRedisKeyPrefix(KeyPostVotedZSetPF+postId), userId)
	} else {
		pipline.ZAdd(ctx, addRedisKeyPrefix(KeyPostVotedZSetPF+postId), redis.Z{
			Score:  value,
			Member: userId,
		})
	}

	// 执行所有pipline的Redis命令
	if _, err := pipline.Exec(ctx); err != nil {
		return err
	}

	return nil
}
