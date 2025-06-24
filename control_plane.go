package main

import (
	"fmt"
	"net"
)

func runControlPlane() {
	fmt.Println("Listening for messages...")
	listen()
}

func listen() {
	listenAddr, err := net.ResolveUDPAddr("udp4", LISTEN_ADDR)
	if err != nil {
		fmt.Printf("Invalid listen address: %v", err)
		return
	}

	conn, err := net.ListenUDP("udp4", listenAddr)
	if err != nil {
		fmt.Printf("Error creating listen connection on %v: %v\n", LISTEN_ADDR, err)
		return
	}

	for {
		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("Error reading data: %v\n", err)
			continue
		}

		var msg UDPMessage
		err = msg.Parse(buffer[:n])
		if err != nil {
			fmt.Printf("Error parsing message: %v\n", err)
			continue
		}
	}
}
