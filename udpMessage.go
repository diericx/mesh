package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Sizes are in bytes
const endpointSize = 32
const reqIdSize = 16

type UDPRequest struct {
	endpoint    string
	reqId       string
	contentSize uint16
	content     string
}

type UDPResponse struct {
	endpoint    string
	reqId       string
	code        int
	contentSize int
	content     string
}

func (m *UDPRequest) Serialize() ([]byte, error) {
	var result [endpointSize + reqIdSize + 2]byte
	if len(m.endpoint) > endpointSize {
		return result[:], fmt.Errorf("endpoint is too large, max size is %v", endpointSize)
	}
	if len(m.reqId) > reqIdSize {
		return result[:], fmt.Errorf("convId is too large, max size is %v", reqIdSize)
	}

	copy(result[0:], m.endpoint[:])
	copy(result[0+32:], m.reqId[:])
	binary.BigEndian.PutUint16(result[0+32+16:], m.contentSize)

	resultWithContent := append(result[:], []byte(m.content)...)

	return resultWithContent, nil
}

func ParseUDPMessage(b []byte) (UDPRequest, error) {
	var parsedMessage UDPRequest

	parsedMessage.endpoint = string(bytes.Trim(b[:endpointSize], "\x00"))
	parsedMessage.reqId = string(bytes.Trim(b[endpointSize:endpointSize+reqIdSize], "\x00"))
	parsedMessage.contentSize = binary.BigEndian.Uint16(b[endpointSize+reqIdSize : endpointSize+reqIdSize+2])
	fmt.Println(len(b))
	parsedMessage.content = string(b[endpointSize+reqIdSize+2 : parsedMessage.contentSize-1])

	return parsedMessage, nil
}
