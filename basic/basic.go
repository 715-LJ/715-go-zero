package basic

import (
	"qiyaowu-go-zero/basic/config"
	"qiyaowu-go-zero/basic/pgsql"
)

func Init() {
	config.Init()
	pgsql.Init()
	//elsearch.Init()
	//redis.Init()
}
