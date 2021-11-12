package log

import (
	"go.uber.org/zap/zapcore"
	"os"
	"sync"

	"go.uber.org/zap"
)

var (
	sugarLogger *zap.SugaredLogger
	once        = &sync.Once{}
)

// GetLogger 获取日志对象
func GetLogger() *zap.SugaredLogger {
	if sugarLogger == nil {
		once.Do(func() {
			encoderConfig := zapcore.EncoderConfig{
				TimeKey:        "time",
				LevelKey:       "level",
				NameKey:        "logger",
				CallerKey:      "caller",
				MessageKey:     "msg",
				StacktraceKey:  "stacktrace",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
				EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
				EncodeDuration: zapcore.SecondsDurationEncoder,
				EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
			}

			// // 设置日志级别
			// atom := zap.NewAtomicLevelAt(zap.DebugLevel)
			// config := zap.Config{
			// 	Level:            atom,                                          // 日志级别
			// 	Development:      true,                                          // 开发模式，堆栈跟踪
			// 	Encoding:         "json",                                        // 输出格式 console 或 json
			// 	EncoderConfig:    encoderConfig,                                 // 编码器配置
			// 	InitialFields:    map[string]interface{}{"serviceName": "init"}, // 初始化字段，如：添加一个服务器名称
			// 	OutputPaths:      []string{"stdout", "init.log"},                // 输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）
			// 	ErrorOutputPaths: []string{"stderr"},
			// }

			enConfig := zapcore.NewJSONEncoder(encoderConfig)
			file, _ := os.OpenFile("init.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
			writer := zapcore.AddSync(file)

			core := zapcore.NewCore(enConfig, writer, zapcore.InfoLevel)
			Logger := zap.New(core, zap.AddCaller())
			sugarLogger = Logger.Sugar()
			Logger.Info("log 初始化成功")

			defer Logger.Sync() // flushes buffer, if any
		})
	}
	return sugarLogger
}
