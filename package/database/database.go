package database

import (
	"github.com/gin-ctl/zero/package/get"
	"github.com/gin-ctl/zero/package/logger"
	"gorm.io/gorm"
	gl "gorm.io/gorm/logger"
	"sync"
	"time"
)

var (
	DB   *gorm.DB
	Once sync.Once
)

// Connect 连接数据库
func Connect(dbConfig gorm.Dialector, _logger gl.Interface) (err error) {
	// 使用 gorm.Open 连接数据库
	Once.Do(func() {
		DB, err = gorm.Open(dbConfig, &gorm.Config{
			Logger: _logger,
		})
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
