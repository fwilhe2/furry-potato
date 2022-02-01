package main

import (
	"fmt"
	"io/ioutil"
	wasmer "github.com/wasmerio/wasmer-go/wasmer"
)

func main() {
    wasmBytes, _ := ioutil.ReadFile("foo.wat")

    engine := wasmer.NewEngine()
    store := wasmer.NewStore(engine)

    // Compiles the module
    module, _ := wasmer.NewModule(store, wasmBytes)

    // Instantiates the module
    importObject := wasmer.NewImportObject()
    instance, _ := wasmer.NewInstance(module, importObject)

    // Gets the `sum` exported function from the WebAssembly instance.
    sum, _ := instance.Exports.GetFunction("add")

    // Calls that exported function with Go standard values. The WebAssembly
    // types are inferred and values are casted automatically.
    result, _ := sum(5, 37)

    fmt.Println(result) // 42!
}