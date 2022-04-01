package log

import (
	"fmt"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
)

func Debugf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	proxywasm.LogDebug(msg)
}

func Infof(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	proxywasm.LogInfo(msg)
}

func Warnf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	proxywasm.LogWarn(msg)
}

func Errorf(err error, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	msg = fmt.Sprintf("%s: %v", msg, err)
	proxywasm.LogError(msg)
}

func Criticalf(err error, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	msg = fmt.Sprintf("%s: %v", msg, err)
	proxywasm.LogCritical(msg)
}

func Debug(msg string) {
	proxywasm.LogDebug(msg)
}

func Info(msg string) {
	proxywasm.LogInfo(msg)
}

func Warn(msg string) {
	proxywasm.LogWarn(msg)
}

func Error(err error, msg string) {
	msg = fmt.Sprintf("%s: %v", msg, err)
	proxywasm.LogError(msg)
}

func Critical(err error, msg string) {
	msg = fmt.Sprintf("%s: %v", msg, err)
	proxywasm.LogCritical(msg)
}
