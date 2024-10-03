package common

import (
	"github.com/sirupsen/logrus"
	"path"
)

var (
	logPath = "../logs"
	logFile = "access.log"
)
var Logger = logrus.New()

type RequestIdHook struct {
	RequestId string
}

func NewRequestIdHook(requestId string) logrus.Hook {
	hook := RequestIdHook{
		RequestId: requestId,
	}
	return &hook
}

func (hook *RequestIdHook) Fire(entry *logrus.Entry) error {
	entry.Data["traceId"] = hook.RequestId
	return nil
}

func (hook *RequestIdHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// 日志初始化
func init() {
	// 打开文件
	logFileName := path.Join(logPath, logFile)
	// 使用滚动压缩方式记录日志
	rolling(logFileName)
	// 设置日志输出JSON格式
	/*Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05", // 设置json里的日期输出格式
	})*/
	// 设置日志输出Text格式
	Logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat:           "2006-01-02 15:04:05",
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
		DisableLevelTruncation:    true,
		QuoteEmptyFields:          true,
	})
	// 设置日志记录级别
	Logger.SetLevel(logrus.DebugLevel)
}

// 日志滚动设置
func rolling(logFile string) {
	// 设置输出
	Logger.SetOutput(&lumberjack.Logger{
		Filename:   logFile, //日志文件位置
		MaxSize:    500,     // 单文件最大容量,单位是MB
		MaxBackups: 3,       // 最大保留过期文件个数
		MaxAge:     28,      // 保留过期文件的最大时间间隔,单位是天
		Compress:   true,    // 是否需要压缩滚动日志, 使用的 gzip 压缩
	})
}
