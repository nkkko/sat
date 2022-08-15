//go:build !wasmer && !wasmedge
// +build !wasmer,!wasmedge

package engine

import (
	"github.com/suborbital/sat/api"
	"github.com/suborbital/sat/engine/moduleref"
	"github.com/suborbital/sat/engine/runtime"
	runtimewasmtime "github.com/suborbital/sat/engine/runtime/wasmtime"
)

func runtimeBuilder(ref *moduleref.WasmModuleRef, api api.HostAPI) runtime.RuntimeBuilder {
	return runtimewasmtime.NewBuilder(ref, api)
}