// Package redis 工具包
package redis

import (
	"context"
	"github.com/gin-ctl/zero/package/logger"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

// RdsClient Redis服务
type RdsClient struct {
	Client  *redis.Client
	Context context.Context
}

// rds 全局 Redis，使用 db 1
var (
	Rds  *RdsClient
	Once sync.Once
)

// ConnectRedis 连接 redis 数据库，设置全局的 Redis 对象
func ConnectRedis(address string, username string, password string, db int) {
	Once.Do(func() {
		Rds = NewClient(address, username, password, db)
	})
}

// NewClient 创建一个新的 redis 连接
func NewClient(address string, username string, password string, db int) *RdsClient {
	// 初始化自定义的 RedisClient
	r := &RdsClient{}
	r.Context = context.Background()           //使用默认的 context
	r.Client = redis.NewClient(&redis.Options{ // 使用 redis 库里的 NewClient 初始化连接
		Addr:     address,
		Username: username,
		Password: password,
		DB:       db,
	})

	// testing
	err := r.Ping()
	if err != nil {
		panic(err.Error())
	}

	return r
}

// Ping 用以测试 redis 连接是否正常
func (rds *RdsClient) Ping() error {
	_, err := rds.Client.Ping(rds.Context).Result()
	return err
}

// Set 存储 key 对应的 value，并设置 expiration 过期时间（单位：秒）
func (rds *RdsClient) Set(key string, value interface{}, expiration time.Duration) bool {
	if err := rds.Client.Set(rds.Context, key, value, expiration).Err(); err != nil {
		logger.ErrorString("Redis", "Set", err.Error())
		return false
	}
	return true
}

// Get 获取 key 对应的 value
func (rds *RdsClient) Get(key string) string {
	result, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Get", err.Error())
		}
		return ""
	}
	return result
}

// Has 判断一个 key 是否存在，内部错误和 redis.Nil 都返回 false
func (rds *RdsClient) Has(key string) bool {
	if _, err := rds.Client.Get(rds.Context, key).Result(); err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Has", err.Error())
		}
		return false
	}
	return true
}

// Del 删除存储在 redis 里的数据，支持多个 key 传参
func (rds *RdsClient) Del(keys ...string) bool {
	if err := rds.Client.Del(rds.Context, keys...).Err(); err != nil {
		logger.ErrorString("Redis", "Del", err.Error())
		return false
	}
	return true
}

// FlushDB 清空当前 redis db 里的所有数据
func (rds *RdsClient) FlushDB() bool {
	if err := rds.Client.FlushDB(rds.Context).Err(); err != nil {
		logger.ErrorString("Redis", "FlushDB", err.Error())
		return false
	}
	return true
}

// Increment 当前参数只有 1 时，为 key 的值自增 1
// 当参数有 2 个时，第一个参数为key，第二个参数为要增加的值（int64 类型）
func (rds *RdsClient) Increment(params ...interface{}) bool {
	switch len(params) {
	case 1:
		key := params[0].(string)
		if err := rds.Client.Incr(rds.Context, key).Err(); err != nil {
			logger.ErrorString("Redis", "Increment", err.Error())
			return false
		}
	case 2:
		key := params[0].(string)
		value := params[1].(int64)
		if err := rds.Client.IncrBy(rds.Context, key, value).Err(); err != nil {
			logger.ErrorString("Redis", "Increment", err.Error())
			return false
		}
	default:
		logger.ErrorString("Redis", "Increment", "参数过多")
		return false
	}
	return true
}

// Decrement 当前参数只有 1 时，为 key 的值自减 1
// 当参数有 2 个时，第一个参数为key，第二个参数为要减去的值（int64 类型）
func (rds *RdsClient) Decrement(params ...interface{}) bool {
	switch len(params) {
	case 1:
		key := params[0].(string)
		if err := rds.Client.Decr(rds.Context, key).Err(); err != nil {
			logger.ErrorString("Redis", "Decrement", err.Error())
			return false
		}
	case 2:
		key := params[0].(string)
		value := params[1].(int64)
		if err := rds.Client.DecrBy(rds.Context, key, value).Err(); err != nil {
			logger.ErrorString("Redis", "Decrement", err.Error())
			return false
		}
	default:
		logger.ErrorString("Redis", "Decrement", "参数过多")
		return false
	}
	return true
}

// IsExists 判断一个 key 是否存在
func (rds *RdsClient) IsExists(key string) bool {
	result, err := rds.Client.Exists(rds.Context, key).Result()
	if err != nil {
		logger.ErrorString("Redis", "IsExists", err.Error())
		return false
	}
	return result == 1
}

// LLen 返回列表 key 的长度
func (rds *RdsClient) LLen(key string) int64 {
	result, err := rds.Client.LLen(rds.Context, key).Result()
	if err != nil {
		logger.ErrorString("Redis", "LLen", err.Error())
		return 0
	}
	return result
}

// LPush 将一个或多个值 value 插入到列表 key 的表头
func (rds *RdsClient) LPush(key string, values ...interface{}) bool {
	if err := rds.Client.LPush(rds.Context, key, values...).Err(); err != nil {
		logger.ErrorString("Redis", "LPush", err.Error())
		return false
	}
	return true
}

// RPop 移除并返回列表 key 的尾元素
func (rds *RdsClient) RPop(key string) string {
	result, err := rds.Client.RPop(rds.Context, key).Result()
	if err != nil {
		logger.ErrorString("Redis", "RPop", err.Error())
		return ""
	}
	return result
}

// RPeek 返回队列 key 最后一个尾元素
func (rds *RdsClient) RPeek(key string) string {
	result, err := rds.Client.LRange(rds.Context, key, -1, -1).Result()
	if err != nil {
		logger.ErrorString("Redis", "RPeek", err.Error())
		return ""
	}
	return result[0]
}

// ZAdd 存储一个有序集合
func (rds *RdsClient) ZAdd(key string, values ...*redis.Z) bool {
	if err := rds.Client.ZAdd(rds.Context, key, values...).Err(); err != nil {
		logger.ErrorString("Redis", "ZAdd", err.Error())
		return false
	}
	return true
}

// ZRange 获取一个有序集合的所有成员
func (rds *RdsClient) ZRange(key string) []string {
	result, err := rds.Client.ZRange(rds.Context, key, 0, -1).Result()
	if err != nil {
		logger.ErrorString("Redis", "ZRange", err.Error())
		return []string{}
	}
	return result
}

// SAdd 存储无序集合
func (rds *RdsClient) SAdd(key string, values ...interface{}) bool {
	if err := rds.Client.SAdd(rds.Context, key, values...).Err(); err != nil {
		logger.ErrorString("Redis", "SAdd", err.Error())
		return false
	}
	return true
}

// SMembers 获取一个无序集合的所有成员
func (rds *RdsClient) SMembers(key string) []string {
	result, err := rds.Client.SMembers(rds.Context, key).Result()
	if err != nil {
		logger.ErrorString("Redis", "SMembers", err.Error())
		return []string{}
	}
	return result
}

// HGet 获取hash元素
func (rds *RdsClient) HGet(key string, field string) string {
	result, err := rds.Client.HGet(rds.Context, key, field).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "HGet", err.Error())
		}
		return ""
	}
	return result
}

// SRandMember 随机获取集合的元素
func (rds *RdsClient) SRandMember(key string) string {
	result, err := rds.Client.SRandMember(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "SRandMember", err.Error())
		}
		return ""
	}
	return result
}
