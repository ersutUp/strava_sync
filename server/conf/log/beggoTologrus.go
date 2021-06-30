//beggo的日志对接到logurs
package log

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/sirupsen/logrus"
)

/*
type Logger interface {
	//实现类的初始化方法，一般为优先的一些时间 比如打开文件、创建客户端等等
	Init(config string) error
	//每次打印日志时调用的方法
	WriteMsg(lm *LogMsg) error
	Destroy()
	Flush()
	//全局统一格式
	SetFormatter(f LogFormatter)
}

func (bl *BeeLogger) setLogger(adapterName string, configs ...string) error {

	...

	//全局统一格式
	// Global formatter overrides the default set formatter
	if len(bl.globalFormatter) > 0 {
		fmtr, ok := GetFormatter(bl.globalFormatter)
		if !ok {
			return errors.New(fmt.Sprintf("the formatter with name: %s not found", bl.globalFormatter))
		}
		lg.SetFormatter(fmtr)
	}

	//Logger 实现类的初始化方法，一般为优先的一些时间 比如打开文件、创建客户端等等
	err := lg.Init(config)

	...
}
*/

const Beggo2logrus = "beggo2logrus"

func init() {
	logs.Register(Beggo2logrus, func() logs.Logger {
		return &BeggoLog2logrus{}
	})
}

type BeggoLog2logrus struct{}

func (b *BeggoLog2logrus) Init(config string) error {
	return nil
}
func (b *BeggoLog2logrus) Destroy()                         {}
func (b *BeggoLog2logrus) Flush()                           {}
func (b *BeggoLog2logrus) SetFormatter(f logs.LogFormatter) {}

func (b *BeggoLog2logrus) WriteMsg(lm *logs.LogMsg) error {
	level := lm.Level
	msg := lm.Msg
	switch level {
	case logs.LevelEmergency, logs.LevelAlert, logs.LevelCritical, logs.LevelError:
		logrus.WithField(LogForm, Beggo).Error(msg)
	case logs.LevelWarning, logs.LevelNotice:
		logrus.WithField(LogForm, Beggo).Warn(msg)
	case logs.LevelInformational:
		logrus.WithField(LogForm, Beggo).Info(msg)
	case logs.LevelDebug:
		logrus.WithField(LogForm, Beggo).Debug(msg)
	default:
		logrus.WithField(LogForm, Beggo).Trace(msg)
	}
	return nil
}
