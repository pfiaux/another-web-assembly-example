package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

type Engine struct {
	color HSLColor
}

type KeyEvent struct {
	key string
}

type HSLColor struct {
	Hue        int `yaml:"hue"`        // 0-360
	Saturation int `yaml:"saturation"` // 0-100
	Lightness  int `yaml:"lightness"`  // 0-100 (also sometimes called luminance)
}

func (engine *Engine) loadConfig(config []byte) error {
	err := yaml.Unmarshal(config, &engine.color)
	if err != nil {
		return fmt.Errorf("Failed to parse config: %v", err)
	}
	log.Printf("Color after loading: %v", engine.color)
	return nil
}

func (engine *Engine) handleKeyEvent(event *KeyEvent) {
	switch event.key {
	case "ArrowUp":
		if engine.color.Hue <= 350 {
			engine.color.Hue += 10
		}
		log.Printf("engine: color updated %v", engine.color)
	case "ArrowDown":
		if engine.color.Hue >= 10 {
			engine.color.Hue -= 10
		}
		log.Printf("engine: color updated %v", engine.color)
	}
}
