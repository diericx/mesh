package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func send() {
	destAddr, err := net.ResolveUDPAddr("udp4", "0.0.0.0:8081")
	if err != nil {
		fmt.Printf("Invalid dest address: %v", err)
		return
	}

	// Dial a UDP connection
	conn, err := net.DialUDP("udp", nil, destAddr)
	if err != nil {
		fmt.Println("Error dialing UDP:", err)
		os.Exit(1)
	}
	defer conn.Close() // Close the connection when the main function exits

	var msg UDPRequest = UDPRequest{
		Endpoint: "hello",
		ReqId:    "world",
	}

	msgSerialized, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Error serializing: ", err)
		os.Exit(1)
	}
	fmt.Println(msgSerialized)

	// Send the message
	_, err = conn.Write(msgSerialized)
	if err != nil {
		fmt.Println("Error sending message:", err)
		os.Exit(1)
	}
}
