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
	Debug(v ...any)
	DebugF(str string, v ...any)
	Info(v ...any)
	InfoF(str string, v ...any)
	Warn(v ...any)
	WarnF(str string, v ...any)
	Error(v ...any)
	ErrorF(str string, v ...any)
	Fatal(v ...any)
	FatalF(str string, v ...any)

	// Json 打印json
	Json(v any)
	JsonFormat(v any)

	// DebugFX 有上下文的Info级别日志接口, format字符串格式
	DebugFX(ctx context.Context, str string, v ...any)
	InfoFX(ctx context.Context, str string, v ...any)
	WarnFX(ctx context.Context, str string, v ...any)
	ErrorFX(ctx context.Context, str string, v ...any)
	FatalFX(ctx context.Context, str string, v ...any)
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
