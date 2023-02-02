package logger

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
)

// myLogger 日志引用对象
var myLogger Logger

func init() {
	myLogger = InitLoggerForTest()
}

// SetLogger 设置 logger 对象
func SetLogger(log Logger) {
	myLogger = log
}

// Debug log
func Debug(v ...interface{}) {
	myLogger.Debug(v...)
}

// Debugf log
func Debugf(format string, v ...interface{}) {
	myLogger.Debugf(format, v...)
}

// Info log
func Info(v ...interface{}) {
	myLogger.Info(v...)
}

// Infof log
func Infof(format string, v ...interface{}) {
	myLogger.Infof(format, v...)
}

// Warn log
func Warn(v ...interface{}) {
	myLogger.Warn(v...)
}

// Warnf log
func Warnf(format string, v ...interface{}) {
	myLogger.Warnf(format, v...)
}

// Error log
func Error(v ...interface{}) {
	myLogger.Error(v...)
}

// Errorf log
func Errorf(format string, v ...interface{}) {
	myLogger.Errorf(format, v...)
}

// LogStack 返回当前调用堆栈信息
// start 起始调用栈层级
// end 结束调用栈层级 输入0则会添加调用栈信息直到没有
func LogStack(start, end int) string {
	stack := bytes.Buffer{}
	for i := start; i < end || end == 0; i++ {
		pc, str, line, _ := runtime.Caller(i)
		if line == 0 {
			break
		}

		// 根据src截短输出路径
		index := strings.Index(str, "src")
		if index != -1 {
			index = index + len("src") + 1
			str = str[index:]
		}
		stack.WriteString(fmt.Sprintf("%s:%d %s\n", str, line, runtime.FuncForPC(pc).Name()))
	}
	return stack.String()
}
