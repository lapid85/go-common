package clients

import "github.com/redis/go-redis/v9"

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

// RedisDefault 获取默认的 redis 连接
func RedisDefault() *redis.Client {
	return Redis("redis://localhost:6379/0")
}
