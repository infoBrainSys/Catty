package main

import (
	"GPT/config"
	"GPT/dao"
	"GPT/routers"
)

func main() {
	// 初始化 viper 读取配置文件
	config.Config()
	// 初始化数据库
	dao.Initdatabase()
	// 初始化路由
	app := routers.Routers()

	app.Run(":80")
}
