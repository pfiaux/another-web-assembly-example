package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestBuildEngine(t *testing.T) {
	is := is.New(t)

	engine := buildEngine()

	is.Equal(engine.color, HSLColor{0, 0, 0})
}

func TestLoadConfig(t *testing.T) {
	is := is.New(t)
	testCases := []struct {
		config        string
		expectedColor HSLColor
		hasError      bool
	}{
		{
			config: `hue: 42
saturation: 42
lightness: 42`,
			expectedColor: HSLColor{42, 42, 42},
			hasError:      false,
		},
		{
			config:   `Not a valid yaml string`,
			hasError: true,
		},
	}

	for _, tc := range testCases {
		engine := buildEngine()

		err := engine.loadConfig([]byte(tc.config))

		if tc.hasError {
			is.True(err != nil)
		} else {
			is.NoErr(err)
			is.Equal(engine.color, tc.expectedColor)
		}
	}
}

func TestGetJSObject(t *testing.T) {
	is := is.New(t)
	engine := buildEngine()
	engine.color = HSLColor{50, 90, 35}

	jsFriendlyObject := engine.getJSObject()

	is.Equal(jsFriendlyObject, map[string]interface{}{
		"hue":        50,
		"saturation": 90,
		"lightness":  35,
	})
}

func TestHandleKeyEvent(t *testing.T) {
	is := is.New(t)
	startColor := HSLColor{180, 50, 50}
	testCases := []struct {
		event         KeyEvent
		expectedColor HSLColor
	}{
		{
			event:         KeyEvent{key: "ArrowUp"},
			expectedColor: HSLColor{190, 50, 50},
		},
		{
			event:         KeyEvent{key: "ArrowDown"},
			expectedColor: HSLColor{170, 50, 50},
		},
		{
			event:         KeyEvent{key: "ArrowLeft"},
			expectedColor: HSLColor{180, 50, 50},
		},
	}

	for _, tc := range testCases {
		engine := buildEngine()
		engine.color = startColor

		engine.handleKeyEvent(&tc.event)

		is.Equal(engine.color, tc.expectedColor)
	}
}
