package echo

import (
	"os"
	"path/filepath"
	"testing"
)

func TestEchoProg(t *testing.T) {
	tests := []struct {
		name string
		exec func()
	}{
		{
			name: "test prog",
			exec: func() {
				e := New()
				if e.Prog.Pid() != os.Getpid() {
					t.Errorf("expecting pid %d, got %d", os.Getpid(), e.Prog.Pid())
				}
				path, err := os.Executable()
				if err != nil {
					t.Fatal(err)
				}
				if e.Prog.Name() != filepath.Base(path) {
					t.Errorf("expecting %s, got %s", filepath.Base(path), path)
				}
				if e.Prog.Path() != path {
					t.Errorf("expecting path %s, got %s", path, e.Prog.Path())
				}
				if e.Prog.Pid() != os.Getpid() {
					t.Errorf("expecting pid %d, got %d", os.Getpid(), e.Prog.Pid())
				}
				if e.Prog.Ppid() != os.Getppid() {
					t.Errorf("expecting ppid %d, got %d", os.Getppid(), e.Prog.Ppid())
				}
				wd, err := os.Getwd()
				if err != nil {
					t.Fatal(err)
				}
				if e.Prog.Workdir() != wd {
					t.Errorf("expecting workdir %s, got %s", wd, e.Prog.Workdir())
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.exec()
		})
	}
}
