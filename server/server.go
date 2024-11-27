package server

import (
	"fmt"
	"net/http"
)

// Start initializes and starts the HTTP server
func Start(port string, endpoint string) error {
	// Define the API route dynamically
	http.HandleFunc(endpoint, hostnameHandler)

	// Start the HTTP server
	address := fmt.Sprintf(":%s", port)
	return http.ListenAndServe(address, nil)
}
