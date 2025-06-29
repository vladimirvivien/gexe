package gexe

import (
	"testing"
)

func TestJoinPath(t *testing.T) {
	tests := []struct {
		name     string
		vars     map[string]string
		elements []string
		expected string
	}{
		{
			name:     "basic path join",
			elements: []string{"home", "user", "documents"},
			expected: "home\\user\\documents",
		},
		{
			name:     "path with single variable",
			vars:     map[string]string{"HOME": "C:\\Users\\user"},
			elements: []string{"${HOME}", "documents", "file.txt"},
			expected: "C:\\Users\\user\\documents\\file.txt",
		},
		{
			name:     "path with multiple variables",
			vars:     map[string]string{"ROOT": "C:\\Program Files", "APP": "myapp"},
			elements: []string{"${ROOT}", "${APP}", "config", "settings.json"},
			expected: "C:\\Program Files\\myapp\\config\\settings.json",
		},
		{
			name:     "relative path",
			vars:     map[string]string{"DIR": "src"},
			elements: []string{"${DIR}", "main", "main.go"},
			expected: "src\\main\\main.go",
		},
		{
			name:     "empty elements",
			elements: []string{},
			expected: "",
		},
		{
			name:     "variable with nested path",
			vars:     map[string]string{"NESTED": "some\\nested\\path"},
			elements: []string{"root", "${NESTED}", "file.txt"},
			expected: "root\\some\\nested\\path\\file.txt",
		},
		{
			name:     "drive letter paths",
			vars:     map[string]string{"DRIVE": "D:\\", "FOLDER": "data"},
			elements: []string{"${DRIVE}", "${FOLDER}", "files", "test.txt"},
			expected: "D:\\data\\files\\test.txt",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := New()
			for key, value := range test.vars {
				g.SetVar(key, value)
			}

			result := g.JoinPath(test.elements...)
			if result != test.expected {
				t.Errorf("JoinPath() = %q, expected %q", result, test.expected)
			}
		})
	}
}
