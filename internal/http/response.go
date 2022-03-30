package http

import (
	"fmt"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
)

type Response struct {
	numHeaders  int
	endOfStream bool
}

func NewResponse(numHeaders int, endOfStream bool) *Response {
	return &Response{numHeaders: numHeaders, endOfStream: endOfStream}
}

func (r *Response) AddHeader(key, value string) error {
	if err := proxywasm.AddHttpResponseHeader("hello", "kurtis"); err != nil {
		return fmt.Errorf("failed to set response header: %v", err)
	}
	return nil
}
