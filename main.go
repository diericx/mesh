package main

import (
	"fmt"
	"os"
)

type Message struct {
	// String
	endpoint [32]byte
	// String
	convId [16]byte
}

var LISTEN_ADDR string = "0.0.0.0:8080"

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 1 {
		fmt.Printf("Please specify either client or ctrl")
		return
	}

	switch argsWithoutProg[1] {
	case "client":
		runClient()
	case "ctrl":
		runControlPlane()
	default:
		fmt.Println("Unrecognized command")

	}
}
