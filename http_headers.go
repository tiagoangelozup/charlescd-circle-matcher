package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

type httpHeaders struct {
	types.DefaultHttpContext
	contextID uint32
}

func newHttpHeaders(contextID uint32) *httpHeaders {
	return &httpHeaders{contextID: contextID}
}

func (ctx *httpHeaders) OnHttpResponseHeaders(numHeaders int, endOfStream bool) types.Action {
	if err := proxywasm.AddHttpResponseHeader("hello", "kurtis"); err != nil {
		proxywasm.LogCriticalf("failed to set response header: %v", err)
	}
	proxywasm.LogDebug("kurtis here! :)")
	return types.ActionContinue
}
