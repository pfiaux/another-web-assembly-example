package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestBuildEngine(t *testing.T) {
	is := is.New(t)

	engine := Engine{}

	is.Equal(engine.color, HSLColor{0, 0, 0})
}

func TestLoadConfig(t *testing.T) {
	testCases := []struct {
		name          string
		config        string
		expectedColor HSLColor
		hasError      bool
	}{
		{
			name: "success",
			config: `hue: 42
saturation: 42
lightness: 42`,
			expectedColor: HSLColor{42, 42, 42},
			hasError:      false,
		},
		{
			name:     "invalid yaml",
			config:   `Not a valid yaml string`,
			hasError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			is := is.New(t)

			engine := Engine{}

			err := engine.loadConfig([]byte(tc.config))

			if tc.hasError {
				is.True(err != nil)
			} else {
				is.NoErr(err)
				is.Equal(engine.color, tc.expectedColor)
			}
		})
	}
}

func TestHandleKeyEvent(t *testing.T) {
	startColor := HSLColor{180, 50, 50}
	testCases := []struct {
		name          string
		event         KeyEvent
		expectedColor HSLColor
	}{
		{
			name:          "ArrowUp",
			event:         KeyEvent{key: "ArrowUp"},
			expectedColor: HSLColor{190, 50, 50},
		},
		{
			name:          "ArrowDown",
			event:         KeyEvent{key: "ArrowDown"},
			expectedColor: HSLColor{170, 50, 50},
		},
		{
			name:          "ArrowLeft",
			event:         KeyEvent{key: "ArrowLeft"},
			expectedColor: HSLColor{180, 50, 50},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			is := is.New(t)

			engine := Engine{
				color: startColor,
			}

			engine.handleKeyEvent(&tc.event)

			is.Equal(engine.color, tc.expectedColor)
		})
	}
}
