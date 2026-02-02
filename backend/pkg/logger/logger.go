package logger

import (
	"backend/config"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Logger    *zap.Logger
	logWriter *lumberjack.Logger
)

// InitLogger 初始化日志
func InitLogger() error {
	// 获取配置中的日志设置，如果未设置则使用默认值
	logFile := "./logs/app.log"
	maxSize := 128
	maxBackups := 30
	maxAge := 7
	compress := true

	// 如果全局配置已初始化，则使用配置文件中的值
	if config.GlobalConfig != nil {
		if config.GlobalConfig.Server.LogFile != "" {
			logFile = config.GlobalConfig.Server.LogFile
		}
		if config.GlobalConfig.Server.MaxSize > 0 {
			maxSize = config.GlobalConfig.Server.MaxSize
		}
		if config.GlobalConfig.Server.MaxBackups > 0 {
			maxBackups = config.GlobalConfig.Server.MaxBackups
		}
		if config.GlobalConfig.Server.MaxAge > 0 {
			maxAge = config.GlobalConfig.Server.MaxAge
		}
		compress = config.GlobalConfig.Server.Compress
	}

	// 创建lumberjack logger实例
	logWriter = &lumberjack.Logger{
		Filename:   logFile,    // 日志文件路径
		MaxSize:    maxSize,    // 日志文件最大大小(MB)
		MaxBackups: maxBackups, // 保留旧文件的最大个数
		MaxAge:     maxAge,     // 保留旧文件的最大天数
		Compress:   compress,   // 是否压缩旧文件
	}

	// 设置日志写入文件
	writerSync := zapcore.AddSync(logWriter)

	// 设置日志编码
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     customTimeEncoder,             // 自定义时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 短路径编码器
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
		zapcore.NewCore(encoder, writerSync, highPriority),
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), lowPriority),
	)

	// 创建logger
	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	// 确保日志被刷新到文件
	zap.RedirectStdLog(Logger)

	return nil
}

// Sync 强制将缓冲的日志写入文件
func Sync() {
	if Logger != nil {
		Logger.Sync()
	}
	if logWriter != nil {
		logWriter.Close()
	}
}

// 自定义时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
