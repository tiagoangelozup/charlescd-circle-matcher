package http

import (
	"fmt"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	"strings"
)

type Request struct {
	numHeaders  int
	endOfStream bool
}

func NewRequest(numHeaders int, endOfStream bool) *Request {
	return &Request{numHeaders: numHeaders, endOfStream: endOfStream}
}

func (r *Request) GetHeader(key string) (string, error) {
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
