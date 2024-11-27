package main

import (
	"animal-hash-digest/server"
	"flag"
	"fmt"
)

func main() {
	// Define command-line arguments
	port := flag.String("port", "8080", "Port number for the HTTP server")
	endpoint := flag.String("endpoint", "/hostname", "API endpoint for generating hostnames")
	flag.Parse()

	// Print configuration for debugging
	fmt.Printf("Starting server on port %s with endpoint %s\n", *port, *endpoint)

	// Start the server with the specified port and endpoint
	if err := server.Start(*port, *endpoint); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
