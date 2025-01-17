package marsLog

import "context"

func init() {
	// 如果没有设置Logger, 则启动时使用默认的MarsDefaultLog对象
	if Logger() == nil {
		SetLogger(&DefaultLog{})
	}
}

type MarsLogger interface {
	// Debug 无上下文的Info级别日志接口, format字符串格式
	Debug(v ...interface{})
	DebugF(str string, v ...interface{})
	Info(v ...interface{})
	InfoF(str string, v ...interface{})
	Warn(v ...interface{})
	WarnF(str string, v ...interface{})
	Error(v ...interface{})
	ErrorF(str string, v ...interface{})
	Fatal(v ...interface{})
	FatalF(str string, v ...interface{})

	// Json 打印json
	Json(v interface{})
	JsonFormat(v interface{})

	// DebugFX 有上下文的Info级别日志接口, format字符串格式
	DebugFX(ctx context.Context, str string, v ...interface{})
	InfoFX(ctx context.Context, str string, v ...interface{})
	WarnFX(ctx context.Context, str string, v ...interface{})
	ErrorFX(ctx context.Context, str string, v ...interface{})
	FatalFX(ctx context.Context, str string, v ...interface{})
}

// MarsLog 默认的KisLog 对象
var MarsLog MarsLogger

// SetLogger 设置KisLog对象, 可以是用户自定义的Logger对象
func SetLogger(newLog MarsLogger) {
	MarsLog = newLog
}

// Logger 获取到kisLog对象
func Logger() MarsLogger {
	return MarsLog
}
