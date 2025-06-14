package gexe

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSprintfFunctionality(t *testing.T) {
	tests := []struct {
		name   string
		testFn func(t *testing.T)
	}{
		{"Run with sprintf", testRunSprintf},
		{"SetVar with sprintf", testSetVarSprintf},
		{"SetEnv with sprintf", testSetEnvSprintf},
		{"Eval with sprintf", testEvalSprintf},
		{"FileRead with sprintf", testFileReadSprintf},
		{"FileWrite with sprintf", testFileWriteSprintf},
		{"PathExists with sprintf", testPathExistsSprintf},
		{"MkDir with sprintf", testMkDirSprintf},
		{"String with sprintf", testStringSprintf},
		{"ProgAvail with sprintf", testProgAvailSprintf},
		{"Combined sprintf and variable expansion", testCombinedSprintfAndVarExpansion},
		{"No args backward compatibility", testNoArgsBackwardCompatibility},
		{"Multiple args", testMultipleArgs},
		{"Edge cases", testEdgeCases},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.testFn)
	}
}

func testRunSprintf(t *testing.T) {
	g := New()

	// Test with string formatting
	result := g.Run("echo %s", "Hello World")
	expected := "Hello World"
	if result != expected {
		t.Errorf("Run with sprintf failed: expected %q, got %q", expected, result)
	}

	// Test with multiple args
	result = g.Run("echo %s %d", "Number:", 42)
	expected = "Number: 42"
	if result != expected {
		t.Errorf("Run with multiple sprintf args failed: expected %q, got %q", expected, result)
	}

	// Test package-level function
	result = Run("echo %s", "Package Level")
	expected = "Package Level"
	if result != expected {
		t.Errorf("Package-level Run with sprintf failed: expected %q, got %q", expected, result)
	}
}

func testSetVarSprintf(t *testing.T) {
	g := New()

	// Test SetVar with sprintf
	g.SetVar("user", "User: %s", "Alice")
	value := g.Val("user")
	expected := "User: Alice"
	if value != expected {
		t.Errorf("SetVar with sprintf failed: expected %q, got %q", expected, value)
	}

	// Test package-level function
	SetVar("count", "Count: %d", 100)
	value = Val("count")
	expected = "Count: 100"
	if value != expected {
		t.Errorf("Package-level SetVar with sprintf failed: expected %q, got %q", expected, value)
	}
}

func testSetEnvSprintf(t *testing.T) {
	g := New()

	// Test SetEnv with sprintf
	g.SetEnv("TEST_VAR", "Value: %s", "formatted")
	value := os.Getenv("TEST_VAR")
	expected := "Value: formatted"
	if value != expected {
		t.Errorf("SetEnv with sprintf failed: expected %q, got %q", expected, value)
	}

	// Cleanup
	os.Unsetenv("TEST_VAR")
}

func testEvalSprintf(t *testing.T) {
	g := New()
	g.SetVar("name", "World")

	// Test Eval with sprintf
	result := g.Eval("Hello %s: ${name}", "Formatted")
	expected := "Hello Formatted: World"
	if result != expected {
		t.Errorf("Eval with sprintf failed: expected %q, got %q", expected, result)
	}

	// Test package-level function (set variable in default session)
	SetVar("name", "World")
	result = Eval("Hello %s: ${name}", "Package")
	expected = "Hello Package: World"
	if result != expected {
		t.Errorf("Package-level Eval with sprintf failed: expected %q, got %q", expected, result)
	}
}

func testFileReadSprintf(t *testing.T) {
	g := New()

	// Create test file
	filename := "test_file.txt"
	g.FileWrite(filename).String("test content")
	defer os.Remove(filename)

	// Test FileRead with sprintf
	content := g.FileRead("%s", filename).String()
	expected := "test content"
	if content != expected {
		t.Errorf("FileRead with sprintf failed: expected %q, got %q", expected, content)
	}

	// Test package-level function
	content = FileRead("%s", filename).String()
	if content != expected {
		t.Errorf("Package-level FileRead with sprintf failed: expected %q, got %q", expected, content)
	}
}

func testFileWriteSprintf(t *testing.T) {
	g := New()

	// Test FileWrite with sprintf
	filename := "test_%s.txt"
	formattedName := "output"
	g.FileWrite(filename, formattedName).String("test content")

	expectedFilename := "test_output.txt"
	defer os.Remove(expectedFilename)

	// Verify file was created with correct name
	if !g.PathExists(expectedFilename) {
		t.Errorf("FileWrite with sprintf failed: file %q was not created", expectedFilename)
	}

	// Test package-level function
	filename2 := "test_%s_%d.txt"
	FileWrite(filename2, "package", 123).String("test content")
	expectedFilename2 := "test_package_123.txt"
	defer os.Remove(expectedFilename2)

	if !PathExists(expectedFilename2) {
		t.Errorf("Package-level FileWrite with sprintf failed: file %q was not created", expectedFilename2)
	}
}

func testPathExistsSprintf(t *testing.T) {
	g := New()

	// Use cross-platform temp directory
	tempDir := os.TempDir()

	// Test PathExists with sprintf
	exists := g.PathExists(tempDir)
	if !exists {
		t.Errorf("PathExists failed on temp dir %q", tempDir)
	}

	// Test with formatting - use current directory which should always exist
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	dirName := filepath.Base(currentDir)
	parentDir := filepath.Dir(currentDir)

	exists = g.PathExists(filepath.Join(parentDir, "%s"), dirName)
	if !exists {
		t.Errorf("PathExists with sprintf failed for path %q", filepath.Join(parentDir, dirName))
	}

	// Test non-existent path with sprintf
	exists = g.PathExists(filepath.Join(tempDir, "nonexistent_%s"), "path")
	if exists {
		t.Error("PathExists with sprintf should return false for non-existent path")
	}
}

func testMkDirSprintf(t *testing.T) {
	g := New()

	// Test MkDir with sprintf using cross-platform temp directory
	tempDir := os.TempDir()
	dirName := filepath.Join(tempDir, "test_%s")
	formattedDir := "sprintf"
	info := g.MkDir(dirName, 0755, formattedDir)
	expectedDir := filepath.Join(tempDir, "test_sprintf")

	if info.Err() != nil {
		t.Errorf("MkDir with sprintf failed: %v", info.Err())
	}

	defer os.RemoveAll(expectedDir)

	if !g.PathExists(expectedDir) {
		t.Errorf("MkDir with sprintf failed: directory %q was not created", expectedDir)
	}
}

func testStringSprintf(t *testing.T) {
	g := New()

	// Test String with sprintf
	str := g.String("Hello %s", "World").String()
	expected := "Hello World"
	if str != expected {
		t.Errorf("String with sprintf failed: expected %q, got %q", expected, str)
	}

	// Test package-level function
	str = String("Count: %d", 42).String()
	expected = "Count: 42"
	if str != expected {
		t.Errorf("Package-level String with sprintf failed: expected %q, got %q", expected, str)
	}
}

func testProgAvailSprintf(t *testing.T) {
	g := New()

	// Use a command that should be available on all platforms
	// Use go command which should be available since we're running go test
	path := g.ProgAvail("%s", "go")
	if path == "" {
		t.Error("ProgAvail with sprintf failed: go command should be available")
	}

	// Test package-level function
	path = ProgAvail("%s", "go")
	if path == "" {
		t.Error("Package-level ProgAvail with sprintf failed: go command should be available")
	}
}

func testCombinedSprintfAndVarExpansion(t *testing.T) {
	g := New()
	g.SetVar("user", "Alice")

	// Test combination of sprintf and variable expansion
	result := g.Run("echo Hello %s, your home is ${HOME}", "formatted")
	// Should contain both "Hello formatted" and the HOME value
	if !contains(result, "Hello formatted") {
		t.Errorf("Combined sprintf and var expansion failed: missing sprintf part in %q", result)
	}

	// Test with Eval
	result = g.Eval("User: %s, Name: ${user}", "Admin")
	expected := "User: Admin, Name: Alice"
	if result != expected {
		t.Errorf("Combined sprintf and var expansion in Eval failed: expected %q, got %q", expected, result)
	}
}

func testNoArgsBackwardCompatibility(t *testing.T) {
	g := New()

	// Test that calling without args works as before
	result := g.Run("echo Hello")
	expected := "Hello"
	if result != expected {
		t.Errorf("Backward compatibility failed: expected %q, got %q", expected, result)
	}

	// Test with variable expansion
	g.SetVar("msg", "World")
	result = g.Run("echo Hello ${msg}")
	expected = "Hello World"
	if result != expected {
		t.Errorf("Backward compatibility with var expansion failed: expected %q, got %q", expected, result)
	}
}

func testMultipleArgs(t *testing.T) {
	g := New()

	// Test with multiple formatting arguments
	result := g.Run("echo %s %d %f %t", "String", 42, 3.14, true)
	expected := "String 42 3.140000 true"
	if result != expected {
		t.Errorf("Multiple args sprintf failed: expected %q, got %q", expected, result)
	}

	// Test SetVar with multiple args
	g.SetVar("complex", "%s-%d-%s", "prefix", 123, "suffix")
	value := g.Val("complex")
	expected = "prefix-123-suffix"
	if value != expected {
		t.Errorf("SetVar with multiple args failed: expected %q, got %q", expected, value)
	}
}

func testEdgeCases(t *testing.T) {
	g := New()

	// Test with no formatting verbs but args provided - should ignore args
	result := g.Run("echo Hello", "unused")
	expected := "Hello"
	if result != expected {
		t.Errorf("Edge case (no verbs with args) failed: expected %q, got %q", expected, result)
	}

	// Test with empty string
	result = g.Eval("%s", "")
	expected = ""
	if result != expected {
		t.Errorf("Edge case (empty string) failed: expected %q, got %q", expected, result)
	}

	// Test with nil args (empty slice)
	result = g.Run("echo Hello")
	expected = "Hello"
	if result != expected {
		t.Errorf("Edge case (no args) failed: expected %q, got %q", expected, result)
	}
}

// Helper function
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > len(substr) && (s[:len(substr)] == substr ||
			s[len(s)-len(substr):] == substr ||
			indexOf(s, substr) >= 0)))
}

func indexOf(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
