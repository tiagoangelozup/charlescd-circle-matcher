package json

import (
	"fmt"
	"github.com/buger/jsonparser"
)

type Value struct {
	value    []byte
	dataType jsonparser.ValueType
	offset   int
}

func (v *Value) AsFloat() (float64, error) {
	if v.dataType != jsonparser.Number {
		return 0, fmt.Errorf("value is not a number: %s", string(v.value))
	}
	return jsonparser.ParseFloat(v.value)
}

func (v *Value) String() string {
	return string(v.value)
}
