package logger

import (
	"fmt"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
)

type contextualized struct {
	contextID   uint32
	contextName string
}

func (c *contextualized) Debugf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	msg = fmt.Sprintf("[%d] %s: %s", c.contextID, c.contextName, msg)
	proxywasm.LogDebug(msg)
}

func (c *contextualized) Infof(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	msg = fmt.Sprintf("[%d] %s: %s", c.contextID, c.contextName, msg)
	proxywasm.LogInfo(msg)
}

func (c *contextualized) Warnf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	msg = fmt.Sprintf("[%d] %s: %s", c.contextID, c.contextName, msg)
	proxywasm.LogWarn(msg)
}

func (c *contextualized) Errorf(err error, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	msg = fmt.Sprintf("[%d] %s: %s", c.contextID, c.contextName, msg)
	msg = fmt.Sprintf("%s: %v", msg, err)
	proxywasm.LogError(msg)
}

func (c *contextualized) Criticalf(err error, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	msg = fmt.Sprintf("[%d] %s: %s", c.contextID, c.contextName, msg)
	msg = fmt.Sprintf("%s: %v", msg, err)
	proxywasm.LogCritical(msg)
}

func (c *contextualized) Debug(msg string) {
	msg = fmt.Sprintf("[%d] %s: %s", c.contextID, c.contextName, msg)
	proxywasm.LogDebug(msg)
}

func (c *contextualized) Info(msg string) {
	msg = fmt.Sprintf("[%d] %s: %s", c.contextID, c.contextName, msg)
	proxywasm.LogInfo(msg)
}

func (c *contextualized) Warn(msg string) {
	msg = fmt.Sprintf("[%d] %s: %s", c.contextID, c.contextName, msg)
	proxywasm.LogWarn(msg)
}

func (c *contextualized) Error(err error, msg string) {
	msg = fmt.Sprintf("[%d] %s: %s", c.contextID, c.contextName, msg)
	msg = fmt.Sprintf("%s: %v", msg, err)
	proxywasm.LogError(msg)
}

func (c *contextualized) Critical(err error, msg string) {
	msg = fmt.Sprintf("[%d] %s: %s", c.contextID, c.contextName, msg)
	msg = fmt.Sprintf("%s: %v", msg, err)
	proxywasm.LogCritical(msg)
}
