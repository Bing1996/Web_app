package redis

// redis key

const (
	// Prefix 项目Key前缀
	Prefix = "bluebell:"

	// KeyPostTimeZSet 帖子及发帖时间
	KeyPostTimeZSet = "post:time"

	// KeyPostScoreZSet 帖子及投票分数
	KeyPostScoreZSet = "post:score"

	// KeyPostVotedZSetPF 记录用户投票类型
	KeyPostVotedZSetPF = "post:voted"
)

// addRedisKeyPrefix
func addRedisKeyPrefix(key string) string {
	return Prefix + key
}
