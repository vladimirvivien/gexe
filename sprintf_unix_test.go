//go:build !windows

package gexe

import (
	"testing"
)

func TestSprintfFunctionalityUnix(t *testing.T) {
	tests := []struct {
		name   string
		testFn func(t *testing.T)
	}{
		{"Unix ProgAvail with sprintf", testProgAvailSprintfUnix},
		{"Unix Run with sprintf", testRunSprintfUnix},
		{"Unix PathExists with /tmp", testPathExistsSprintfUnix},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.testFn)
	}
}

func testProgAvailSprintfUnix(t *testing.T) {
	g := New()

	// Test ProgAvail with sprintf - using Unix-specific commands
	path := g.ProgAvail("%s", "ls")
	if path == "" {
		t.Error("ProgAvail with sprintf failed: ls command should be available on Unix")
	}

	// Test package-level function
	path = ProgAvail("%s", "echo")
	if path == "" {
		t.Error("Package-level ProgAvail with sprintf failed: echo command should be available on Unix")
	}
}

func testRunSprintfUnix(t *testing.T) {
	g := New()

	// Test with Unix echo command
	result := g.Run("echo %s", "Hello World")
	expected := "Hello World"
	if result != expected {
		t.Errorf("Unix Run with sprintf failed: expected %q, got %q", expected, result)
	}

	// Test with ls command
	result = g.Run("echo %s", "Unix Test")
	expected = "Unix Test"
	if result != expected {
		t.Errorf("Unix echo with sprintf failed: expected %q, got %q", expected, result)
	}
}

func testPathExistsSprintfUnix(t *testing.T) {
	g := New()

	// Test PathExists with /tmp on Unix systems
	exists := g.PathExists("/tmp")
	if !exists {
		t.Error("PathExists failed on /tmp")
	}

	// Test with formatting
	exists = g.PathExists("/%s", "tmp")
	if !exists {
		t.Error("PathExists with sprintf failed")
	}
}
