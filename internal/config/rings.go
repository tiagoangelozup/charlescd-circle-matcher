package config

import (
	"encoding/json"
	"fmt"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
)

type Rings []*Ring

func RingsFromPlugin() (Rings, error) {
	data, err := proxywasm.GetPluginConfiguration()
	if err != nil {
		return nil, fmt.Errorf("error reading rings from plugin configuration: %w", err)
	}
	return unmarshalRings(data)
}

func unmarshalRings(data []byte) (Rings, error) {
	var r Rings
	if err := json.Unmarshal(data, &r); err != nil {
		return nil, fmt.Errorf("error unmarshalling configurations: %w", err)
	}
	return r, nil
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
