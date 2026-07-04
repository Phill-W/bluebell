package redis

import (
	"bluebell/setting"
	"fmt"

	"github.com/go-redis/redis"
)

var (
	// client redis客户端
	client *redis.Client
	// Nil redis未查到记录时返回的错误
	Nil = redis.Nil
)

// Init 初始化连接
func Init(cfg *setting.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password, // no password set
		DB:           cfg.DB,       // use default DB
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	_, err = client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

// Close 关闭redis连接
func Close() {
	_ = client.Close()
}
