package main

// Sizes are in bytes
const endpointSize = 32
const reqIdSize = 16

type UDPRequest struct {
	Endpoint string `json:"endpoint"`
	ReqId    string `json:"reqId"`
	Content  string `json:"content"`
}

type UDPResponse struct {
	endpoint    string
	reqId       string
	code        int
	contentSize int
	content     string
}
