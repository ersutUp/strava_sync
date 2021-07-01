package controllers

import (
	"fit_sync_server/conf/db"
	"fit_sync_server/models"
	beego "github.com/beego/beego/v2/server/web"
)

type UserInfoController struct {
	beego.Controller
}

// @router / [get]
func (this *UserInfoController) Get()  {
	userInfo := &models.UserInfo{}

	db := db.Mydb.Last(userInfo)

	err := db.Error
	if err != nil {
		panic("db err")
	}

	this.Data["json"] = userInfo

	this.ServeJSON()
}