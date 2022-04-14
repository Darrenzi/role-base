package main

import (
	"blog/common/global"
	"blog/router"
	"flag"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var conf string

func main() {
	flag.StringVar(&conf, "conf", "", "configuration file path")
	flag.Parse()

	if conf == "" {
		conf = "D:/blog_2022/blog-server/config/"
	}
	if err := global.ParseConfig(conf); err != nil {
		log.Error("解析配置文件失败:", err)
	}

	global.Init()

	host := fmt.Sprintf("localhost:%s", global.Config.Port)
	router := router.InitRouter()
	server := &http.Server{
		Addr:    host,
		Handler: router,
	}

	log.Info("启动服务器, 服务地址:", host)
	if err := server.ListenAndServe(); err != nil {
		log.Error("启动服务器失败: ", err)
	}
}
