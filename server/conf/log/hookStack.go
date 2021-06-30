//logurs的hook：error级以上打印堆栈信息
package log

import (
	"github.com/sirupsen/logrus"
	"runtime"
)

type CallerHook struct {
}

//那些等级下执行此hook
func (h CallerHook) Levels() []logrus.Level {
	//return logrus.AllLevels
	return []logrus.Level{logrus.ErrorLevel}
}

//hook具体处理的事情
func (h CallerHook) Fire(entry *logrus.Entry) error {
	buf := make([]byte, 1<<20)
	len := runtime.Stack(buf, true)
	if len > 0 {
		s := string(buf[:len])
		entry.Data["stack"] = s
	}
	return nil
}
