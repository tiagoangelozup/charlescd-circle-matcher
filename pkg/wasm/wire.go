//go:build wireinject
// +build wireinject

package wasm

import (
	"github.com/google/wire"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/config"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/context"
)

func newHttpContext(context.HttpID, context.PluginID, config.PluginRawData, config.VMRawData) types.HttpContext {
	wire.Build(ProviderSet)
	return nil
}
