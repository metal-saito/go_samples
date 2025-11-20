package main

import (
	"fmt"
	"os"
)

// main is the entry point for the Go samples repository
func main() {
	fmt.Println("Go Samples - Reservation Tool")
	fmt.Println("This repository contains Go language samples.")
	fmt.Println("Please see samples/ directory for individual examples.")
	
	if len(os.Args) > 1 {
		fmt.Printf("\nUsage: See README.md for instructions\n")
		os.Exit(0)
	}
	
	// Demonstrate basic Go functionality
	fmt.Println("\nRepository structure:")
	fmt.Println("  samples/01_cli_reservation - CLI tool with cobra")
	fmt.Println("  samples/02_http_api - HTTP API with Gin")
	fmt.Println("  samples/03_worker_pool - Worker pool pattern")
}

