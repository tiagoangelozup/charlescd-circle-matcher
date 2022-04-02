package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tiagoangelozup/charlescd-circle-matcher/pkg/wasm"
)

func main() {
	proxywasm.SetVMContext(wasm.NewVM())
}
