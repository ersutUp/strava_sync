//logurs的配置
package log

import (
	log "github.com/sirupsen/logrus"
	"os"
	"fit_sync_server/utils"
)

const (
	LogForm = "log_from"
	Beggo   = "beggo"
	Gorm    = "gorm"
)

//logrus 日志配置
func LogrusConf() {

	if utils.IsProd() {
		//json 格式输出
		/*
			添加前
				time="2020-12-31T16:57:23+08:00" level=info msg="hello world"
			添加后
				{"level":"info","msg":"hello world","time":"2020-12-31T16:53:01+08:00"}


			FieldMap 更改key值，如下（log.FieldMap{log.FieldKeyLevel:"l"}）
				{"l":"info","msg":"hello world","time":"2020-12-31T17:16:16+08:00"}

			TimestampFormat 格式化日期 ，如下（TimestampFormat:"2006-01-02 15:04:05.000"）
				{"l":"info","msg":"hello world","time":"2020-12-31 17:20:55.945"}

			PrettyPrint 格式化输出，如下（PrettyPrint:true）
				{
				  "l": "info",
				  "msg": "hello world",
				  "time": "2020-12-31 17:22:38.800"
				}

			DataKey 将 log.Fields 放在其中
				添加前
					{"num":1,"l":"info","msg":"hello world","time":"2020-12-31 17:38:10.566"}
				添加后（DataKey:"data"）
					{"data":{"num":1},"l":"info","msg":"hello world","time":"2020-12-31 17:40:28.509"}

		*/
		log.SetFormatter(&log.JSONFormatter{DataKey: "data", TimestampFormat: "2006-01-02 15:04:05.000"})
	} else {
		//纯文本输出
		/*
			ForceColors 展示颜色
			FullTimestamp 展示完整时间戳
			TimestampFormat 格式化时间
		*/
		log.SetFormatter(&log.TextFormatter{FullTimestamp: true, ForceColors: false, DisableColors: true, TimestampFormat: "2006-01-02 15:04:05.000"})
	}
	/*
		展示具体打印日志的文件路径、行号以及方法，分别在file、func中展示
		{"file":"D:/code/go/go_student/log/log.go:37","func":"go_student/log.Print","l":"info","msg":"hello world","time":"2020-12-31 17:28:28.152"}
	*/
	log.SetReportCaller(false)

	//最低可打印的日志等级
	if utils.IsProd() {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.DebugLevel)
	}

	//输出到文件
	if !utils.IsDev() {
		file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			log.SetOutput(file)
		} else {
			log.Info("Failed to log to file, using default stderr")
		}
	}

	//添加hook
	log.AddHook(CallerHook{})
}
