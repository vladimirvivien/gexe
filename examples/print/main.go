package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/vladimirvivien/gexe"
)

func main() {
	fmt.Println("=== Gexe Print Functionality Demo ===")

	// 1. Basic Print functionality
	fmt.Println("\n1. Basic Print functionality:")
	gexe.Print("Hello ")
	gexe.Println("World!")

	// 2. Print with sprintf formatting
	fmt.Println("\n2. Print with sprintf formatting:")
	gexe.Print("User: %s, ", "Alice")
	gexe.Println("ID: %d", 42)

	// 3. Print with variable expansion
	fmt.Println("\n3. Print with variable expansion:")
	gexe.SetVar("name", "Bob")
	gexe.SetVar("role", "Admin")
	gexe.Println("User ${name} has role: ${role}")

	// 4. Combined sprintf and variable expansion
	fmt.Println("\n4. Combined sprintf and variable expansion:")
	gexe.SetVar("env", "production")
	gexe.Println("Deploying app %s to ${env} environment", "myapp")

	// 5. Method chaining
	fmt.Println("\n5. Method chaining:")
	g := gexe.New()
	g.SetVar("service", "web-server").
		Print("Starting ").
		Print("${service} ").
		Println("on port %d", 8080)

	// 6. PrintTo functionality
	fmt.Println("\n6. PrintTo functionality:")
	var buf bytes.Buffer
	gexe.SetVar("logLevel", "INFO")
	gexe.PrintTo(&buf, "[${logLevel}] Application started at %s", "2023-12-01 10:00:00")
	fmt.Printf("Buffer contents: %s\n", buf.String())

	// 7. PrintTo with file
	fmt.Println("\n7. PrintTo with file:")
	tempFile, err := os.CreateTemp("", "gexe_print_demo_*.log")
	if err != nil {
		fmt.Printf("Failed to create temp file: %v\n", err)
		return
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	gexe.SetVar("user", "admin")
	gexe.PrintTo(tempFile, "User ${user} performed action: %s\n", "login")

	// Read back the file content
	content := gexe.FileRead(tempFile.Name()).String()
	fmt.Printf("File contents: %s", content)

	// 8. Real-world example: log-like output
	fmt.Println("\n8. Real-world example:")
	session := gexe.New()
	session.SetVar("timestamp", "2023-12-01 10:05:00").
		SetVar("level", "INFO").
		SetVar("component", "auth-service")

	session.Print("[${timestamp}] ").
		Print("[${level}] ").
		Print("[${component}] ").
		Println("User %s authentication %s", "alice@example.com", "successful")

	fmt.Println("\n=== Demo Complete ===")
}
