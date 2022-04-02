package wasm

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/config"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/log"
	"sync"
)

type VM struct {
	types.DefaultVMContext

	sync.RWMutex
	vmConfig config.VMRawData
}

func NewVM() *VM { return &VM{} }

func (v *VM) OnVMStart(vmConfigurationSize int) types.OnVMStartStatus {
	cfg, err := proxywasm.GetVMConfiguration()
	if err != nil && err != types.ErrorStatusNotFound {
		log.Critical(err, "error reading vm configuration")
		return types.OnVMStartStatusFailed
	}
	if err != types.ErrorStatusNotFound {
		v.SetVMConfig(cfg)
	}
	log.Infof(`wasm vm started! {"vm_config_size":%d}`, vmConfigurationSize)
	return types.OnVMStartStatusOK
}

func (v *VM) NewPluginContext(contextID uint32) types.PluginContext {
	vmConfig := v.VMConfig()
	plugin := NewPlugin(contextID)
	plugin.SetVMConfig(vmConfig)
	return plugin
}

func (v *VM) VMConfig() config.VMRawData {
	v.RLock()
	defer v.RUnlock()
	return v.vmConfig
}

func (v *VM) SetVMConfig(cfg config.VMRawData) {
	v.Lock()
	defer v.Unlock()
	v.vmConfig = cfg
}
