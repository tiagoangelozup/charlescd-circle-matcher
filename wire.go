//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

func newHttpContext(contextID uint32) types.HttpContext {
	wire.Build(providers)
	return nil
}