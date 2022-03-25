package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

type rootContext struct {
	types.DefaultVMContext
}

type pluginContext struct {
	types.DefaultPluginContext
}

func (r *rootContext) NewPluginContext(contextID uint32) types.PluginContext {
	return &pluginContext{}
}

func (*pluginContext) NewHttpContext(contextID uint32) types.HttpContext {
	return newHttpContext(contextID)
}
