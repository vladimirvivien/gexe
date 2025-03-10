package gexe

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFileReadWrite(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test_writer_reader.txt")
	content := "Hello from file"
	defer os.RemoveAll(path)

	if err := FileWrite(path).String(content).Err(); err != nil {
		t.Fatal(err)
	}
	if FileRead(path).String() != content {
		t.Error("unexpected file content")
	}
}
