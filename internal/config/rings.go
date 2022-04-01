package config

import (
	"encoding/json"
	"fmt"
)

func GetRings(data []byte) ([]*Ring, error) {
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
	Key      string        `json:"key,omitempty"`
	Operator string        `json:"operator,omitempty"`
	Values   []interface{} `json:"values,omitempty"`
}
