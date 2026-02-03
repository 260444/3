package logger

import (
	"backend/config"
	"log"
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

// customTimeEncoder 自定义时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// InitLogger 初始化并返回一个基于配置设置的新 zap.Logger 实例
func InitLogger() error {
	zapCfg := config.GlobalConfig.Server

	// 创建一个用于日志输出的 writeSyncer
	writeSyncer := getLogWriter(zapCfg.LogFile, zapCfg.MaxSize, zapCfg.MaxBackups, zapCfg.MaxAge)

	// 如果配置了控制台输出，则添加控制台输出
	if zapCfg.IsConsole {
		writeSyncer = zapcore.NewMultiWriteSyncer(writeSyncer, zapcore.AddSync(os.Stdout))
	}

	// 创建日志格式化的编码器
	encoder := getEncoder()

	// 根据配置确定日志级别
	var logLevel zapcore.Level

	if err := logLevel.UnmarshalText([]byte(zapCfg.Mode)); err != nil {
		log.Fatalf("Failed to parse log level: %v", err)
	}

	// 创建核心和日志实例
	core := zapcore.NewCore(encoder, writeSyncer, logLevel)
	Logger = zap.New(core, zap.AddCaller())
	return nil
}

// getLogWriter 返回一个 zapcore.WriteSyncer，该写入器利用 lumberjack 包，实现日志的滚动记录
func getLogWriter(filename string, maxSize, maxBackups, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,   // 日志文件的位置
		MaxSize:    maxSize,    // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: maxBackups, // 保留旧文件的最大个数
		MaxAge:     maxAge,     // 保留旧文件的最大天数
	}
	return zapcore.AddSync(lumberJackLogger)
}

// getEncoder 返回一个为生产日志配置的 JSON 编码器
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	encoderConfig.TimeKey = "time"
	encoderConfig.LevelKey = "level"
	encoderConfig.NameKey = "logger"
	encoderConfig.CallerKey = "caller"
	encoderConfig.MessageKey = "msg"
	encoderConfig.StacktraceKey = "stacktrace"
	// 小写编码器
	encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder

	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	// 短路径编码器
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	// 自定义时间格式
	encoderConfig.EncodeTime = customTimeEncoder
	// encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
