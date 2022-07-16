package main

import (
	"drive/config"
	"drive/db"
	"drive/logger"
	"drive/routers"
	"drive/services"
	"drive/tools"
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {
	// 初始化日志
	logger.InitLog()

	// 初始化配置文件
	err := config.Init()
	if err != nil {
		return
	}

	// 初始化通用工具：例如雪花 ID 生成器
	err = tools.Init()
	if err != nil {
		return
	}

	// 数据库准备工作：例如迁移
	db.Init()

	// 业务服务层初始化工作：例如建立数据库连接供各个业务使用
	services.Init()

	// 初始化路由
	r := routers.Init()
	err = r.Run(fmt.Sprintf(":%d", config.Conf.APIConfig.Port))
	if err != nil {
		logrus.Errorf("run server failed, err: %v", err)
		return
	}
}
