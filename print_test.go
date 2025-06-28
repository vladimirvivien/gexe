package gexe

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestPrintFunctionality(t *testing.T) {
	tests := []struct {
		name   string
		testFn func(t *testing.T)
	}{
		{"Print basic", testPrintBasic},
		{"Print with sprintf", testPrintSprintf},
		{"Print with variable expansion", testPrintVarExpansion},
		{"Print combined sprintf and variables", testPrintCombined},
		{"Println basic", testPrintlnBasic},
		{"Println with sprintf", testPrintlnSprintf},
		{"Println with variable expansion", testPrintlnVarExpansion},
		{"PrintTo basic", testPrintToBasic},
		{"PrintTo with sprintf", testPrintToSprintf},
		{"PrintTo with variable expansion", testPrintToVarExpansion},
		{"Print method chaining", testPrintChaining},
		{"Print backward compatibility", testPrintBackwardCompatibility},
		{"Print edge cases", testPrintEdgeCases},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.testFn)
	}
}

func testPrintBasic(t *testing.T) {
	// Capture stdout
	output := captureOutput(t, func() {
		g := New()
		result := g.Print("Hello World")
		if result != g {
			t.Error("Print should return the session for chaining")
		}
	})

	expected := "Hello World"
	if output != expected {
		t.Errorf("Print failed: expected %q, got %q", expected, output)
	}

	// Test package-level function
	output = captureOutput(t, func() {
		Print("Package Level")
	})

	expected = "Package Level"
	if output != expected {
		t.Errorf("Package-level Print failed: expected %q, got %q", expected, output)
	}
}

func testPrintSprintf(t *testing.T) {
	// Test with sprintf formatting
	output := captureOutput(t, func() {
		g := New()
		g.Print("Hello %s", "World")
	})

	expected := "Hello World"
	if output != expected {
		t.Errorf("Print with sprintf failed: expected %q, got %q", expected, output)
	}

	// Test with multiple args
	output = captureOutput(t, func() {
		g := New()
		g.Print("User: %s, ID: %d", "Alice", 42)
	})

	expected = "User: Alice, ID: 42"
	if output != expected {
		t.Errorf("Print with multiple sprintf args failed: expected %q, got %q", expected, output)
	}
}

func testPrintVarExpansion(t *testing.T) {
	output := captureOutput(t, func() {
		g := New()
		g.SetVar("name", "World")
		g.Print("Hello ${name}")
	})

	expected := "Hello World"
	if output != expected {
		t.Errorf("Print with variable expansion failed: expected %q, got %q", expected, output)
	}
}

func testPrintCombined(t *testing.T) {
	output := captureOutput(t, func() {
		g := New()
		g.SetVar("greeting", "Hello")
		g.Print("${greeting} %s", "World")
	})

	expected := "Hello World"
	if output != expected {
		t.Errorf("Print with combined sprintf and variables failed: expected %q, got %q", expected, output)
	}
}

func testPrintlnBasic(t *testing.T) {
	// Test Println adds newline
	output := captureOutput(t, func() {
		g := New()
		result := g.Println("Hello World")
		if result != g {
			t.Error("Println should return the session for chaining")
		}
	})

	expected := "Hello World\n"
	if output != expected {
		t.Errorf("Println failed: expected %q, got %q", expected, output)
	}

	// Test package-level function
	output = captureOutput(t, func() {
		Println("Package Level")
	})

	expected = "Package Level\n"
	if output != expected {
		t.Errorf("Package-level Println failed: expected %q, got %q", expected, output)
	}
}

func testPrintlnSprintf(t *testing.T) {
	output := captureOutput(t, func() {
		g := New()
		g.Println("Hello %s", "World")
	})

	expected := "Hello World\n"
	if output != expected {
		t.Errorf("Println with sprintf failed: expected %q, got %q", expected, output)
	}
}

func testPrintlnVarExpansion(t *testing.T) {
	output := captureOutput(t, func() {
		g := New()
		g.SetVar("name", "World")
		g.Println("Hello ${name}")
	})

	expected := "Hello World\n"
	if output != expected {
		t.Errorf("Println with variable expansion failed: expected %q, got %q", expected, output)
	}
}

func testPrintToBasic(t *testing.T) {
	var buf bytes.Buffer
	g := New()
	result := g.PrintTo(&buf, "Hello World")

	if result != g {
		t.Error("PrintTo should return the session for chaining")
	}

	expected := "Hello World"
	if buf.String() != expected {
		t.Errorf("PrintTo failed: expected %q, got %q", expected, buf.String())
	}

	// Test package-level function
	var buf2 bytes.Buffer
	PrintTo(&buf2, "Package Level")

	expected = "Package Level"
	if buf2.String() != expected {
		t.Errorf("Package-level PrintTo failed: expected %q, got %q", expected, buf2.String())
	}
}

func testPrintToSprintf(t *testing.T) {
	var buf bytes.Buffer
	g := New()
	g.PrintTo(&buf, "Hello %s", "World")

	expected := "Hello World"
	if buf.String() != expected {
		t.Errorf("PrintTo with sprintf failed: expected %q, got %q", expected, buf.String())
	}
}

func testPrintToVarExpansion(t *testing.T) {
	var buf bytes.Buffer
	g := New()
	g.SetVar("name", "World")
	g.PrintTo(&buf, "Hello ${name}")

	expected := "Hello World"
	if buf.String() != expected {
		t.Errorf("PrintTo with variable expansion failed: expected %q, got %q", expected, buf.String())
	}
}

func testPrintChaining(t *testing.T) {
	output := captureOutput(t, func() {
		g := New()
		g.SetVar("name", "Alice").
			Print("Hello ").
			Print("%s", "${name}").
			Println("!")
	})

	expected := "Hello Alice!\n"
	if output != expected {
		t.Errorf("Print chaining failed: expected %q, got %q", expected, output)
	}
}

func testPrintBackwardCompatibility(t *testing.T) {
	// Test no args
	output := captureOutput(t, func() {
		g := New()
		g.Print("Hello World")
	})

	expected := "Hello World"
	if output != expected {
		t.Errorf("Print backward compatibility failed: expected %q, got %q", expected, output)
	}

	// Test with variable expansion but no sprintf
	output = captureOutput(t, func() {
		g := New()
		g.SetVar("msg", "World")
		g.Print("Hello ${msg}")
	})

	expected = "Hello World"
	if output != expected {
		t.Errorf("Print with variables only failed: expected %q, got %q", expected, output)
	}
}

func testPrintEdgeCases(t *testing.T) {
	// Test empty string
	output := captureOutput(t, func() {
		g := New()
		g.Print("")
	})

	expected := ""
	if output != expected {
		t.Errorf("Print empty string failed: expected %q, got %q", expected, output)
	}

	// Test PrintTo with nil writer should panic gracefully
	defer func() {
		if r := recover(); r == nil {
			t.Error("PrintTo with nil writer should panic")
		}
	}()

	g := New()
	g.PrintTo(nil, "test")
}

// Helper function to capture stdout during test execution
func captureOutput(t *testing.T, fn func()) string {
	t.Helper()
	// Create a pipe to capture output
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}

	// Save original stdout
	originalStdout := os.Stdout

	// Redirect stdout to our pipe
	os.Stdout = w

	// Create a channel to receive the captured output
	outputChan := make(chan string)

	// Start a goroutine to read from the pipe
	go func(tt *testing.T) {
		var buf bytes.Buffer
		_, err := io.Copy(&buf, r)
		if err != nil {
			tt.Errorf("Failed to read from pipe: %v", err)
		}
		outputChan <- buf.String()
	}(t)

	// Execute the function
	fn()

	// Close the write end and restore stdout
	w.Close()
	os.Stdout = originalStdout

	// Get the captured output
	output := <-outputChan
	r.Close()

	return output
}
