package main

import (
	"GinStudy/global"
	"GinStudy/router"
	"github.com/gin-gonic/gin"
	"log"
)

// 初始化 工作
func init() {
	err := global.SetupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = global.SetupDB()
	if err != nil {
		log.Fatalf("init.setupDB err: %v", err)
	}
	err = global.SetupRedisDb()
	if err != nil {
		log.Fatalf("init.setupRedisDb err: %v", err)
	}
}
func main() {
	//设置运行模式
	gin.SetMode(global.ServerSetting.RunMode)
	//引入路由
	r := router.Router()
	//启动
	r.Run(global.ServerSetting.HttpHost + ":" + global.ServerSetting.HttpPort)
}
