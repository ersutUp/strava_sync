package utils

import beego "github.com/beego/beego/v2/server/web"

func IsProd() bool {
	var mode = beego.BConfig.RunMode
	if mode == "" || beego.PROD == mode {
		return true
	}
	return false
}
func IsDev() bool {
	var mode = beego.BConfig.RunMode
	if beego.DEV == mode {
		return true
	}
	return false
}


//获取配置出现错误时 panic
func GetAppConfig(key string) string {
	if val, err := beego.AppConfig.String(key); err != nil {
		panic("\"" + key + "\" not config")
	} else {
		return val
	}
}