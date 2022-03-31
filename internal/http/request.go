package http

import (
	"fmt"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	"strings"
)

type Request interface {
	GetHeader(key string) (string, error)
}

type envoyRequest struct {
	numHeaders  int
	endOfStream bool
}

func NewRequest(numHeaders int, endOfStream bool) *envoyRequest {
	return &envoyRequest{numHeaders: numHeaders, endOfStream: endOfStream}
}

func (r *envoyRequest) GetHeader(key string) (string, error) {
	hs, err := proxywasm.GetHttpRequestHeaders()
	if err != nil && err != types.ErrorStatusNotFound {
		return "", fmt.Errorf("error getting http request header %q: %w", key, err)
	}
	for _, h := range hs {
		if k, v := h[0], h[1]; strings.EqualFold(k, key) {
			return v, nil
		}
	}
	return "", nil
}

func (r *envoyRequest) AddHeader(key, value string) error {
	err := proxywasm.AddHttpRequestHeader(key, value)
	if err != nil {
		return fmt.Errorf("error adding value %q to http request header %q: %w", value, key, err)
	}
	return nil
}
