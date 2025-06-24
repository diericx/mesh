package main

import (
	"fmt"
	"net"
	"os"
)

func runClient() {
	fmt.Println("Sending a message to the control plane...")
	send()
}

func send() {
	destAddr, err := net.ResolveUDPAddr("udp4", "0.0.0.0:8080")
	if err != nil {
		fmt.Printf("Invalid listen address: %v", err)
		return
	}

	// Dial a UDP connection
	conn, err := net.DialUDP("udp", nil, destAddr)
	if err != nil {
		fmt.Println("Error dialing UDP:", err)
		os.Exit(1)
	}
	defer conn.Close() // Close the connection when the main function exits

	var msg UDPMessage = UDPMessage{}
	msg.SetEndpoint("hello")
	msg.SetConvID("world")

	// Send the message
	_, err = conn.Write(msg.Serialize())
	if err != nil {
		fmt.Println("Error sending message:", err)
		os.Exit(1)
	}
}
