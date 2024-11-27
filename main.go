package main

import (
	"flag"
	"fmt"
	"os"

	"animal-hash-digest/generator"
)

func main() {
	// Define a command-line argument for the number of hostnames to generate
	count := flag.Int("count", 1, "Number of hostnames to generate")
	flag.Parse()

	// Validate the argument
	if *count <= 0 {
		fmt.Println("Please specify a positive number for the count.")
		os.Exit(1)
	}

	// Generate and print hostnames
	for i := 0; i < *count; i++ {
		fmt.Println(generator.GenerateHostname())
	}
}
