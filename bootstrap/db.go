package bootstrap

import (
	"errors"
	"fmt"
	"github.com/gin-ctl/zero/package/get"
	"github.com/gin-ctl/zero/package/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gl "gorm.io/gorm/logger"
	"time"
)

var (
	DB *gorm.DB
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
	err := connect(dbConfig, logger.NewGormLogger())
	if err != nil {
		panic(errors.New("database connection failure"))
	}
}

// Connect 连接数据库
func connect(dbConfig gorm.Dialector, _logger gl.Interface) (err error) {
	// 使用 gorm.Open 连接数据库
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})

	// 处理错误
	if err != nil {
		logger.ErrorJSON("database", "connect", err)
		return err
	}

	if get.String("db.connection") == "mysql" {
		// 获取底层的 SqlDB
		sqlDB, errs := DB.DB()
		if errs != nil {
			logger.ErrorJSON("database", "connect", err)
			return errs
		}
		// 设置最大连接数
		sqlDB.SetMaxOpenConns(get.Int("db.mysql.max_open_connections"))
		// 设置最大空闲连接数
		sqlDB.SetMaxIdleConns(get.Int("db.mysql.max_idle_connections"))
		// 设置每个连接的过期时间
		sqlDB.SetConnMaxLifetime(time.Duration(get.Int("db.mysql.max_life_seconds")) * time.Second)
	}
	return nil
}
