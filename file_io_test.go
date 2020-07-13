package echo

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestFile_Read(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(t *testing.T) string
		test     func(*testing.T, string)
		teardown func(*testing.T, string)
	}{
		{
			name: "file read",
			setup: func(t *testing.T) string {
				name := "/tmp/echo_test_readfile.txt"
				if err := ioutil.WriteFile(name, []byte("Hello from echo"), 0744); err != nil {
					t.Error(err)
				}
				return name
			},
			test: func(t *testing.T, filename string) {
				e := New()
				f := e.OpenFile(filename)
				if f.Err() != nil {
					t.Error(f.Err())
				}
				expected := "Hello from echo"
				actual := strings.TrimSpace(f.Read())
				if actual != expected {
					t.Errorf("unexpected result for File.Read: %s", actual)
				}
			},
			teardown: func(t *testing.T, fileName string) {
				if err := os.RemoveAll(fileName); err != nil {
					t.Error(err)
				}
			},
		},

		{
			name: "file read lines",
			setup: func(t *testing.T) string {
				name := "/tmp/echo_test_readfile_lines.txt"
				if err := ioutil.WriteFile(name, []byte("Hello from\necho\necho\necho"), 0744); err != nil {
					t.Error(err)
				}
				return name
			},
			test: func(t *testing.T, filename string) {
				e := New()
				f := e.OpenFile(filename)
				if f.Err() != nil {
					t.Error(f.Err())
				}
				expecteds := []string{"Hello from", "echo", "echo", "echo"}
				actuals := f.ReadLines()
				for i := range expecteds {
					if expecteds[i] != strings.TrimSpace(actuals[i]) {
						t.Errorf("File.ReadLines expecting %s but got %s", expecteds[i], actuals[i])
					}
				}
			},
			teardown: func(t *testing.T, fileName string) {
				if err := os.RemoveAll(fileName); err != nil {
					t.Error(err)
				}
			},
		},

		{
			name: "file read bytes",
			setup: func(t *testing.T) string {
				name := "/tmp/echo_test_readfile_bytes.txt"
				if err := ioutil.WriteFile(name, []byte("Hello from echo"), 0744); err != nil {
					t.Error(err)
				}
				return name
			},
			test: func(t *testing.T, filename string) {
				e := New()
				f := e.OpenFile(filename)
				if f.Err() != nil {
					t.Error(f.Err())
				}
				expected := "Hello from echo"
				actual := strings.TrimSpace(string(f.ReadBytes()))
				if actual != expected {
					t.Errorf("unexpected result for File.Read: %s", actual)
				}
			},
			teardown: func(t *testing.T, fileName string) {
				if err := os.RemoveAll(fileName); err != nil {
					t.Error(err)
				}
			},
		},

		{
			name: "file reader",
			setup: func(t *testing.T) string {
				name := "/tmp/echo_test_filereader.txt"
				if err := ioutil.WriteFile(name, []byte("Hello from echo"), 0744); err != nil {
					t.Error(err)
				}
				return name
			},
			test: func(t *testing.T, filename string) {
				e := New()
				f := e.OpenFile(filename)
				if f.Err() != nil {
					t.Error(f.Err())
				}
				rdr := f.Reader()
				if f.Err() != nil {
					t.Fatal(f.Err())
				}
				defer rdr.Close()
				expected := "Hello from echo"
				buf := new(bytes.Buffer)
				if _, err := io.Copy(buf, f.Reader()); err != nil {
					t.Fatal(err)
				}
				actual := strings.TrimSpace(buf.String())
				if actual != expected {
					t.Errorf("unexpected result for File.Read: %s", actual)
				}
			},
			teardown: func(t *testing.T, fileName string) {
				if err := os.RemoveAll(fileName); err != nil {
					t.Error(err)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			filename := test.setup(t)
			test.test(t, filename)
			test.teardown(t, filename)
		})
	}
}

func TestFile_Write(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(t *testing.T) string
		test     func(*testing.T, string)
		teardown func(*testing.T, string)
	}{
		{
			name: "read write",
			setup: func(t *testing.T) string {
				name := "/tmp/echo_test_writefile.txt"
				e := New()
				f := e.OpenFile(name)
				f.Write("Hello from echo")
				if f.Err() != nil {
					t.Fatal(f.Err())
				}
				return name
			},
			test: func(t *testing.T, filename string) {
				e := New()
				actual := e.OpenFile(filename).Read()
				expected := "Hello from echo"
				if actual != expected {
					t.Errorf("unexpected result for File.Write: %s", actual)
				}
			},
			teardown: func(t *testing.T, fileName string) {
				if err := os.RemoveAll(fileName); err != nil {
					t.Error(err)
				}
			},
		},
		{
			name: "write append read",
			setup: func(t *testing.T) string {
				name := "/tmp/echo_test_appendfile.txt"
				e := New()
				f := e.OpenFile(name)

				if f.Write("Hello"); f.Err() != nil{
					t.Fatal(f.Err())
				}

				f.Append(" from echo")
				if f.Err() != nil {
					t.Fatal(f.Err())
				}
				return name
			},
			test: func(t *testing.T, filename string) {
				e := New()
				actual := e.OpenFile(filename).Read()
				expected := "Hello from echo"
				if actual != expected {
					t.Errorf("unexpected result for File.Write: %s", actual)
				}
			},
			teardown: func(t *testing.T, fileName string) {
				if err := os.RemoveAll(fileName); err != nil {
					t.Error(err)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			filename := test.setup(t)
			test.test(t, filename)
			test.teardown(t, filename)
		})
	}
}

func TestFile_Stream(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(t *testing.T) string
		test     func(*testing.T, string)
		teardown func(*testing.T, string)
	}{
		{
			name: "stream from",
			setup: func(t *testing.T) string {
				return "/tmp/echo_test_copyfile.txt"
			},
			test: func(t *testing.T, filename string) {
				New().OpenFile(filename).StreamFrom(strings.NewReader("Hello World!"))
				expected := "Hello World!"
				actual := New().OpenFile(filename).Read()
				if actual != expected {
					t.Errorf("unexpected result for File.StreamFrom: %s", actual)
				}
			},
			teardown: func(t *testing.T, fileName string) {
				if err := os.RemoveAll(fileName); err != nil {
					t.Error(err)
				}
			},
		},

		{
			name: "stream to",
			setup: func(t *testing.T) string {
				return "/tmp/echo_test_copyfile.txt"
			},
			test: func(t *testing.T, filename string) {
				New().OpenFile(filename).Write("Hello from Echo!")
				buf := new(bytes.Buffer)
				New().OpenFile(filename).StreamTo(buf)
				expected := "Hello from Echo!"
				actual := strings.TrimSpace(buf.String())
				if actual != expected {
					t.Errorf("unexpected result for File.StreamTo: %s", actual)
				}
			},
			teardown: func(t *testing.T, fileName string) {
				if err := os.RemoveAll(fileName); err != nil {
					t.Error(err)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			filename := test.setup(t)
			test.test(t, filename)
			test.teardown(t, filename)
		})
	}
}