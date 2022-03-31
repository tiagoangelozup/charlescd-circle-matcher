package json

import (
	"encoding/base64"
	"fmt"
	"github.com/buger/jsonparser"
)

type Object struct{ data []byte }

func FromBase64(encoded string) (*Object, error) {
	data, err := base64.RawStdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, fmt.Errorf("error decoding base64: %w", err)
	}
	return &Object{data: data}, nil
}

func (o *Object) GetValue(key string) (*Value, error) {
	value, dataType, offset, err := jsonparser.Get(o.data, SplitKey(key)...)
	if err != nil {
		return nil, fmt.Errorf("error getting value from key %q: %w", key, err)
	}
	return &Value{value: value, dataType: dataType, offset: offset}, nil
}
