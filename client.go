package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func sendUDPRequest(destAddrStr string, req UDPRequest) {
	destAddr, err := net.ResolveUDPAddr("udp4", destAddrStr)
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

	msgSerialized, err := json.Marshal(req)
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
