package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	server   = "127.0.0.1"
	port     = "3306"
	username = "root"
	password = ""
	database = "order"
)
var DB = Conn()

func Conn() *gorm.DB {

	MysqlDB, _ := gorm.Open(mysql.New(mysql.Config{
		DSN:                       GetMysqlDsn(), // DSN data source name
		DefaultStringSize:         191,           // string 类型字段的默认长度
		DisableDatetimePrecision:  true,          // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,          // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,          // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,         // 根据版本自动配置
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   Logger(),
	})

	return MysqlDB.Debug()
}

func GetMysqlDsn() string {
	mysqlConn := username + ":" + password + "@tcp(" + server + ":" + port + ")/" + database + "?charset=utf8mb4,utf8&parseTime=True&loc=Local"
	return mysqlConn
}

func Logger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,  // 慢 SQL 阈值
			IgnoreRecordNotFoundError: true,         // 忽视
			LogLevel:                  logger.Error, // Log level
			Colorful:                  true,         // 彩色打印
		},
	)
}
