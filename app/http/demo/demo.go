package main

import (
	"fmt"
	"github.com/gin-ctl/zero/package/bootstrap"
	"github.com/gin-ctl/zero/package/get"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// 加载配置文件
	get.NewViper("env.yaml", "./config")

	// 启动基础服务
	bootstrap.SetupLogger()
	bootstrap.SetupDB()
	bootstrap.SetupRedis()

	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，适用于生产环境
	gin.SetMode(gin.ReleaseMode)
	// gin 实例
	router := gin.New()
	// 初始化路由绑定
	bootstrap.RegisterDemoApiRoute(router)
	// 运行http服务
	log.Println("demo api serve start: " + get.String("app.host") +
		":" + get.String("app.port"))
	err := router.Run(fmt.Sprintf("%s:%d",
		get.Get("app.host"), get.Int("app.port")))
	if err != nil {
		panic("Unable to start server, error: " + err.Error())
	}
}
