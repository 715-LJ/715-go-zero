package pgsql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"qiyaowu-go-zero/basic/config"
	"time"
)

func initPgsql() {

	var err error
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.GetPgsqlConfig().GetURL(),
		config.GetPgsqlConfig().GetPort(),
		config.GetPgsqlConfig().GetUser(),
		config.GetPgsqlConfig().GetPassword(),
		config.GetPgsqlConfig().GetDbname(),
	)
	// 创建连接
	pgsqlDB, err = gorm.Open(
		postgres.New(
			postgres.Config{
				DSN:                  dsn, // DSN data source name
				PreferSimpleProtocol: true,
			}),
		&gorm.Config{
			PrepareStmt: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
			},
		})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	dbConn, _ := pgsqlDB.DB()

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	dbConn.SetMaxOpenConns(config.GetPgsqlConfig().GetMaxOpenConnection())

	//SetMaxIdleConns 设置空闲连接池中连接的最大数量
	dbConn.SetMaxIdleConns(config.GetPgsqlConfig().GetMaxIdleConnection())

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	dbConn.SetConnMaxLifetime(time.Hour)

	// 激活链接
	if err = dbConn.Ping(); err != nil {
		log.Fatal(err)
	}
}
