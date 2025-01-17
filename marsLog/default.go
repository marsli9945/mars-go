package marsLog

import (
	"context"
	"fmt"
	"github.com/marsli9945/mars-go/marsJson"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// DefaultLog 默认提供的日志对象
type DefaultLog struct{}

func (log *DefaultLog) Debug(v ...any) {
	if os.Getenv("DEBUG_ENABLE") == "true" {
		setPrefix(DEBUG)
		getLogger().Println(v...)
	}
}

func (log *DefaultLog) DebugF(str string, v ...any) {
	if os.Getenv("DEBUG_ENABLE") == "true" {
		setPrefix(DEBUG)
		getLogger().Printf(str, v...)
	}
}

func (log *DefaultLog) Info(v ...any) {
	setPrefix(INFO)
	getLogger().Println(v...)
}

func (log *DefaultLog) InfoF(str string, v ...any) {
	setPrefix(INFO)
	getLogger().Printf(str, v...)
}

func (log *DefaultLog) Warn(v ...any) {
	setPrefix(WARN)
	getLogger().Println(v...)
}

func (log *DefaultLog) WarnF(str string, v ...any) {
	setPrefix(WARN)
	getLogger().Printf(str, v...)
}

func (log *DefaultLog) Error(v ...any) {
	setPrefix(ERROR)
	getErrorLogger().Println(v...)
}

func (log *DefaultLog) ErrorF(str string, v ...any) {
	setPrefix(ERROR)
	getErrorLogger().Printf(str, v...)
}

func (log *DefaultLog) Fatal(v ...any) {
	setPrefix(FATAL)
	getErrorLogger().Println(v...)
}

func (log *DefaultLog) FatalF(str string, v ...any) {
	setPrefix(FATAL)
	getErrorLogger().Printf(str, v...)
}

func (log *DefaultLog) Json(v any) {
	setPrefix(JSON)
	getLogger().Println(marsJson.Marshal(v))
}

func (log *DefaultLog) JsonFormat(v any) {
	setPrefix(JSON)
	getLogger().Println("\r\n" + marsJson.PrettyString(marsJson.Marshal(v)))
}

func (log *DefaultLog) DebugFX(ctx context.Context, str string, v ...any) {
	if os.Getenv("DEBUG_ENABLE") == "true" {
		setPrefix(DEBUG)
		getLogger().Println(ctx)
		getLogger().Printf(str, v...)
	}
}

func (log *DefaultLog) InfoFX(ctx context.Context, str string, v ...any) {
	setPrefix(INFO)
	getLogger().Println(ctx)
	getLogger().Printf(str, v...)
}

func (log *DefaultLog) WarnFX(ctx context.Context, str string, v ...any) {
	setPrefix(WARN)
	getLogger().Println(ctx)
	getLogger().Printf(str, v...)
}

func (log *DefaultLog) ErrorFX(ctx context.Context, str string, v ...any) {
	setPrefix(ERROR)
	getErrorLogger().Println(ctx)
	getErrorLogger().Printf(str, v...)
}

func (log *DefaultLog) FatalFX(ctx context.Context, str string, v ...any) {
	setPrefix(FATAL)
	getErrorLogger().Println(ctx)
	getErrorLogger().Printf(str, v...)
}

type Level int

const (
	green   = "\033[1;32m"
	white   = "\033[1;37m"
	yellow  = "\033[1;33m"
	red     = "\033[1;31m"
	blue    = "\033[1;34m"
	magenta = "\033[1;35m"
	cyan    = "\033[1;36m"
	reset   = "\033[0m"
)

var (
	DefaultCallerDepth = 2

	logger      *log.Logger
	errLogger   *log.Logger
	logPrefix   = ""
	levelFlags  = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL", "JSON"}
	levelColors = []string{white, green, yellow, red, magenta, blue}
)

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
	JSON
)

// setup initialize the marsLog instance
func setup() {
	logger = log.New(io.MultiWriter(os.Stdout), "", 0)
	errLogger = log.New(io.MultiWriter(os.Stdout), "", 0)
}

// getLogger init and get logger
func getLogger() *log.Logger {
	if logger == nil {
		setup()
	}
	return logger
}

// getErrorLogger init and get errLogger
func getErrorLogger() *log.Logger {
	if errLogger == nil {
		setup()
	}
	return errLogger
}

func getTime() string {
	return time.Now().Format(time.DateTime)
}

// setPrefix set the prefix of the marsLog output
func setPrefix(level Level) {
	_, runtimeFile, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%sMARS%s] %s |%s %-5s %s| %s:%-3d ", cyan, reset, getTime(), levelColors[level], levelFlags[level], reset, filepath.Base(runtimeFile), line)
	} else {
		logPrefix = fmt.Sprintf("[%sMARS%s] %s |%s %-5s %s| ", cyan, reset, getTime(), levelColors[level], levelFlags[level], reset)
	}
	getLogger().SetPrefix(logPrefix)
	getErrorLogger().SetPrefix(logPrefix)
}
