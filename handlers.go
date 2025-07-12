package main

// Handler for listing networks
func handleNetworks(req UDPRequest) UDPResponse {
	return UDPResponse{
		Endpoint: req.Endpoint,
		ReqId:    req.ReqId,
		Code:     200,
		Content:  "[]", // Empty array as requested
	}
}

// InitializeRouter sets up the router with all endpoint handlers
func InitializeRouter() *APIRouter {
	router := NewAPIRouter()

	// Register handlers for different endpoints
	router.RegisterHandler("/api/v1/networks", handleNetworks)

	return router
}
