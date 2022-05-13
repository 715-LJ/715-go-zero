package config

import (
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"path/filepath"
	"sync"
)

var (
	pgsqlConfig    defaultPgsqlConfig
	m              sync.RWMutex
	inited         bool
	esConf         defaultEs
	redisConfig    defaultRedisConfig
	sentinelConfig redisSentinel
)

// Init 初始化配置
func Init() {

	m.Lock()
	defer m.Unlock()

	if inited {
		logx.Info("[Init] 配置已经初始化过")
		return
	}
	// 加载yml配置
	// 先加载基础配置
	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join("../../../", string(filepath.Separator))))
	pt := filepath.Join(appPath, "etc")

	os.Chdir(appPath)
	conf.MustLoad(pt+"/pgsql.yaml", &pgsqlConfig)
	//conf.MustLoad(pt+"/elasticsearch.yaml",&esConf)
	//conf.MustLoad(pt+"/redis.yaml",&redisConfig)
	// 标记已经初始化
	inited = true
}

//esConfig  获取elasticsaseRbacp配置
func GetEsConfig() (ret EsConfig) {
	return esConf
}
func GetPgsqlConfig() (ret PgsqlConfig) {
	return pgsqlConfig
}

// GetRedisConfig 获取redis配置
func GetRedisConfig() (ret RedisConfig) {
	return redisConfig
}
func GetSentinelConfig() (ret RedisSentinelConfig) {
	return sentinelConfig
}
