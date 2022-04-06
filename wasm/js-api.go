package main

import (
	"log"
	"syscall/js"
)

func RegisterCallbacks(engine *Engine) {
	js.Global().Set("parseConfig", js.FuncOf(func(this js.Value, inputs []js.Value) interface{} {
		initWrapper(engine, inputs)
		return nil
	}))

	js.Global().Set("handleKeyEvent", js.FuncOf(func(this js.Value, inputs []js.Value) interface{} {
		handleKeyEventWrapper(engine, inputs)
		return nil
	}))

	js.Global().Set("getColor", js.FuncOf(func(this js.Value, inputs []js.Value) interface{} {
		log.Print("js-api: Returning the color...")
		return engine.getJSObject()
	}))
}

func initWrapper(engine *Engine, inputs []js.Value) {
	if len(inputs) != 1 {
		log.Fatalf("js-api: Expecting 1 argument, the config file: %v", len(inputs))
		return
	}
	if inputs[0].Type() != js.TypeObject {
		log.Fatalf("js-api: Argument 1 for parseConfig(config) must be a %v not %v", js.TypeObject, inputs[0].Type())
		return
	}

	jsConfigFile := inputs[0]
	config := copyBytesFromJS(jsConfigFile)
	log.Printf("js-api: got config.yaml... %s", string(config))
	err := engine.loadConfig(config)
	if err != nil {
		log.Fatalf("js-api: parseConfig(config) failed to load config %v", err)
	}
}

func handleKeyEventWrapper(engine *Engine, inputs []js.Value) {
	if len(inputs) != 1 {
		log.Printf("js-api: Expecting 1 argument, the config file: %v", len(inputs))
		return
	}
	if inputs[0].Type() != js.TypeString {
		log.Printf("js-api: Argument 1 for handleKeyEvent() must be of %v not %v", js.TypeString, inputs[0].Type())
		return
	}
	event := KeyEvent{
		key: inputs[0].String(),
	}
	log.Printf("js-api: forwarding key event... %v", event)
	engine.handleKeyEvent(&event)
}

func copyBytesFromJS(input js.Value) []byte {
	data := make([]uint8, input.Get("byteLength").Int())
	js.CopyBytesToGo(data, input)
	return data
}
