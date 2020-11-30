package main

import (
	"fmt"
	"go.uber.org/zap"
	"single-login/dao/mysql"
	"single-login/dao/redis"
	"single-login/logger"
	"single-login/router"
	"single-login/setting"
)

func main() {

	//1. 加载配置
	err := setting.Init()
	if err != nil {
		zap.L().Debug("Init settings failed", zap.Error(err))
		fmt.Printf("错误:%v\n", err)
		return
	}
	//2. 初始化日志
	err = logger.Init(setting.Cfg.LoggerConf, setting.Cfg.Mode)
	if err != nil {
		zap.L().Debug("Init log failed", zap.Error(err))
		return
	}

	//将缓冲区的日志追加到文件中
	defer zap.L().Sync()

	//3. 初始化Mysql
	err = mysql.Init(setting.Cfg.MysqlConf)
	if err != nil {
		zap.L().Debug("Init mysql failed", zap.Error(err))
		return
	}
	defer mysql.Close()

	//4. 初始化Redis
	err = redis.Init(setting.Cfg.RedisConf)
	if err != nil {
		zap.L().Debug("Init redis failed", zap.Error(err))
		return
	}
	defer redis.Close()

	//5. 注册路由
	r := router.Setup()
	_ = r.Run(":8888")

}
