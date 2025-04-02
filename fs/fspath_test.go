package fs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewPath(t *testing.T) {
	p := Path("/tmp/sometest.txt")
	if p.path != "/tmp/sometest.txt" {
		t.Errorf("path not set")
	}
}

func TestPathMethods(t *testing.T) {
	tests := []struct {
		name string
		path string
		test func(t *testing.T, path string)
	}{
		{
			name: "path not exist",
			path: "/fake/path",
			test: func(t *testing.T, path string) {
				p := Path(path)
				if p.Exists() {
					t.Errorf("unexpected path exists: %s", path)
				}
			},
		},
		{
			name: "path exist",
			path: filepath.Join(t.TempDir(), "myfile.txt"),
			test: func(t *testing.T, path string) {
				if err := Write(path).String("hello").Err(); err != nil {
					t.Fatalf("unable to write test file: %s", err)
					return
				}
				defer os.RemoveAll(path)

				p := Path(path)
				if !p.Exists() {
					t.Errorf("path missing: %s", path)
				}
			},
		},
		{
			name: "make dir",
			path: filepath.Join(t.TempDir(), "mydir"),
			test: func(t *testing.T, path string) {
				p := Path(path).MkDir(0644)
				if p.Err() != nil {
					t.Fatalf("unable to create dir: %s", p.err)
					return
				}
				defer os.RemoveAll(path)
				if _, err := os.Stat(path); err != nil {
					t.Fatalf("dir not created: %s", err)
				}
			},
		},
		{
			name: "remove file",
			path: filepath.Join(t.TempDir(), "myfile.txt"),
			test: func(t *testing.T, path string) {
				if err := Write(path).String("hello").Err(); err != nil {
					t.Fatalf("unable to write test file: %s", err)
					return
				}

				p := Path(path).Remove()
				if p.Err() != nil {
					t.Fatalf("fail to remove file: %s", p.Err())
				}

				if _, err := os.Stat(path); err == nil {
					t.Fatalf("file was not removed")
				}
			},
		},
		{
			name: "count dir",
			path: filepath.Join(t.TempDir(), "mydir"),
			test: func(t *testing.T, path string) {
				if err := os.Mkdir(path, 0766); err != nil {
					t.Fatalf("failed to create dir for test: %s", err)
				}

				if err := Write(path + string(os.PathSeparator) + "file1.txt").String("hello1").Err(); err != nil {
					t.Fatalf("unable to write test file: %s", err)
					return
				}
				if err := Write(path + string(os.PathSeparator) + "file2.txt").String("hello2").Err(); err != nil {
					t.Fatalf("unable to write test file: %s", err)
					return
				}
				if err := Write(path + string(os.PathSeparator) + "file3.txt").String("hello3").Err(); err != nil {
					t.Fatalf("unable to write test file: %s", err)
					return
				}
				defer os.RemoveAll(path)
				if l := len(Path(path).Dirs()); l != 3 {
					t.Errorf("expecting 3 files in path, got %d", l)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.test(t, test.path)
			if err := os.RemoveAll(test.path); err != nil {
				t.Logf("warninig: unable to remove path: %s", err)
			}
		})
	}
}
