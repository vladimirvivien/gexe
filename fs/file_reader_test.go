package fs

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestFileReader(t *testing.T) {
	type testFile struct {
		path    string
		content string
	}
	tests := []struct {
		name string
		file testFile
		test func(*testing.T, testFile)
	}{
		{
			name: "read.string",
			file: testFile{path: "/tmp/echo_test_read_string.txt", content: "Hello from gexe"},
			test: func(t *testing.T, file testFile) {
				fr := Read(file.path)
				if fr.Err() != nil {
					t.Fatal(fr.Err())
				}
				actual := strings.TrimSpace(fr.String())
				if actual != file.content {
					t.Errorf("Read().String(): unexpected result: want %s, got %s", file.content, actual)
				}
			},
		},
		{
			name: "read.lines",
			file: testFile{path: "/tmp/echo_test_read_lines.txt", content: "Hello from\ngexe\ngexe\ngexe"},
			test: func(t *testing.T, file testFile) {
				fr := Read(file.path)
				if fr.Err() != nil {
					t.Fatal(fr.Err())
				}
				actual := fr.Lines()
				expected := []string{"Hello from", "gexe", "gexe", "gexe"}
				if len(actual) != len(expected) {
					t.Errorf("Read().Lines(): unexpected length: want %d, got %d", len(expected), len(actual))
				}

				for i, str := range actual {
					if str != expected[i] {
						t.Errorf("Read().Lines(): unexpected item: want %s, got %s", expected[i], str)
					}
				}
			},
		},
		{
			name: "read.bytes",
			file: testFile{path: "/tmp/echo_test_read_bytes.txt", content: "Hello from gexe"},
			test: func(t *testing.T, file testFile) {
				fr := Read(file.path)
				if fr.Err() != nil {
					t.Fatal(fr.Err())
				}
				actual := strings.TrimSpace(string(fr.Bytes()))
				if actual != file.content {
					t.Errorf("Read().String(): unexpected result: want %s, got %s", file.content, actual)
				}
			},
		},
		{
			name: "read.writeTo",
			file: testFile{path: "/tmp/echo_test_read_writeTo.txt", content: "Hello from gexe"},
			test: func(t *testing.T, file testFile) {
				buf := new(bytes.Buffer)
				fr := Read(file.path).Into(buf)
				if fr.Err() != nil {
					t.Fatal(fr.Err())
				}
				actual := strings.TrimSpace(buf.String())
				if actual != file.content {
					t.Errorf("Read().String(): unexpected result: want %s, got %s", file.content, actual)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := os.WriteFile(test.file.path, []byte(test.file.content), 0744); err != nil {
				t.Fatal(err)
			}
			test.test(t, test.file)
			if err := os.RemoveAll(test.file.path); err != nil {
				t.Fatal(err)
			}
		})
	}
}
