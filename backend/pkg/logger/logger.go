package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

// InitLogger 初始化日志
func InitLogger() error {
	// 设置日志写入文件
	logWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./logs/app.log", // 日志文件路径
		MaxSize:    128,              // 日志文件最大大小(MB)
		MaxBackups: 30,               // 保留旧文件的最大个数
		MaxAge:     7,                // 保留旧文件的最大天数
		Compress:   true,             // 是否压缩旧文件
	})

	// 设置日志编码
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     customTimeEncoder,              // 自定义时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 短路径编码器
	}

	// 创建encoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// 设置日志级别
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel && lvl >= zapcore.DebugLevel
	})

	// 创建core
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, logWriter, highPriority),
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), lowPriority),
	)

	// 创建logger
	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return nil
}

// 自定义时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}