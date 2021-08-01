package fs

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestFileWriter(t *testing.T) {
	type testFile struct {
		path    string
		content string
	}
	tests := []struct {
		name  string
		write func(*testing.T) testFile
		test  func(*testing.T, testFile)
	}{
		{
			name: "write.string",
			write: func(t *testing.T) testFile {
				file := testFile{path: "/tmp/echo_test_write_string.txt", content: "Hello Write File"}
				f := Write(file.path).String(file.content)
				if f.Err() != nil {
					t.Fatal(f.Err())
				}
				return file
			},
			test: func(t *testing.T, file testFile) {
				actual := Read(file.path).String()
				if actual != file.content {
					t.Fatalf("Write().String unexpected result: wrote %s, got %s", file.content, actual)
				}
			},
		},
		{
			name: "write.lines",
			write: func(t *testing.T) testFile {
				file := testFile{path: "/tmp/echo_test_write_lines.txt", content: ""}
				f := Write(file.path).Lines([]string{"Hello", "gexe", "gexe"})
				if f.Err() != nil {
					t.Fatal(f.Err())
				}
				return file
			},
			test: func(t *testing.T, file testFile) {
				actual := Read(file.path).Lines()
				expected := []string{"Hello", "gexe", "gexe"}
				if len(actual) != len(expected) {
					t.Fatalf("Write().Lines unexpected length: want %d, got %d", len(expected), len(actual))
				}
			},
		},
		{
			name: "write.bytes",
			write: func(t *testing.T) testFile {
				file := testFile{path: "/tmp/echo_test_write_bytes.txt", content: "Hello Write File"}
				f := Write(file.path).Bytes([]byte(file.content))
				if f.Err() != nil {
					t.Fatal(f.Err())
				}
				return file
			},
			test: func(t *testing.T, file testFile) {
				actual := Read(file.path).String()
				if actual != file.content {
					t.Fatalf("Write().String unexpected result: wrote %s, got %s", file.content, actual)
				}
			},
		},
		{
			name: "write.readFrom",
			write: func(t *testing.T) testFile {
				file := testFile{path: "/tmp/echo_test_write_bytes.txt", content: "hello from buffer"}
				buf := bytes.NewBufferString(file.content)
				f := Write(file.path).From(buf)
				if f.Err() != nil {
					t.Fatal(f.Err())
				}
				return file
			},
			test: func(t *testing.T, file testFile) {
				actual := Read(file.path).String()
				if actual != file.content {
					t.Fatalf("Write().String unexpected result: wrote %s, got %s", file.content, actual)
				}
			},
		},
		{
			name: "write.truncate",
			write: func(t *testing.T) testFile {
				file := testFile{path: "/tmp/echo_test_write_truncate.txt", content: "Hello Write File"}
				f := Write(file.path).String("I will be truncated").String(file.content)
				if f.Err() != nil {
					t.Fatal(f.Err())
				}
				return file
			},
			test: func(t *testing.T, file testFile) {
				actual := Read(file.path).String()
				if actual != file.content {
					t.Fatalf("Write().String unexpected result: wrote %s, got %s", file.content, actual)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			file := test.write(t)
			test.test(t, file)
			if err := os.RemoveAll(file.path); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestFileAppender(t *testing.T) {
	type testFile struct {
		path    string
		content string
	}
	tests := []struct {
		name  string
		write func(*testing.T) testFile
		test  func(*testing.T, testFile)
	}{
		{
			name: "append.string",
			write: func(t *testing.T) testFile {
				file := testFile{path: "/tmp/echo_test_write_string.txt", content: "how are you?"}
				f := Write(file.path).String("Hello, ")
				if f.Err() != nil {
					t.Fatal(f.Err())
				}
				Append(file.path).String(file.content)
				return file
			},
			test: func(t *testing.T, file testFile) {
				actual := Read(file.path).String()
				if actual != fmt.Sprintf("Hello, %s", file.content) {
					t.Fatalf("Write().String unexpected result: wrote %s, got %s", file.content, actual)
				}
			},
		},
		{
			name: "append.lines",
			write: func(t *testing.T) testFile {
				file := testFile{path: "/tmp/echo_test_write_lines.txt", content: ""}
				f := Write(file.path).String("Alo\n")
				if f.Err() != nil {
					t.Fatal(f.Err())
				}
				Append(file.path).Lines([]string{"Hello", "gexe", "gexe"})
				return file
			},
			test: func(t *testing.T, file testFile) {
				actual := Read(file.path).Lines()
				expected := []string{"Alo", "Hello", "gexe", "gexe"}
				if len(actual) != len(expected) {
					t.Fatalf("Write().Lines unexpected length: want %d, got %d", len(expected), len(actual))
				}
			},
		},
		{
			name: "append.readFrom",
			write: func(t *testing.T) testFile {
				file := testFile{path: "/tmp/echo_test_write_bytes.txt", content: "hello from buffer"}
				buf := bytes.NewBufferString(file.content)
				f := Write(file.path).String("Hi! ")
				if f.Err() != nil {
					t.Fatal(f.Err())
				}
				Append(file.path).From(buf)
				return file
			},
			test: func(t *testing.T, file testFile) {
				actual := Read(file.path).String()
				if actual != fmt.Sprintf("Hi! %s", file.content) {
					t.Fatalf("Write().String unexpected result: wrote %s, got %s", file.content, actual)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			file := test.write(t)
			test.test(t, file)
			if err := os.RemoveAll(file.path); err != nil {
				t.Fatal(err)
			}
		})
	}
}
