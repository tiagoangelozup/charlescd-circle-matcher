package config

import (
	"fmt"
	"github.com/buger/jsonparser"
)

type (
	converter func([]byte, jsonparser.ValueType, int, error)
	Rings     []*Ring
	Ring      struct {
		ID    string
		Match Match
	}
	Match struct {
		All []*Rule
		Any []*Rule
	}
	Rule struct {
		Key      string
		Operator string
		Values   []interface{}
	}
)

type builder struct {
	rings Rings
}

func NewRings(raw PluginRawData) Rings {
	b := &builder{rings: make(Rings, 0)}
	jsonparser.ArrayEach(raw, funcName(b), "rings")
	return b.rings
}

func funcName(b *builder) converter {
	return func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		ring := new(Ring)
		ring.ID, _ = jsonparser.GetString(value, "id")
		jsonparser.ArrayEach(value, funcName2(ring), "match", "any")
		b.rings = append(b.rings, ring)
	}
}

func funcName2(r *Ring) converter {
	return func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		rule := new(Rule)
		rule.Key, _ = jsonparser.GetString(value, "key")
		rule.Operator, _ = jsonparser.GetString(value, "operator")
		jsonparser.ArrayEach(value, funcName3(rule), "values")
		r.Match.Any = append(r.Match.Any, rule)
	}
}

func funcName3(rule *Rule) converter {
	return func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		fmt.Println(value)
	}
}
