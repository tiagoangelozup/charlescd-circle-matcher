package main

import (
	"github.com/google/wire"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

var providers = wire.NewSet(
	newHttpHeaders,
	wire.Bind(new(types.HttpContext), new(*httpHeaders)),
)
