package wasm

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/config"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/context"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/log"
	"sync"
)

type Plugin struct {
	types.DefaultPluginContext

	contextID context.PluginID

	sync.RWMutex
	pluginConfig config.PluginRawData
	vmConfig     config.VMRawData
}

func NewPlugin(contextID uint32) *Plugin {
	return &Plugin{contextID: context.PluginID(contextID)}
}

func (p *Plugin) OnPluginStart(pluginConfigurationSize int) types.OnPluginStartStatus {
	cfg, err := proxywasm.GetPluginConfiguration()
	if err != nil && err != types.ErrorStatusNotFound {
		log.Critical(err, "error reading plugin configuration")
		return types.OnPluginStartStatusFailed
	}
	if err != types.ErrorStatusNotFound {
		p.SetPluginConfig(cfg)
	}
	log.Infof(`plugin started! {"plugin_config_size":%d}`, pluginConfigurationSize)
	return types.OnPluginStartStatusOK
}

func (p *Plugin) NewHttpContext(contextID uint32) types.HttpContext {
	return newHttpContext(
		context.HttpID(contextID),
		p.ContextID(),
		p.PluginConfig(),
		p.VMConfig(),
	)
}

func (p *Plugin) ContextID() context.PluginID {
	return p.contextID
}

func (p *Plugin) VMConfig() config.VMRawData {
	p.RLock()
	defer p.RUnlock()
	return p.vmConfig
}

func (p *Plugin) SetVMConfig(config config.VMRawData) {
	p.Lock()
	defer p.Unlock()
	p.vmConfig = config
}

func (p *Plugin) PluginConfig() config.PluginRawData {
	p.RLock()
	defer p.RUnlock()
	return p.pluginConfig
}

func (p *Plugin) SetPluginConfig(cfg config.PluginRawData) {
	p.Lock()
	defer p.Unlock()
	p.pluginConfig = cfg
}
