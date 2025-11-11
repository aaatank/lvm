package lvmsdk

import (
	_ "embed"
)

//go:embed lvm_fe/lvm_fe.wasm
var wasm []byte
