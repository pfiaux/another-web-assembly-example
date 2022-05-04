package main

import (
	"log"
)

func main() {
	log.Print("WASM example initializing...")
	engine := Engine{}
	RegisterCallbacks(&engine)
	log.Print("Registered Callbacks")
	// The channel makes sure that the main function does not end
	// and keeps the application running.
	<-make(chan struct{})
}
