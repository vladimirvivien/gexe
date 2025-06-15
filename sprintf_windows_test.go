//go:build windows

package gexe

import (
	"testing"
)

func TestSprintfFunctionalityWindows(t *testing.T) {
	tests := []struct {
		name   string
		testFn func(t *testing.T)
	}{
		{"Windows ProgAvail with sprintf", testProgAvailSprintfWindows},
		{"Windows Run with sprintf", testRunSprintfWindows},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.testFn)
	}
}

func testProgAvailSprintfWindows(t *testing.T) {
	g := New()

	// Test ProgAvail with sprintf - using Windows-specific commands
	path := g.ProgAvail("%s", "cmd")
	if path == "" {
		t.Error("ProgAvail with sprintf failed: cmd command should be available on Windows")
	}

	// Test package-level function
	path = ProgAvail("%s", "powershell")
	if path == "" {
		t.Skip("PowerShell not available, skipping test")
	}
}

func testRunSprintfWindows(t *testing.T) {
	g := New()

	// Test with Windows echo command
	result := g.Run("echo %s", "Hello World")
	expected := "Hello World"
	if result != expected {
		t.Errorf("Windows Run with sprintf failed: expected %q, got %q", expected, result)
	}

	// Test with dir command (Windows equivalent of ls)
	result = g.Run("cmd /c echo %s", "Windows Test")
	expected = "Windows Test"
	if result != expected {
		t.Errorf("Windows cmd with sprintf failed: expected %q, got %q", expected, result)
	}
}
