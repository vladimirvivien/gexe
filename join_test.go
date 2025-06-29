//go:build !windows

package gexe

import (
	"testing"
)

func TestJoin(t *testing.T) {
	tests := []struct {
		name      string
		vars      map[string]string
		separator string
		elements  []string
		expected  string
	}{
		{
			name:      "basic string join",
			separator: ",",
			elements:  []string{"apple", "banana", "cherry"},
			expected:  "apple,banana,cherry",
		},
		{
			name:      "join with single variable",
			vars:      map[string]string{"FRUIT": "apple"},
			separator: "-",
			elements:  []string{"${FRUIT}", "banana", "cherry"},
			expected:  "apple-banana-cherry",
		},
		{
			name:      "join with multiple variables",
			vars:      map[string]string{"FIRST": "hello", "SECOND": "world"},
			separator: " ",
			elements:  []string{"${FIRST}", "${SECOND}", "!"},
			expected:  "hello world !",
		},
		{
			name:      "join with empty separator",
			vars:      map[string]string{"PREFIX": "test"},
			separator: "",
			elements:  []string{"${PREFIX}", "123", "suffix"},
			expected:  "test123suffix",
		},
		{
			name:      "empty elements",
			separator: ",",
			elements:  []string{},
			expected:  "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := New()
			for key, value := range test.vars {
				g.SetVar(key, value)
			}

			result := g.Join(test.separator, test.elements...)
			if result != test.expected {
				t.Errorf("Join() = %q, expected %q", result, test.expected)
			}
		})
	}
}

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
			expected: "home/user/documents",
		},
		{
			name:     "path with single variable",
			vars:     map[string]string{"HOME": "/home/user"},
			elements: []string{"${HOME}", "documents", "file.txt"},
			expected: "/home/user/documents/file.txt",
		},
		{
			name:     "path with multiple variables",
			vars:     map[string]string{"ROOT": "/opt", "APP": "myapp"},
			elements: []string{"${ROOT}", "${APP}", "config", "settings.json"},
			expected: "/opt/myapp/config/settings.json",
		},
		{
			name:     "relative path",
			vars:     map[string]string{"DIR": "src"},
			elements: []string{"${DIR}", "main", "main.go"},
			expected: "src/main/main.go",
		},
		{
			name:     "empty elements",
			elements: []string{},
			expected: "",
		},
		{
			name:     "variable with nested path",
			vars:     map[string]string{"NESTED": "some/nested/path"},
			elements: []string{"root", "${NESTED}", "file.txt"},
			expected: "root/some/nested/path/file.txt",
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
