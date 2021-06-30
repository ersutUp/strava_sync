package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "fit_sync_server/conf/filters"
	"fit_sync_server/conf/log"
	_ "fit_sync_server/routers"
)

func main() {
	//logrus 日志配置
	log.LogrusConf()
	//beego.logs 配置
	log.BeggoConf()

	var mode = beego.BConfig.RunMode
	if mode == beego.DEV {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

	} else if mode == beego.PROD {

	}

	beego.Run()
}
