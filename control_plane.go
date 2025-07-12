package main

import (
	"fmt"
	"net"
)

func udpAPIlistener(listenAddrStr string) {
	fmt.Println("UDP listener started.")

	listenAddr, err := net.ResolveUDPAddr("udp4", listenAddrStr)
	if err != nil {
		fmt.Printf("Invalid listen address: %v", err)
		return
	}

	conn, err := net.ListenUDP("udp4", listenAddr)
	if err != nil {
		fmt.Printf("Error creating listen connection on %v: %v\n", listenAddr, err)
		return
	}

	for {
		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("Error reading data: %v\n", err)
			continue
		}

		msg, err := ParseUDPMessage(buffer[:n])
		if err != nil {
			fmt.Printf("Error parsing message: %v\n", err)
			continue
		}

		fmt.Println(msg.endpoint)
		fmt.Println(msg.reqId)
	}
}
