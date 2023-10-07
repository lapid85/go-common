package clients

import (
	"consts/consts"

	"github.com/redis/go-redis/v9"
)

// RedisClients 代码 - Redis客户端
var redisClients = map[string]*redis.Client{}

// Redis 获取 redis 连接
func Redis(connStr string) *redis.Client {
	// "redis://<user>:<pass>@localhost:6379/<db>"
	opt, err := redis.ParseURL(connStr)
	if err != nil {
		panic(err)
	}

	return redis.NewClient(opt)
	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "", // 没有密码，默认值
	// 	DB:       0,  // 默认DB 0
	// })
	// return rdb, nil
}

// RedisSystem 获取默认的 redis 连接
func RedisSystem() *redis.Client {
	return Redis("redis://localhost:6379/0")
}

// GetRedisBySite 依据平台获取DB
func GetRedisBySite(siteCode string) *redis.Client {
	if siteCode == "" {
		panic("未指定平台名称")
	}
	if val, exists := redisClients[siteCode]; exists {
		return val
	}

	if val, exists := consts.SiteRedisStrings[siteCode]; !exists {
		panic("未找到平台(" + siteCode + ")的数据库连接信息")
	} else {
		db := Redis(val)
		redisClients[siteCode] = db
		return db
	}
}
