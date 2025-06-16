package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/vladimirvivien/gexe"
)

func main() {
	fmt.Println("=== Gexe Sprintf Functionality Demo ===")

	// 1. Basic string formatting
	fmt.Println("\n1. Basic string formatting:")
	result := gexe.Run("echo Hello %s!", "World")
	fmt.Printf("   Result: %s\n", result)

	// 2. Multiple arguments
	fmt.Println("\n2. Multiple arguments:")
	result = gexe.Run("echo User: %s, ID: %d, Active: %t", "Alice", 42, true)
	fmt.Printf("   Result: %s\n", result)

	// 3. Combined with variable expansion
	fmt.Println("\n3. Combined with variable expansion:")
	gexe.SetVar("greeting", "Welcome")
	result = gexe.Run("echo ${greeting} %s to our platform!", "Bob")
	fmt.Printf("   Result: %s\n", result)

	// 4. File operations with formatting
	fmt.Println("\n4. File operations with formatting:")
	tempDir := os.TempDir()
	filename := fmt.Sprintf("demo_%s.txt", time.Now().Format("20060102_150405"))
	fullPath := filepath.Join(tempDir, "%s")
	gexe.FileWrite(fullPath, filename).String("This is a demo file created with sprintf formatting")

	if gexe.PathExists(fullPath, filename) {
		content := gexe.FileRead(fullPath, filename).String()
		fmt.Printf("   Created file: %s\n", filepath.Join(tempDir, filename))
		fmt.Printf("   Content: %s\n", content)
	}

	// 5. Variable setting with formatting
	fmt.Println("\n5. Variable setting with formatting:")
	gexe.SetVar("status", "User %s has %d points", "Charlie", 150)
	fmt.Printf("   Variable value: %s\n", gexe.Val("status"))

	// 6. Backward compatibility (no formatting verbs)
	fmt.Println("\n6. Backward compatibility:")
	result = gexe.Run("echo Hello World", "unused_arg", "another_unused")
	fmt.Printf("   Result: %s (args ignored when no format verbs)\n", result)

	// 7. Complex example: log file creation
	fmt.Println("\n7. Complex example - log file creation:")
	user := "admin"
	action := "login"
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	logDir := filepath.Join(tempDir, "logs")
	gexe.SetVar("log_dir", logDir)
	gexe.MkDir("${log_dir}", 0755)

	logEntry := fmt.Sprintf("[%s] User %s performed %s", timestamp, user, action)
	gexe.FileWrite("${log_dir}/app_%s.log", time.Now().Format("20060102")).String(logEntry)

	fmt.Printf("   Log entry created: %s\n", logEntry)

	// Cleanup
	gexe.Run("rm -f %s/demo_*.txt", tempDir)
	gexe.Run("rm -rf %s", logDir)

	fmt.Println("\n=== Demo Complete ===")
}
