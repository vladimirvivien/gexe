package prog

import (
	"os"
	"path/filepath"
	"testing"
)

func TestEchoProg(t *testing.T) {
	tests := []struct {
		name string
		exec func(*testing.T)
	}{
		{
			name: "test Info",
			exec: func(t *testing.T) {
				if Prog().Pid() != os.Getpid() {
					t.Errorf("expecting pid %d, got %d", os.Getpid(), Prog().Pid())
				}
				path, err := os.Executable()
				if err != nil {
					t.Fatal(err)
				}
				if Prog().Name() != filepath.Base(path) {
					t.Errorf("expecting %s, got %s", filepath.Base(path), path)
				}
				if Prog().Path() != path {
					t.Errorf("expecting path %s, got %s", path, Prog().Path())
				}
				if Prog().Ppid() != os.Getppid() {
					t.Errorf("expecting ppid %d, got %d", os.Getppid(), Prog().Ppid())
				}
				wd, err := os.Getwd()
				if err != nil {
					t.Fatal(err)
				}
				if Prog().Workdir() != wd {
					t.Errorf("expecting workdir %s, got %s", wd, Prog().Workdir())
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.exec(t)
		})
	}
}
