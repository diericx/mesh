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
	Endpoint string `json:"endpoint"`
	ReqId    string `json:"reqId"`
	Code     int    `json:"code"`
	Content  string `json:"content"`
}

// Handler function type for API endpoints
type APIHandler func(req UDPRequest) UDPResponse

// APIRouter manages endpoint routing
type APIRouter struct {
	handlers map[string]APIHandler
}

// NewAPIRouter creates a new router instance
func NewAPIRouter() *APIRouter {
	return &APIRouter{
		handlers: make(map[string]APIHandler),
	}
}

// RegisterHandler registers a handler for a specific endpoint
func (r *APIRouter) RegisterHandler(endpoint string, handler APIHandler) {
	r.handlers[endpoint] = handler
}

// Route processes a request and returns a response
func (r *APIRouter) Route(req UDPRequest) UDPResponse {
	handler, exists := r.handlers[req.Endpoint]
	if !exists {
		return UDPResponse{
			Endpoint: req.Endpoint,
			ReqId:    req.ReqId,
			Code:     404,
			Content:  "[]",
		}
	}
	return handler(req)
}
