package config

import (
	"encoding/json"
	"fmt"
)

type Rings []*Ring

func UnmarshalRings(data []byte) (Rings, error) {
	var r Rings
	if err := json.Unmarshal(data, &r); err != nil {
		return nil, fmt.Errorf("error inmarshalling configurations: %w", err)
	}
	return r, nil
}

func (r *Rings) Marshal() ([]byte, error) {
	return json.Marshal(r)
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
