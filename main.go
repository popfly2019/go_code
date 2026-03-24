package main

import (
	"go_demo/config"
	"go_demo/pkg/db"
	"go_demo/router"
	"log"
)

func main() {

	if err := config.InitConfig(); err != nil {
		log.Fatalf("初始化配置失败: %v", err)
	}

	if err := db.InitDB(&config.AppConfig.Database); err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}

	defer db.CloseDB()

	log.Printf("服务器启动在端口: %s", config.AppConfig.Port)

	r := router.Setup()

	_ = r.Run()
}
