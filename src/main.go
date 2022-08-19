package main

import (
	"flag"
	"webhook/src/global"
	"webhook/src/global/http"
	"webhook/src/logger"
)

func main() {
	// 初始化配置
	global.InitConfig()
	// 初始化日志
	logger.InitLogger()
	// 启动Banner
	global.Welcome()
	// 初始化指令
	global.InitCommand()
	// 解析指令
	flag.Parse()
	// Http 服务
	http.Start()
}
