package json

import (
	"encoding/base64"
	"fmt"
	"github.com/buger/jsonparser"
	"regexp"
	"strings"
)

var arrayIndexRegex = regexp.MustCompile(`(?m)\[[0-9]+]`)

type Object struct{ data []byte }

func FromBase64(encoded string) (*Object, error) {
	data, err := base64.RawStdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, fmt.Errorf("error decoding base64: %w", err)
	}
	return &Object{data: data}, nil
}

func (o *Object) GetString(key string) (string, error) {
	value, err := jsonparser.GetString(o.data, SplitKey(key)...)
	if err != nil {
		return "", fmt.Errorf("error getting string value from key %q: %w", key, err)
	}
	return value, nil
}

func (o *Object) GetBoolean(key string) (bool, error) {
	value, err := jsonparser.GetBoolean(o.data, SplitKey(key)...)
	if err != nil {
		return false, fmt.Errorf("error getting string value from key %q: %w", key, err)
	}
	return value, nil
}

func (o *Object) GetInt(key string) (int64, error) {
	value, err := jsonparser.GetInt(o.data, SplitKey(key)...)
	if err != nil {
		return 0, fmt.Errorf("error getting string value from key %q: %w", key, err)
	}
	return value, nil
}

func (o *Object) GetFloat(key string) (float64, error) {
	value, err := jsonparser.GetFloat(o.data, SplitKey(key)...)
	if err != nil {
		return 0, fmt.Errorf("error getting string value from key %q: %w", key, err)
	}
	return value, nil
}

func SplitKey(key string) []string {
	results := make([]string, 0)
	for _, splitted := range strings.Split(key, ".") {
		arraykeys := make([]string, 0)
		arraykeys = append(arraykeys, arrayIndexRegex.FindAllString(splitted, -1)...)
		if idx := strings.LastIndex(splitted, strings.Join(arraykeys, "")); idx >= 0 {
			splitted = splitted[:idx]
		}
		results = append(results, splitted)
		results = append(results, arraykeys...)
	}
	return results
}
