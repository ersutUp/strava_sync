package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/prometheus/common/log"
)

type ErrorController struct {
	web.Controller

}

func (c ErrorController) Error503()  {
	log.Info("NotLogin")
}