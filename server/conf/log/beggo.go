package log

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

//beggo 的日志配置
func BeggoConf() {
	//开始处理日志
	logs.Reset()

	//var pattern *logs.PatternLogFormatter
	//pattern = &logs.PatternLogFormatter{
	//	Pattern:    "{time:\"%w\",level:\"%t\",file:\"%f\",line:\"%n\",msg:\"%m\"}\r\n",
	//	//Pattern:"%F:%n|%w%t>> %m\r\n",
	//	WhenFormat: "2006-01-02 15:04:05.000",
	//}
	////注册打印格式
	//logs.RegisterFormatter("pattern", pattern)
	////设置全局打印格式
	//logs.SetGlobalFormatter("pattern")

	//文件以及行号
	beego.BConfig.Log.FileLineNum = false

	/*//所有日志输出到文件中
	logs.SetLogger(logs.AdapterFile,`{"filename":"beego.log","maxdays":10}`)
	//错误日志输出到文件中
	logs.SetLogger(logs.AdapterMultiFile,`{"filename":"beego.error.log","maxdays":10,"separate":["emergency", "alert", "critical", "error"]}}`)
	//打印到控制台
	logs.SetLogger(logs.AdapterConsole)*/

	//beggo的日志托管给logrus
	logs.SetLogger(Beggo2logrus)

}
