package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
	"qiyaowu-go-zero/basic/config"
	"sync"
)

var (
	client *redis.Client
	m      sync.RWMutex
	inited bool
)

// Init 初始化Redis
func Init() {
	m.Lock()
	defer m.Unlock()

	if inited {
		logx.Errorf("已经初始化过Redis...")
		return
	}

	redisConfig := config.GetRedisConfig()
	redisSentConfig := config.GetSentinelConfig()
	// 打开才加载
	if redisConfig != nil && redisConfig.GetEnabled() {

		logx.Errorf("初始化Redis...")

		// 加载哨兵模式
		//if redisSentConfig.GetSentinelConfig() != nil && redisConfig.GetSentinelConfig().GetEnabled() {
		if redisSentConfig.GetEnabled() {
			logx.Errorf("初始化Redis，哨兵模式...")
			initSentinel(redisConfig, redisSentConfig)
		} else { // 普通模式
			logx.Errorf("初始化Redis，普通模式...")
			initSingle(redisConfig)
		}

		logx.Errorf("初始化Redis，检测连接...")

		pong, err := client.Ping(context.TODO()).Result()
		if err != nil {
			logx.Errorf(err.Error())
		}

		logx.Errorf("初始化Redis，检测连接Ping.")
		logx.Errorf("初始化Redis，检测连接Ping..")
		logx.Errorf("初始化Redis，检测连接Ping... %s", pong)
	}
}

// GetRedis 获取redis
func GetRedis() *redis.Client {
	return client
}

func initSentinel(redisConfig config.RedisConfig, redisSentConfig config.RedisSentinelConfig) {
	client = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    redisSentConfig.GetMaster(),
		SentinelAddrs: redisSentConfig.GetNodes(),
		DB:            redisConfig.GetDBNum(),
		Password:      redisConfig.GetPassword(),
	})

}

func initSingle(redisConfig config.RedisConfig) {
	client = redis.NewClient(&redis.Options{
		Addr:     redisConfig.GetConn(),
		Password: redisConfig.GetPassword(), // no password set
		DB:       redisConfig.GetDBNum(),    // use default DB
	})
}
