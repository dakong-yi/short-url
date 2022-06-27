package mysql

import (
	"fmt"
	"short-url/config"

	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlClient *gorm.DB

func InitMysql() {
	mysqlConfig := config.Cfg.MysqlCfg
	// user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database, mysqlConfig.Charset,
		mysqlConfig.ParseTime, mysqlConfig.TimeZone)
	var err error
	MysqlClient, err = gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic(fmt.Sprintf("创建mysql客户端失败: %v,%s", MysqlClient, err))
	}
}

func GetMysqlClient() *gorm.DB {
	return MysqlClient
}
