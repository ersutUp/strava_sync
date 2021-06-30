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