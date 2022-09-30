package main

import (
	"blog/dao"
	"blog/logger"
	"blog/models"
	"blog/routers"
	"blog/settings"
	"fmt"
)

func main() {
	//加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("settings Init error : %v", err)
		return
	}

	//加载日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("logger init err: %v", err)
		return
	}

	//加载mysql
	if err := dao.InitMysql(settings.Conf.MysqlConfig); err != nil {
		fmt.Printf("mysql init err: %v", err)
		return
	}
	defer dao.Close()

	//模型绑定 model => 数据表
	dao.DB.AutoMigrate(
		new(models.User),
		new(models.Comment),
		new(models.Category),
		new(models.Config),
		new(models.Post),
	)

	//注册路由
	r := routers.SetRouter()
	if err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port)); err != nil {
		fmt.Printf("server start err: %v", err)
		return
	}
}
