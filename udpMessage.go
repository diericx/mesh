package main

import (
	"bytes"
	"fmt"
)

const endpointMaxLength = 32
const convIdMaxLength = 16

type UDPMessage struct {
	endpoint string
	convId   string
}

type serializedUDPMessage = [endpointMaxLength + convIdMaxLength]byte

func (m *UDPMessage) Serialize() (serializedUDPMessage, error) {
	var result serializedUDPMessage
	if len(m.endpoint) > endpointMaxLength {
		return result, fmt.Errorf("endpoint is too large, max size is %v", endpointMaxLength)
	}
	if len(m.convId) > convIdMaxLength {
		return result, fmt.Errorf("convId is too large, max size is %v", convIdMaxLength)
	}

	copy(result[0:], m.endpoint)
	copy(result[endpointMaxLength:], m.convId)

	return result, nil
}

func (m *UDPMessage) Parse(b []byte) error {
	var _m serializedUDPMessage
	if len(b) != len(_m) {
		return fmt.Errorf("invalid message size: %v", len(_m))
	}

	m.endpoint = string(bytes.Trim(b[:endpointMaxLength], "\x00"))
	m.convId = string(bytes.Trim(b[endpointMaxLength:endpointMaxLength+convIdMaxLength], "\x00"))

	return nil
}
