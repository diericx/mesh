package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func udpAPIlistener(listenAddrStr string) {
	fmt.Printf("UDP listener started on %s\n", listenAddrStr)

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

		var req UDPRequest
		err = json.Unmarshal(buffer[:n], &req)
		if err != nil {
			fmt.Printf("Error parsing message: %v\n", err)
			continue
		}
		fmt.Println(buffer[:n])

		fmt.Println(req.Endpoint)
		fmt.Println(req.ReqId)
	}
}
