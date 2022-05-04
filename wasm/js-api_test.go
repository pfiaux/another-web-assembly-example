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
	testCases := []struct {
		name          string
		inputs        []js.Value
		expectedColor HSLColor
	}{
		{
			name:          "Empty input is ignored",
			inputs:        []js.Value{},
			expectedColor: HSLColor{0, 0, 0},
		},
		{
			name:          "Invalid input type is ignored",
			inputs:        []js.Value{js.ValueOf(42)},
			expectedColor: HSLColor{0, 0, 0},
		},
		{
			name:          "valid input",
			inputs:        []js.Value{js.ValueOf("ArrowUp")},
			expectedColor: HSLColor{10, 0, 0},
		},
	}

	for _, tc := range testCases {
		// REVIEW: I would run each case as its own sub-test.
		t.Run(tc.name, func(t *testing.T) {
			is := is.New(t)

			engine := Engine{}

			handleKeyEventWrapper(&engine, tc.inputs)

			is.Equal(engine.color, tc.expectedColor)
		})
	}
}

func TestGetColorForJS(t *testing.T) {
	is := is.New(t)
	engine := &Engine{
		color: HSLColor{50, 90, 35},
	}

	jsFriendlyObject := getColorForJS(engine)

	is.Equal(jsFriendlyObject, map[string]interface{}{
		"hue":        50,
		"saturation": 90,
		"lightness":  35,
	})
}
