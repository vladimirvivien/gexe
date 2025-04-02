//go:build windows

package exec

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/vladimirvivien/gexe/vars"
)

func TestNewProc(t *testing.T) {
	tests := []struct {
		name   string
		cmdStr string
		exec   func(*testing.T, string)
	}{
		{
			name:   "access result",
			cmdStr: `powershell -Command "Write-Output 'Hello World'"`,
			exec: func(t *testing.T, cmd string) {
				proc := NewProc(cmd)
				if err := proc.Start().Err(); err != nil {
					t.Fatal(err)
				}
				if err := proc.Wait().Err(); err != nil {
					t.Fatal(err)
				}

				if strings.TrimSpace(proc.Result()) != "Hello World" {
					t.Errorf("unexpected result: %s", proc.Result())
				}
			},
		},
		{
			name:   "access proc.Out()",
			cmdStr: `powershell -Command "Write-Output 'Hello World'"`,
			exec: func(t *testing.T, cmd string) {
				proc := NewProc(cmd)
				if err := proc.Start().Wait().Err(); err != nil {
					t.Fatal(err)
				}

				buf := new(bytes.Buffer)
				if _, err := buf.ReadFrom(proc.Out()); err != nil {
					t.Fatal(err)
				}

				if strings.TrimSpace(buf.String()) != "Hello World" {
					t.Errorf("unexpected result: %s", buf.String())
				}
			},
		},
		{
			name:   "with expansion",
			cmdStr: `powershell -Command "Write-Output '$MSG'"`,
			exec: func(t *testing.T, cmd string) {
				proc := NewProcWithVars(cmd, vars.New().SetVar("MSG", "Hello World"))
				if err := proc.Start().Err(); err != nil {
					t.Fatal(err)
				}
				if err := proc.Wait().Err(); err != nil {
					t.Fatal(err)
				}
				if strings.TrimSpace(proc.Result()) != "Hello World" {
					t.Errorf("unexpected result: %s", proc.Result())
				}
			},
		},
		{
			name:   "custom stdout/stderr",
			cmdStr: `powershell -Command "Write-Output 'Hello World'"`,
			exec: func(t *testing.T, cmd string) {
				proc := NewProc(cmd)
				buf := new(bytes.Buffer)
				proc.SetStdout(buf)
				proc.SetStderr(buf)

				if err := proc.Start().Wait().Err(); err != nil {
					t.Fatalf("Failed to start/wait for proc: %s", err)
				}

				if strings.TrimSpace(buf.String()) != "Hello World" {
					t.Errorf("unexpected result: %s", buf.String())
				}
			},
		},
		{
			name:   "start with error",
			cmdStr: `foo "Hello World"`,
			exec: func(t *testing.T, cmd string) {
				proc := NewProc(cmd)
				if err := proc.Start().Err(); err == nil {
					t.Fatalf("Expecting error, but got none")
				}

				if err := proc.Wait().Err(); err == nil {
					t.Fatal("Expecting error, but got none")
				}

				if strings.TrimSpace(proc.Result()) == "Hello World" {
					t.Errorf("Expecting error but did not get it")
				}
			},
		},
		{
			name:   "run with result",
			cmdStr: `powershell -Command "Write-Output 'Hello World'"`,
			exec: func(t *testing.T, cmd string) {
				proc := NewProc(cmd)
				if err := proc.Run().Err(); err != nil {
					t.Fatal(err)
				}
				if strings.TrimSpace(proc.Result()) != "Hello World" {
					t.Errorf("unexpected result: %s", proc.Result())
				}
			},
		},
		{
			name:   "run with proc.Out()",
			cmdStr: `powershell -Command "Write-Output 'Hello World'"`,
			exec: func(t *testing.T, cmd string) {
				proc := NewProc(cmd)
				if err := proc.Run().Err(); err != nil {
					t.Fatal(err)
				}
				buf := new(bytes.Buffer)
				if _, err := buf.ReadFrom(proc.Out()); err != nil {
					t.Fatal(err)
				}

				if strings.TrimSpace(buf.String()) != "Hello World" {
					t.Errorf("unexpected result: %s", buf.String())
				}
			},
		},
		{
			name:   "missing command with Result",
			cmdStr: `foobar "Hello World"`,
			exec: func(t *testing.T, cmd string) {
				proc := NewProc(cmd)

				if err := proc.Run().Err(); err == nil {
					t.Fatalf("Expecting error, but got none")
				}
				result := proc.Result()
				if strings.TrimSpace(result) == "Hello World" {
					t.Errorf("Expecting error but did not get it")
				}
				if !strings.Contains(result, "not recognized") && !strings.Contains(result, "not found") {
					t.Errorf("Expecting result with error message, got: %s", result)
				}
			},
		},
		{
			name:   "missing command with proc.Out",
			cmdStr: `foobar "Hello World"`,
			exec: func(t *testing.T, cmd string) {
				proc := NewProc(cmd)

				if err := proc.Run().Err(); err == nil {
					t.Fatalf("Expecting error, but got none")
				}
				result := proc.Result()
				if !strings.Contains(result, "not recognized") && !strings.Contains(result, "not found") {
					t.Errorf("Expecting result with error message, got: %s", result)
				}
			},
		},
		{
			name:   "bad command with result",
			cmdStr: `powershell -Command "Get-Date -InvalidParameter"`,
			exec: func(t *testing.T, cmd string) {
				proc := NewProc(cmd)

				err := proc.Run().Err()
				if err == nil {
					t.Fatalf("Expecting error, but got none")
				}

				result := strings.TrimSpace(proc.Result())
				if !strings.Contains(strings.ToLower(result), "parameter") {
					t.Errorf("Expecting error but did not get it")
				}
			},
		},
		{
			name:   "bad command with proc.Out",
			cmdStr: `powershell -Command "Get-Date -InvalidParameter"`,
			exec: func(t *testing.T, cmd string) {
				proc := NewProc(cmd)

				err := proc.Run().Err()
				if err == nil {
					t.Fatalf("Expecting error, but got none")
				}

				buf := new(bytes.Buffer)
				if _, err := buf.ReadFrom(proc.Out()); err != nil {
					t.Fatal(err)
				}

				result := strings.TrimSpace(buf.String())
				if !strings.Contains(strings.ToLower(result), "parameter") {
					t.Errorf("Expecting error message, but got: %s", result)
				}
			},
		},
		{
			name:   "proc status",
			cmdStr: `powershell -Command "Write-Output 'HELLO WORLD!'"`,
			exec: func(t *testing.T, cmd string) {
				p := NewProc(cmd).Start()

				if p.Err() != nil {
					t.Fatal("Unexpected error:", p.Err().Error())
				}
				if p.ExitCode() != -1 {
					t.Fatal("Expecting -1, got:", p.ExitCode())
				}
				if p.IsSuccess() {
					t.Fatal("Success should be false")
				}
				if !p.hasStarted() {
					t.Fatal("proc has not started")
				}

				// wait for proc to finish
				if err := p.Wait().Err(); err != nil {
					t.Fatalf("failed to wait for proc to finish: %s", err)
				}

				if p.Result() != "HELLO WORLD!" {
					t.Errorf("Unexpected proc.Result(): %s", p.Result())
				}

				if p.ExitCode() != 0 {
					t.Fatal("Expecting exit code 0, got:", p.ExitCode())
				}
				if !p.IsSuccess() {
					t.Fatal("Process should be success")
				}
			},
		},
		{
			name:   "with working dir",
			cmdStr: `powershell -Command "New-Item -Path 'testfile.txt' -ItemType 'file' -Force"`,
			exec: func(t *testing.T, cmd string) {
				dir, err := os.MkdirTemp("", "TESTDIR")
				if err != nil {
					t.Fatalf("Failed to create temp dir: %s", err)
				}
				defer os.RemoveAll(dir)

				p := NewProc(cmd).SetWorkDir(dir)
				if p.cmd.Dir != dir {
					t.Fatalf("Not setting working dir")
				}

				run := p.Run()
				if err := run.Err(); err != nil {
					t.Fatalf("failed to run: %s", err)
				}
				if _, err := os.Stat(filepath.Join(dir, "testfile.txt")); err != nil {
					t.Fatalf("Unexpected error looking for file: %s", err)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.exec(t, test.cmdStr)
		})
	}
}

func TestStartProc(t *testing.T) {
	tests := []struct {
		name   string
		cmdStr string
		exec   func(*testing.T, string)
	}{
		{
			name:   "access result",
			cmdStr: `powershell -Command "Write-Output 'Hello World'"`,
			exec: func(t *testing.T, cmd string) {
				proc := StartProc(cmd).Wait()
				if err := proc.Err(); err != nil {
					t.Fatalf("failed to start/wait proc: %s", err)
				}
				if strings.TrimSpace(proc.Result()) != "Hello World" {
					t.Errorf("unexpected result: %s", proc.Result())
				}
			},
		},
		{
			name:   "access proc.Out",
			cmdStr: `powershell -Command "Write-Output 'Hello World'"`,
			exec: func(t *testing.T, cmd string) {
				proc := StartProc(cmd)

				if err := proc.Wait().Err(); err != nil {
					t.Fatalf("failed to wait: %s", err)
				}

				buf := new(bytes.Buffer)

				if _, err := buf.ReadFrom(proc.Out()); err != nil {
					t.Error(err)
				}
				if strings.TrimSpace(buf.String()) != "Hello World" {
					t.Errorf("unexpected result: %s", buf.String())
				}
			},
		},
		{
			name:   "with expansion",
			cmdStr: `powershell -Command "Write-Output '$MSG'"`,
			exec: func(t *testing.T, cmd string) {
				proc := StartProcWithVars(cmd, vars.New().SetVar("MSG", "Hello World")).Wait()
				if err := proc.Err(); err != nil {
					t.Fatal(err)
				}
				if strings.TrimSpace(proc.Result()) != "Hello World" {
					t.Errorf("unexpected result: %s", proc.Result())
				}
			},
		},
		{
			name:   "test proc status",
			cmdStr: `powershell -Command "Write-Output 'HELLO WORLD!'"`,
			exec: func(t *testing.T, cmd string) {
				p := StartProc(cmd)
				if p.Err() != nil {
					t.Fatal("Unexpected error:", p.Err().Error())
				}
				if p.ExitCode() != -1 {
					t.Fatal("Expecting -1, got:", p.ExitCode())
				}
				if p.IsSuccess() {
					t.Fatal("Success should be false")
				}
				if !p.hasStarted() {
					t.Fatal("proc has not started")
				}

				// wait for proc to finish
				if err := p.Wait().Err(); err != nil {
					t.Fatalf("failed to wait for proc to finish: %s", err)
				}

				if p.Result() != "HELLO WORLD!" {
					t.Errorf("Unexpected proc.Result(): %s", p.Result())
				}

				if p.ExitCode() != 0 {
					t.Fatal("Expecting exit code 0, got:", p.ExitCode())
				}
				if !p.IsSuccess() {
					t.Fatal("Process should be success")
				}
			},
		},
		{
			name:   "long-running",
			cmdStr: `powershell -Command "for($i=1; $i -le 3; $i++) { Write-Output 'HELLO WORLD!'; Start-Sleep -m 700 }"`,
			exec: func(t *testing.T, cmd string) {
				p := StartProc(cmd).Wait()

				if p.Err() != nil {
					t.Fatal(p.Err())
				}

				data := &bytes.Buffer{}
				if _, err := data.ReadFrom(p.Out()); err != nil {
					t.Fatal(err)
				}

				result := strings.TrimSpace(data.String())
				lines := strings.Split(result, "\n")
				if len(lines) != 3 {
					t.Fatal("unexpected lines generated:", len(lines))
				}
				if !strings.Contains(result, "HELLO WORLD!") {
					t.Fatal("Unexpected result:", result)
				}
			},
		},
		{
			name:   "missing command with result",
			cmdStr: `foobar "HELLO WORLD!"`,
			exec: func(t *testing.T, cmd string) {
				proc := StartProc(cmd).Wait()
				if proc.Err() == nil {
					t.Error("expecting command to fail")
				}
				result := proc.Result()
				if !strings.Contains(result, "not recognized") && !strings.Contains(result, "not found") {
					t.Errorf("Expecting result with error message, got: %s", result)
				}
			},
		},
		{
			name:   "missing command with proc.Out",
			cmdStr: `foobar "HELLO WORLD!"`,
			exec: func(t *testing.T, cmd string) {
				proc := StartProc(cmd).Wait()
				if proc.Err() == nil {
					t.Error("expecting command to fail")
				}
				buf := new(bytes.Buffer)
				if _, err := buf.ReadFrom(proc.Out()); err != nil {
					t.Fatal(err)
				}
				if buf.Len() != 0 {
					t.Error("Unexpected output from stdout/stderr")
				}
			},
		},
		{
			name:   "bad command",
			cmdStr: `powershell -Command "Get-Date -InvalidParameter"`,
			exec: func(t *testing.T, cmd string) {
				proc := StartProc(cmd).Wait()
				if proc.Err() == nil {
					t.Error("expecting command to fail")
				}
				result := proc.Result()
				if !strings.Contains(strings.ToLower(result), "parameter") {
					t.Errorf("Expecting error message, got: %s", result)
				}
			},
		},
		{
			name:   "bad command with proc.Out",
			cmdStr: `powershell -Command "Get-Date -InvalidParameter"`,
			exec: func(t *testing.T, cmd string) {
				proc := StartProc(cmd).Wait()
				if proc.Err() == nil {
					t.Error("expecting command to fail")
				}

				buf := new(bytes.Buffer)
				if _, err := buf.ReadFrom(proc.Out()); err != nil {
					t.Fatal(err)
				}

				result := strings.TrimSpace(buf.String())
				if !strings.Contains(strings.ToLower(result), "parameter") {
					t.Errorf("Expecting error message, got: %s", result)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.exec(t, test.cmdStr)
		})
	}
}

func TestRunProc(t *testing.T) {
	tests := []struct {
		name   string
		cmdStr string
		exec   func(*testing.T, string)
	}{
		{
			name:   "access result",
			cmdStr: `powershell -Command "Write-Output 'HELLO WORLD!'"`,
			exec: func(t *testing.T, cmd string) {
				p := RunProc(cmd)
				if p.ExitCode() != 0 {
					t.Fatal("Expecting exit code 0, got:", p.ExitCode())
				}
				if !p.IsSuccess() {
					t.Fatal("Process should be success")
				}
				if p.Result() != "HELLO WORLD!" {
					t.Fatal("Unexpected command result:", p.Result())
				}
			},
		},
		{
			name:   "access proc.Out()",
			cmdStr: `powershell -Command "Write-Output 'HELLO WORLD!'"`,
			exec: func(t *testing.T, cmd string) {
				p := RunProc(cmd)
				if p.ExitCode() != 0 {
					t.Fatal("Expecting exit code 0, got:", p.ExitCode())
				}
				if !p.IsSuccess() {
					t.Fatal("Process should be success")
				}
				result := new(bytes.Buffer)
				if _, err := result.ReadFrom(p.Out()); err != nil {
					t.Fatal(err)
				}

				if strings.TrimSpace(result.String()) != "HELLO WORLD!" {
					t.Fatal("Unexpected command result:", result.String())
				}
			},
		},
		{
			name:   "with expansion",
			cmdStr: `powershell -Command "Write-Output '$MSG'"`,
			exec: func(t *testing.T, cmd string) {
				v := vars.New().Vars("MSG='Hello World'")
				proc := RunProcWithVars(cmd, v)
				if err := proc.Err(); err != nil {
					t.Fatalf("failed to run proc: %s", err)
				}
				if proc.Result() != v.Val("MSG") {
					t.Fatal("Unexpected command result:", proc.Result())
				}
			},
		},
		{
			name:   "missing command with result",
			cmdStr: `foobar "HELLO WORLD!"`,
			exec: func(t *testing.T, cmd string) {
				proc := RunProc(cmd)
				if proc.Err() == nil {
					t.Error("expecting command to fail")
				}
				result := proc.Result()
				if !strings.Contains(result, "not recognized") && !strings.Contains(result, "not found") {
					t.Errorf("Expecting result with error message, got: %s", result)
				}
			},
		},
		{
			name:   "missing command with proc.Out",
			cmdStr: `foobar "HELLO WORLD!"`,
			exec: func(t *testing.T, cmd string) {
				proc := RunProc(cmd)
				if proc.Err() == nil {
					t.Error("expecting command to fail")
				}

				buf := new(bytes.Buffer)
				if _, err := buf.ReadFrom(proc.Out()); err != nil {
					t.Fatal(err)
				}

				result := buf.String()
				if result != "" {
					t.Errorf("Expecting no output, but got %s", result)
				}
			},
		},
		{
			name:   "bad command with result",
			cmdStr: `powershell -Command "Get-Date -InvalidParameter"`,
			exec: func(t *testing.T, cmd string) {
				proc := RunProc(cmd)
				if proc.Err() == nil {
					t.Error("expecting command to fail")
				}
				result := proc.Result()
				if !strings.Contains(strings.ToLower(result), "parameter") {
					t.Errorf("Expecting error message, got: %s", result)
				}
			},
		},
		{
			name:   "bad command option with proc.Out",
			cmdStr: `powershell -Command "Get-Date -InvalidParameter"`,
			exec: func(t *testing.T, cmd string) {
				proc := RunProc(cmd)
				if proc.Err() == nil {
					t.Error("expecting command to fail")
				}

				buf := new(bytes.Buffer)
				if _, err := buf.ReadFrom(proc.Out()); err != nil {
					t.Fatal(err)
				}

				result := buf.String()
				if !strings.Contains(strings.ToLower(result), "parameter") {
					t.Errorf("Expecting error message, got: %s", result)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.exec(t, test.cmdStr)
		})
	}
}

func TestRun(t *testing.T) {
	tests := []struct {
		name   string
		cmdStr string
		exec   func(string)
	}{
		{
			name:   "simple run",
			cmdStr: `powershell -Command "Write-Output 'HELLO WORLD!'"`,
			exec: func(cmd string) {
				result := Run(cmd)
				if result != "HELLO WORLD!" {
					t.Fatal("Unexpected command result:", result)
				}
			},
		},
		{
			name:   "run with expansion",
			cmdStr: `powershell -Command "Write-Output '$MSG'"`,
			exec: func(cmd string) {
				result := RunWithVars(cmd, vars.New().SetVar("MSG", "Hello World"))
				if result != "Hello World" {
					t.Fatal("Unexpected command result:", result)
				}
			},
		},
		{
			name:   "missing command",
			cmdStr: `foobar "HELLO WORLD!"`,
			exec: func(cmd string) {
				result := Run(cmd)
				if !strings.Contains(result, "not recognized") && !strings.Contains(result, "not found") {
					t.Errorf("Expecting result with error message, got: %s", result)
				}
			},
		},
		{
			name:   "bad command",
			cmdStr: `powershell -Command "Get-Command -InvalidParameter"`,
			exec: func(cmd string) {
				result := Run(cmd)
				if !strings.Contains(strings.ToLower(result), "parameter") {
					t.Errorf("Expecting error message, got: %s", result)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.exec(test.cmdStr)
		})
	}
}
