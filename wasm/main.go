package main

import (
	"log"
)

func main() {
	log.Print("WASM example initializing...")
	engine := Engine{}
	RegisterCallbacks(&engine)
	log.Print("Registered Callbacks")
	// REVIEW: In my opinion, this channel should not be part of the engine,
	// because the purpose is just to make sure, that the main function does not
	// end. Therefore you could create this channel right here
	<-make(chan struct{})
}
