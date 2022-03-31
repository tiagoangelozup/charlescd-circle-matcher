package http

import (
	"fmt"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
)

type Response interface {
	AddHeader(key, value string) error
}

type envoyResponse struct {
	numHeaders  int
	endOfStream bool
}

func NewResponse(numHeaders int, endOfStream bool) *envoyResponse {
	return &envoyResponse{numHeaders: numHeaders, endOfStream: endOfStream}
}

func (r *envoyResponse) AddHeader(key, value string) error {
	if err := proxywasm.AddHttpResponseHeader("hello", "kurtis"); err != nil {
		return fmt.Errorf("failed to set response header: %v", err)
	}
	return nil
}
