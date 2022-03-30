package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
)

func main() {
	proxywasm.SetVMContext(&vm{})
}
