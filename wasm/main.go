package main

import (
	"log"
)

func main() {
	log.Print("WASM example initializing...")
	engine := buildEngine()
	RegisterCallbacks(&engine)
	log.Print("Registered Callbacks")
	<-engine.shutdownChannel
}
