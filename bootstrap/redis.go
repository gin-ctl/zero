package bootstrap

import (
	"fmt"
	"zero/package/get"
	rds "zero/package/redis"
)

// SetupRedis 初始化 Redis
func SetupRedis() {
	// 建立 Redis 连接
	rds.ConnectRedis(
		fmt.Sprintf("%s:%d", get.String("redis.host"),
			get.Int("redis.port")),
		get.String("redis.username"),
		get.String("redis.password"),
		get.Int("redis.database", 0),
	)
}
