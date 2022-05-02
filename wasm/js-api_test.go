package main

import (
	"syscall/js"
	"testing"

	"github.com/matryer/is"
)

var testColorYAML = []byte(`hue: 2
saturation: 5
lightness: 6`)

func TestInitWrapper(t *testing.T) {
	is := is.New(t)
	inputs := []js.Value{
		js.Global().Get("Uint8Array").New(len(testColorYAML)),
	}
	js.CopyBytesToJS(inputs[0], testColorYAML)
	engine := Engine{}

	initWrapper(&engine, inputs)

	is.Equal(engine.color, HSLColor{2, 5, 6})
}

func TestHandleKeyEventWrapper(t *testing.T) {
	is := is.New(t)
	testCases := []struct {
		inputs        []js.Value
		expectedColor HSLColor
	}{
		{
			inputs:        []js.Value{}, // Empty input is ignored
			expectedColor: HSLColor{0, 0, 0},
		},
		{
			inputs:        []js.Value{js.ValueOf(42)}, // Invalid input type is ignored
			expectedColor: HSLColor{0, 0, 0},
		},
		{
			inputs:        []js.Value{js.ValueOf("ArrowUp")},
			expectedColor: HSLColor{10, 0, 0},
		},
	}

	for _, tc := range testCases {
		engine := Engine{}

		handleKeyEventWrapper(&engine, tc.inputs)

		is.Equal(engine.color, tc.expectedColor)
	}
}
