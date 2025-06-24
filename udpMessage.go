package main

import (
	"bytes"
	"fmt"
	unsafe "unsafe"
)

var endpointMaxLength = 32
var convIdMaxLength = 16

type UDPMessage struct {
	// String
	endpoint [32]byte
	// String
	convId [16]byte
}

func (m *UDPMessage) SetEndpoint(v string) error {
	if len(v) > 32 {
		return fmt.Errorf("string too large")
	}
	copy(m.endpoint[:], v)
	return nil
}

func (m *UDPMessage) SetConvID(v string) error {
	if len(v) > 16 {
		return fmt.Errorf("string too large")
	}
	copy(m.convId[:], v)
	return nil
}

func (m *UDPMessage) Serialize() []byte {
	return append(m.endpoint[:], m.convId[:]...)
}

func (m *UDPMessage) Parse(b []byte) error {
	inputBytesLength := int(unsafe.Sizeof(m))
	if len(b) != inputBytesLength {
		return fmt.Errorf("invalid message size: %v", inputBytesLength)
	}

	copy(bytes.Trim(b[:32], "\x00"), m.endpoint[:])
	copy(bytes.Trim(b[32:], "\x00"), m.convId[:])
	return nil
}
