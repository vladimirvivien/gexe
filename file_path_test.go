package echo

import (
	"os"
	"testing"
)

func TestEchoMkdirs(t *testing.T) {
	tests := []struct {
		name  string
		paths string
	}{
		{
			name:  "single dir",
			paths: "/tmp/echo-test",
		},
		{
			name:  "multiple  dirs",
			paths: "/tmp/echo-test /tmp/test-echo",
		},
		{
			name:  "dirs with expansion",
			paths: "/tmp/echo-test $HOME/echo-test",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e := New()
			e.Mkdirs(test.paths)
			for _, path := range e.Split(test.paths) {
				if _, err := os.Stat(path); err != nil {
					t.Errorf("Faild to Mkdirs directory %s: %s", path, err)
				}
				t.Log("removing dir", path)
				if err := os.RemoveAll(path); err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}

func TestEchoRmdirs(t *testing.T) {
	tests := []struct {
		name  string
		paths string
	}{
		{
			name:  "single dir",
			paths: "/tmp/echo-test",
		},
		{
			name:  "multiple  dirs",
			paths: "/tmp/echo-test /tmp/test-echo",
		},
		{
			name:  "dirs with expansion",
			paths: "/tmp/echo-test $HOME/echo-test",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e := New()
			e.Mkdirs(test.paths)
			e.Rmdirs(test.paths)
			for _, path := range e.Split(test.paths) {
				if _, err := os.Stat(path); err == nil {
					t.Errorf("Expected stat error after removing %s", path)
				}
			}
		})
	}
}
