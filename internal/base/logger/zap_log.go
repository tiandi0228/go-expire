package logger

import (
	opfile "hongcha/go-expire/pkg/opfile"
	"os"
	"path"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitZap logPath: 日志打印目录
// maxAge: 日志最大存在时间，单位：天
// rotationTime: 日志切分时间，单位：小时
// projectName: 项目名称
func InitZap(projectName, logPath string, maxAge, rotationTime time.Duration, stdout bool) *zap.Logger {
	if len(projectName) == 0 {
		panic("logger init fail, project name is empty")
	}

	// 创建日志存放目录
	if err := opfile.CreateDirIfNotExist(logPath); err != nil {
		panic(err)
	}
	logPath = path.Join(logPath, projectName)

	// error日志文件配置
	errWriter, err := rotatelogs.New(
		logPath+"_err_%Y-%m-%d.log",
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)
	if err != nil {
		panic(err)
	}

	// info日志文件配置
	infoWriter, err := rotatelogs.New(
		logPath+"_info_%Y-%m-%d.log",
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)
	if err != nil {
		panic(err)
	}

	// 优先级设置（一个日志输出全部信息，一个日志输出error信息）
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})

	// 控制台输出设置
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleEncoderConfig := zap.NewDevelopmentEncoderConfig()
	consoleEncoderConfig.EncodeTime = timeEncoder
	consoleEncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoderConfig.EncodeCaller = customCallerEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)

	// 文件输出设置
	errorCore := zapcore.AddSync(errWriter)
	infoCore := zapcore.AddSync(infoWriter)
	fileEncodeConfig := zap.NewProductionEncoderConfig()
	fileEncodeConfig.EncodeTime = timeEncoder
	fileEncodeConfig.EncodeCaller = customCallerEncoder
	fileEncoder := zapcore.NewConsoleEncoder(fileEncodeConfig)

	cores := make([]zapcore.Core, 0)
	cores = append(cores, zapcore.NewCore(fileEncoder, errorCore, highPriority))
	cores = append(cores, zapcore.NewCore(fileEncoder, infoCore, lowPriority))
	// 如果需要控制台输出则加入cores中
	if stdout {
		cores = append(cores, zapcore.NewCore(consoleEncoder, consoleDebugging, zapcore.DebugLevel))
	}
	core := zapcore.NewTee(cores...)

	// 显示行号
	caller := zap.AddCaller()
	// 设置打印堆栈深度
	callerSkip := zap.AddCallerSkip(2)

	development := zap.Development()
	logger := zap.New(core, caller, callerSkip, development)

	// 替换全局日志
	zap.ReplaceGlobals(logger)

	// 将系统输出重定向到zap中，保证所有出现异常均能打印到文件中
	if _, err := zap.RedirectStdLogAt(logger, zapcore.ErrorLevel); err != nil {
		panic(err)
	}

	return logger
}

// InitZapConsole 初始化控制台日志用于单元测试
func InitZapConsole() *zap.Logger {
	// 控制台输出设置
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleEncoderConfig := zap.NewDevelopmentEncoderConfig()
	consoleEncoderConfig.EncodeTime = timeEncoder
	consoleEncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoderConfig.EncodeCaller = customCallerEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)

	// 显示行号
	caller := zap.AddCaller()
	// 设置打印堆栈深度
	callerSkip := zap.AddCallerSkip(2)

	development := zap.Development()
	zapLogger := zap.New(zapcore.NewCore(consoleEncoder, consoleDebugging, zapcore.DebugLevel), caller, callerSkip, development)

	// 替换全局日志
	zap.ReplaceGlobals(zapLogger)

	// 将系统输出重定向到zap中，保证所有出现异常均能打印到文件中
	if _, err := zap.RedirectStdLogAt(zapLogger, zapcore.ErrorLevel); err != nil {
		panic(err)
	}

	return zapLogger
}

// 自定义打印路径，减少输出日志打印路径长度，根据输入项目名进行减少
func customCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	str := caller.String()
	index := strings.Index(str, "src")
	if index == -1 {
		enc.AppendString(caller.FullPath())
	} else {
		index = index + len("src") + 1
		enc.AppendString(str[index:])
	}
}

// 格式化日志时间，官方的不好看
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
