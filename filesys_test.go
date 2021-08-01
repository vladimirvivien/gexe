package gexe

import (
	"os"
	"testing"
)

func TestFileReader(t *testing.T) {
	path := "/tmp/test_writer_reader.txt"
	content := "Hello from file"
	defer os.RemoveAll(path)

	if err := Write(path).String(content).Err(); err != nil {
		t.Fatal(err)
	}
	if Read(path).String() != content {
		t.Error("unexpected file content")
	}
}
