package config

import (
	"encoding/json"
	"fmt"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

func RingsFromPlugin() ([]*Ring, error) {
	data, err := proxywasm.GetPluginConfiguration()
	if err != nil && err != types.ErrorStatusNotFound {
		return nil, fmt.Errorf("error getting plugin configurations: %w", err)
	}
	if data == nil {
		return nil, nil
	}
	plugin, err := unmarshalPlugin(data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling plugin configurations: %w", err)
	}
	return plugin.Rings, err
}
func unmarshalPlugin(data []byte) (*Plugin, error) {
	r := new(Plugin)
	if err := json.Unmarshal(data, r); err != nil {
		return nil, fmt.Errorf("error unmarshaling plugin configurations: %w", err)
	}
	return r, nil
}

type Plugin struct {
	Rings []*Ring `json:"rings,omitempty"`
}

type Ring struct {
	ID    string `json:"id,omitempty"`
	Match *Match `json:"match,omitempty"`
}

type Match struct {
	All []*Rule `json:"all,omitempty"`
	Any []*Rule `json:"any,omitempty"`
}

type Rule struct {
	Key      string   `json:"key,omitempty"`
	Operator string   `json:"operator,omitempty"`
	Values   []string `json:"values,omitempty"`
}
