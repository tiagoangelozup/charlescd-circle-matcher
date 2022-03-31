//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/config"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/http"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/logger"
	"github.com/tiagoangelozup/charlescd-circle-matcher/pkg/ring"
	"github.com/tiagoangelozup/charlescd-circle-matcher/pkg/router"
)

var providers = wire.NewSet(
	ring.NewService,
	http.NewRequest,
	http.NewResponse,
	logger.NewFactory,
	newPlugin,
	router.NewJWT,
	wire.Bind(new(router.RingService), new(*ring.Service)),
	wire.Bind(new(types.HttpContext), new(*router.JWT)),
	wire.Bind(new(types.PluginContext), new(*plugin)),
)

func newPluginContext(contextID uint32) types.PluginContext {
	wire.Build(providers)
	return nil
}

func newHttpContext(contextID uint32, rings []*config.Ring) types.HttpContext {
	wire.Build(providers)
	return nil
}
