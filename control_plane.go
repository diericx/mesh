package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func udpAPIlistener(listenAddrStr string) {
	fmt.Printf("UDP listener started on %s\n", listenAddrStr)

	// Initialize the API router
	router := InitializeRouter()

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
		n, clientAddr, err := conn.ReadFromUDP(buffer)
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

		fmt.Printf("Received request: Endpoint=%s, ReqId=%s\n", req.Endpoint, req.ReqId)

		// Route the request and get response
		response := router.Route(req)

		// Send response back to client
		err = sendUDPResponse(conn, clientAddr, response)
		if err != nil {
			fmt.Printf("Error sending response: %v\n", err)
		}
	}
}

// sendUDPResponse sends a UDP response back to the client
func sendUDPResponse(conn *net.UDPConn, clientAddr *net.UDPAddr, response UDPResponse) error {
	responseData, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("error marshaling response: %v", err)
	}

	_, err = conn.WriteToUDP(responseData, clientAddr)
	if err != nil {
		return fmt.Errorf("error writing response: %v", err)
	}

	fmt.Printf("Sent response: Endpoint=%s, ReqId=%s, Code=%d\n", response.Endpoint, response.ReqId, response.Code)
	return nil
}
