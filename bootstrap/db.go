package bootstrap

import (
	"errors"
	"fmt"
	"github.com/gin-ctl/zero/package/database"
	"github.com/gin-ctl/zero/package/get"
	"github.com/gin-ctl/zero/package/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDB() {
	var dbConfig gorm.Dialector
	switch get.String("db.connection") {
	case "mysql":
		// 构建 DSN 信息
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
			get.String("db.mysql.username"),
			get.String("db.mysql.password"),
			get.String("db.mysql.host"),
			get.String("db.mysql.port"),
			get.String("db.mysql.database"),
			get.String("db.mysql.charset"),
		)
		dbConfig = mysql.New(mysql.Config{
			DSN:                       dsn,
			SkipInitializeWithVersion: get.Bool("db.mysql.skip_initialize_with_version"),
		})
	case "tidb":
		// 构建 DSN 信息
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&tls=%s",
			get.String("db.tidb.username"),
			get.String("db.tidb.password"),
			get.String("db.tidb.host"),
			get.String("db.tidb.port"),
			get.String("db.tidb.database"),
			get.String("db.tidb.ssl", false),
		)
		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	//case "sqlite":
	//	dbConfig = sqlite.Open(config.Get("db.sqlite.database"))
	default:
		panic(errors.New("database connection not supported"))
	}

	// 连接数据库，并设置 GORM 的日志模式
	err := database.Connect(dbConfig, logger.NewGormLogger())
	if err != nil {
		panic(errors.New("database connection failure"))
	}
}
