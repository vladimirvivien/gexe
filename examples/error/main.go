package main

import (
	"fmt"

	"github.com/vladimirvivien/gexe"
)

func main() {
	// Set some variables for demonstration
	gexe.SetVar("service", "database")
	gexe.SetVar("operation", "connection")

	// Example 1: Basic error creation with formatting
	filename := "config.json"
	err1 := gexe.Error("Failed to process file %s", filename)
	fmt.Printf("Error 1: %v\n", err1)

	// Example 2: Error with variable expansion
	err2 := gexe.Error("${service} ${operation} failed in ${HOME}")
	fmt.Printf("Error 2: %v\n", err2)

	// Example 3: Combined formatting and variable expansion
	port := 5432
	err3 := gexe.Error("Could not connect to ${service} on port %d in ${HOME}", port)
	fmt.Printf("Error 3: %v\n", err3)

	// Example 4: Using Session.Error method directly
	session := gexe.New()
	session.SetVar("component", "auth-service")
	err4 := session.Error("${component} authentication failed for user %s", "john_doe")
	fmt.Printf("Error 4: %v\n", err4)

	// Example 5: Error in a function that returns an error
	if err := processFile("nonexistent.txt"); err != nil {
		fmt.Printf("Function error: %v\n", err)
	}

	// Example 6: Error with multiple variables and formatting
	user := "alice"
	action := "delete"
	gexe.SetVar("resource", "user-data")
	err6 := gexe.Error("User %s is not authorized to perform %s operation on ${resource} in ${HOME}", user, action)
	fmt.Printf("Error 6: %v\n", err6)
}

// processFile demonstrates using gexe.Error in a function
func processFile(filename string) error {
	// Simulate checking if file exists
	if !gexe.PathExists(filename) {
		return gexe.Error("File %s does not exist in ${HOME}", filename)
	}
	return nil
}
