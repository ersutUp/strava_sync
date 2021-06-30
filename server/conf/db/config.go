package db

import "github.com/beego/beego/v2/server/web"

//获取配置出现错误时 panic
func getAppConfig(key string) string {
	if val, err := web.AppConfig.String(key); err != nil {
		panic("\"" + key + "\" not config")
	} else {
		return val
	}
}
