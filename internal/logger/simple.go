package logger

import (
	"fmt"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
)

type simple struct{}

func (s *simple) Debugf(format string, args ...interface{}) {
	proxywasm.LogDebugf(format, args...)
}

func (s *simple) Infof(format string, args ...interface{}) {
	proxywasm.LogInfof(format, args...)
}

func (s *simple) Warnf(format string, args ...interface{}) {
	proxywasm.LogWarnf(format, args...)
}

func (s *simple) Errorf(err error, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	msg = fmt.Sprintf("%s: %v", msg, err)
	proxywasm.LogError(msg)
}

func (s *simple) Criticalf(err error, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	msg = fmt.Sprintf("%s: %v", msg, err)
	proxywasm.LogCritical(msg)
}

func (s *simple) Debug(msg string) {
	proxywasm.LogDebug(msg)
}

func (s *simple) Info(msg string) {
	proxywasm.LogInfo(msg)
}

func (s *simple) Warn(msg string) {
	proxywasm.LogWarn(msg)
}

func (s *simple) Error(err error, msg string) {
	msg = fmt.Sprintf("%s: %v", msg, err)
	proxywasm.LogError(msg)
}

func (s *simple) Critical(err error, msg string) {
	msg = fmt.Sprintf("%s: %v", msg, err)
	proxywasm.LogCritical(msg)
}
