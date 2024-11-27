package main

import (
	"animal-hash-digest/generator"
	"animal-hash-digest/server"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define command-line arguments
	port := flag.String("port", "8080", "Port number for the HTTP server")
	endpoint := flag.String("endpoint", "/hostname", "API endpoint for generating hostnames")
	mode := flag.String("mode", "server", "Mode of operation: 'server' to run the HTTP server, 'direct' to run the binary directly")
	count := flag.Int("count", 1, "Number of hostnames to generate in direct mode")
	flag.Parse()

	// Handle mode argument
	switch *mode {
	case "server":
		// Start the HTTP server
		fmt.Printf("Starting server on port %s with endpoint %s\n", *port, *endpoint)
		if err := server.Start(*port, *endpoint); err != nil {
			fmt.Printf("Error starting server: %v\n", err)
			os.Exit(1)
		}
	case "direct":
		// Run the binary's direct functionality
		runDirectMode(*count)
	default:
		// Invalid mode
		fmt.Printf("Invalid mode: %s. Use 'server' or 'direct'.\n", *mode)
		os.Exit(1)
	}
}


func runDirectMode(count int) {
	// Generate and print hostnames
	hostnames := make([]string, count)
	for i := 0; i < count; i++ {
		hostnames[i] = generator.GenerateHostname()
		fmt.Println(hostnames[i])
	}
}
