package ring

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewService,
)
