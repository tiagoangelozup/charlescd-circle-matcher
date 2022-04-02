package wasm

import (
	"github.com/google/wire"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/config"
	"github.com/tiagoangelozup/charlescd-circle-matcher/pkg/ring"
	"github.com/tiagoangelozup/charlescd-circle-matcher/pkg/router"
)

var ProviderSet = wire.NewSet(
	// providers
	config.ProviderSet,
	ring.ProviderSet,
	router.ProviderSet,

	// interfaces
	wire.Bind(new(router.RingService), new(*ring.Service)),
	wire.Bind(new(types.HttpContext), new(*router.JWT)),
)
