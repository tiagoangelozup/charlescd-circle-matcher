package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/config"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/logger"
	"sync"
)

type PluginLogger logger.Interface

type vm struct {
	types.DefaultVMContext
}

func (v *vm) NewPluginContext(contextID uint32) types.PluginContext {
	return newPluginContext(contextID)
}

type plugin struct {
	sync.RWMutex
	types.DefaultPluginContext
	log   PluginLogger
	rings []*config.Ring
}

func newPlugin(log PluginLogger) *plugin {
	return &plugin{log: log}
}

func (p *plugin) Rings() []*config.Ring {
	p.RLock()
	defer p.RUnlock()
	return p.rings
}

func (p *plugin) AddRings(rings []*config.Ring) {
	p.Lock()
	defer p.Unlock()
	p.rings = append(p.rings, rings...)
}

func (p *plugin) NewHttpContext(contextID uint32) types.HttpContext {
	return newHttpContext(contextID, p.Rings())
}

func (p *plugin) OnPluginStart(pluginConfigurationSize int) types.OnPluginStartStatus {
	rings, err := config.RingsFromPlugin()
	if err != nil {
		p.log.Critical(err, "error reading plugin configuration")
		return types.OnPluginStartStatusFailed
	}
	p.AddRings(rings)
	return types.OnPluginStartStatusOK
}
