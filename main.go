package main

import (
	"mango-user-center/config"
	"mango-user-center/internal/controller"
	"mango-user-center/pkg/db"
	"net/http"
	"time"
)

func main() {

	// 连接数据库
	db.Connect()

	// 启动HTTP服务
	s := &http.Server{
		Addr:           ":" + config.App.Port,
		Handler:        controller.Engine,
		WriteTimeout:   6 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1M
	}
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
