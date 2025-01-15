package marsLog

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// MarsDefaultLog 默认提供的日志对象
type MarsDefaultLog struct{}

func (log *MarsDefaultLog) Debug(v ...interface{}) {
	if os.Getenv("DEBUG_ENABLE") == "true" {
		setPrefix(DEBUG)
		getLogger().Println(v...)
	}
}

func (log *MarsDefaultLog) DebugF(str string, v ...interface{}) {
	if os.Getenv("DEBUG_ENABLE") == "true" {
		setPrefix(DEBUG)
		getLogger().Printf(str, v...)
	}
}

func (log *MarsDefaultLog) Info(v ...interface{}) {
	setPrefix(INFO)
	getLogger().Println(v...)
}

func (log *MarsDefaultLog) InfoF(str string, v ...interface{}) {
	setPrefix(INFO)
	getLogger().Printf(str, v...)
}

func (log *MarsDefaultLog) Warn(v ...interface{}) {
	setPrefix(WARN)
	getLogger().Println(v...)
}

func (log *MarsDefaultLog) WarnF(str string, v ...interface{}) {
	setPrefix(WARN)
	getLogger().Printf(str, v...)
}

func (log *MarsDefaultLog) Error(v ...interface{}) {
	setPrefix(ERROR)
	getErrorLogger().Println(v...)
}

func (log *MarsDefaultLog) ErrorF(str string, v ...interface{}) {
	setPrefix(ERROR)
	getErrorLogger().Printf(str, v...)
}

func (log *MarsDefaultLog) Fatal(v ...interface{}) {
	setPrefix(FATAL)
	getErrorLogger().Println(v...)
}

func (log *MarsDefaultLog) FatalF(str string, v ...interface{}) {
	setPrefix(FATAL)
	getErrorLogger().Printf(str, v...)
}

func (log *MarsDefaultLog) Json(v interface{}) {
	setPrefix(JSON)
	getLogger().Println(jsonMarshal(v))
}

func (log *MarsDefaultLog) JsonFormat(v interface{}) {
	setPrefix(JSON)
	getLogger().Println("\r\n" + prettyString(jsonMarshal(v)))
}

func (log *MarsDefaultLog) DebugFX(ctx context.Context, str string, v ...interface{}) {
	if os.Getenv("DEBUG_ENABLE") == "true" {
		setPrefix(DEBUG)
		getLogger().Println(ctx)
		getLogger().Printf(str, v...)
	}
}

func (log *MarsDefaultLog) InfoFX(ctx context.Context, str string, v ...interface{}) {
	setPrefix(INFO)
	getLogger().Println(ctx)
	getLogger().Printf(str, v...)
}

func (log *MarsDefaultLog) WarnFX(ctx context.Context, str string, v ...interface{}) {
	setPrefix(WARN)
	getLogger().Println(ctx)
	getLogger().Printf(str, v...)
}

func (log *MarsDefaultLog) ErrorFX(ctx context.Context, str string, v ...interface{}) {
	setPrefix(ERROR)
	getErrorLogger().Println(ctx)
	getErrorLogger().Printf(str, v...)
}

func (log *MarsDefaultLog) FatalFX(ctx context.Context, str string, v ...interface{}) {
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

func jsonMarshal(v any) string {
	marshal, err := json.Marshal(v)
	if err != nil {
		return "{}"
	}
	return string(marshal)
}

func prettyString(str string) string {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", " "); err != nil {
		return ""
	}
	return prettyJSON.String()
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
