package main

import (
	"bluebell/domain/mysql"
	"bluebell/internal/logger"
	"bluebell/internal/setting"
	"bluebell/routers"
	"fmt"
	"log"
)

func main() {
	//加载配置
	if err := setting.Init(); err != nil {
		log.Fatalln("配置文件初始化失败", err)
		return
	}
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		log.Fatalln("日志初始化失败", err)
		return
	}
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		log.Fatalln("mysql初始化失败", err)
		return
	}
	defer mysql.Close() // 关闭数据库连接

	//注册路由
	r := routers.SetupRouter(setting.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
