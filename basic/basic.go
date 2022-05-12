package basic

import (
	"qiyaowu-go-zero/basic/config"
	"qiyaowu-go-zero/basic/elsearch"
	"qiyaowu-go-zero/basic/pgsql"
	"qiyaowu-go-zero/basic/redis"
)

func Init() {
	config.Init()
	pgsql.Init()
	elsearch.Init()
	redis.Init()
}
