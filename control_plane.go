package main

import (
	"fmt"
	"net"

	unsafe "unsafe"
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

		var msg Message

		receivedMessage := buffer[:n]
		if len(receivedMessage) != int(unsafe.Sizeof(msg)) {
			fmt.Printf("Invalid message size: %v\n", err)
			continue
		}

		fmt.Println(string(receivedMessage[:32]))
		fmt.Println(string(receivedMessage[32:]))
	}
}
