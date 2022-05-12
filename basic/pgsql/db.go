package pgsql

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"qiyaowu-go-zero/basic/config"
	"sync"
)

var (
	inited  bool
	pgsqlDB *gorm.DB
	m       sync.RWMutex
)

// Init 初始化数据库
func Init() {
	m.Lock()
	defer m.Unlock()

	var err error

	if inited {
		err = fmt.Errorf("[Init] pgsql 已经初始化过")
		log.Fatalln(err.Error())
		return
	}

	// 如果配置声明使用mysql
	if config.GetPgsqlConfig().GetEnabled() {
		initPgsql()
	}

	inited = true
}

// GetDB 获取db
func GetDB() *gorm.DB {
	return pgsqlDB
}
