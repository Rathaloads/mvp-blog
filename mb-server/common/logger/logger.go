package logger

import (
	"io"
	"os"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	LOG_TIME_FORMAT = "2006-01-02 15:04:05.999" // 日志时间格式:到毫秒
)

var (
	logger *zap.SugaredLogger

	Debug  func(a ...interface{})
	Debugf func(format string, a ...interface{})
	Info   func(a ...interface{})
	Infof  func(format string, a ...interface{})
	Warn   func(a ...interface{})
	Warnf  func(format string, a ...interface{})
	Error  func(a ...interface{})
	Errorf func(format string, a ...interface{})
	Fatal  func(a ...interface{})
	Fatalf func(format string, a ...interface{})
	Panic  func(a ...interface{})
	Panicf func(format string, a ...interface{})
)

func InitLog(log_path string) {
	getLogger(log_path)

	Debug = logger.Debug
	Debugf = logger.Debugf
	Info = logger.Info
	Infof = logger.Infof
	Warn = logger.Warn
	Warnf = logger.Warnf
	Error = logger.Error
	Errorf = logger.Errorf
	Fatal = logger.Fatal
	Fatalf = logger.Fatalf
	Panic = logger.Panic
	Panicf = logger.Panicf
}

func getLogger(log_path string) {
	// 文件log格式
	encoder := getEncoder()

	// 控制台log格式
	console_encoder := getConsoleEncoder()

	// 输出文件：不同等级分文件存放
	info_writer := getWriter(log_path + "/info/info")
	info_level := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})
	warn_writer := getWriter(log_path + "/warn/warn")
	warn_level := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})
	error_writer := getWriter(log_path + "/error/error")
	error_level := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	// 创建logger
	newCore := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(info_writer), info_level),            // 写入INFO文件
		zapcore.NewCore(encoder, zapcore.AddSync(warn_writer), warn_level),            // 写入WRAN文件
		zapcore.NewCore(encoder, zapcore.AddSync(error_writer), error_level),          // 写入ERROR文件
		zapcore.NewCore(console_encoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel), // 写入控制台
	)

	log := zap.New(newCore, zap.AddCaller())
	logger = log.Sugar()
}

// 按要求写文件
func getWriter(filename string) io.Writer {
	hook, err := rotatelogs.New(
		strings.Replace(filename, ".log", "", -1)+"-%Y%m%d.log", // 文件名format格式，一天一个文件
		rotatelogs.WithLinkName(filename),                       // 文件名
		rotatelogs.WithMaxAge(time.Hour*24*15),                  // 保存时间
		rotatelogs.WithRotationTime(time.Hour),                  // 多长时间轮询
	)

	if err != nil {
		panic(err)
	}

	return hook
}

// 自定义的输出格式
func getEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(
		zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller_line",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     "\n",
			EncodeLevel:    cEncodeLevel,
			EncodeTime:     cEncodeTime,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   cEncodeCaller,
		})
}

// 输出日志到控制台
func getConsoleEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
}

// 自定义日志级别显示
func cEncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

// 自定义时间格式显示
func cEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + t.Format(LOG_TIME_FORMAT) + "]")
}

// 自定义行号显示
func cEncodeCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + caller.TrimmedPath() + "]")
}
