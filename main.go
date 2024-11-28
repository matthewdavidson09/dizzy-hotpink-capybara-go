package main

import (
	"dizzy-pink-capybara-go/generator"
	"flag"
	"fmt"
)

func main() {
	// Define command-line arguments
	count := flag.Int("count", 1, "Number of hostnames to generate in direct mode")
	flag.Parse()

	generateHostname(*count)
}

func generateHostname(count int) {
	// Generate and print hostnames
	hostnames := make([]string, count)
	for i := 0; i < count; i++ {
		hostnames[i] = generator.GenerateHostname()
		fmt.Println(hostnames[i])
	}
}
