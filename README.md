[![Go Reference](https://pkg.go.dev/badge/github.com/vladimirvivien/gexe.svg)](https://pkg.go.dev/github.com/vladimirvivien/gexe)
[![Go Report Card](https://goreportcard.com/badge/github.com/vladimirvivien/gexe)](https://goreportcard.com/report/github.com/vladimirvivien/gexe)
![Build](https://github.com/vladimirvivien/gexe/actions/workflows/build.yml/badge.svg)
# Project `gexe`
Package with script-like API for system operation and automation!

The goal of project `gexe` is to make it simple to write code for system operation and task automation using a script-like API that offers the security and the type safety of the Go programming language (see [/examples](/examples/)).

## What can you do with `gexe`?
* Parse and execute OS plain text commands, as you would in a shell.
* Support for variable expansion in command string (i.e. `gexe.Run("echo $HOME")`)
* Support for Go's `fmt.Sprintf` formatting in all string parameters (i.e. `gexe.Run("echo %s", "Hello")`)
* Ability to pipe processes: `gexe.Pipe("cat /etc/hosts", "wc -l")`
* Run processes concurrently: `gexe.RunConcur('wget https://example.com/files'; "date")`
* Get process information (i.e. PID, status, exit code, etc)
* Get program information (i.e. args, binary name, working dir, etc)
* Easily read and write file content using different sources (string, bytes, io.Writer, etc)
* Integrate with your shell script using `go run`

## Using `gexe`

### Get the package
```bash
go get github.com/vladimirvivien/gexe
```

### Run a process
The following executes command `echo "Hello World!"` and prints the result:
```go
fmt.Println(gexe.Run(`echo "Hello World!"`))
```

Alternatively, you can create your own `gexe` session for more control and error hanlding:

```go
g := gexe.New()
proc := g.RunProc(`echo "Hello World"`)
if proc.Err() != nil {
    fmt.Println(proc.Err())
    os.Exit(proc.ExitCode())    
}
fmt.Println(proc.Result())
```

### String formatting with Go's fmt syntax
`gexe` methods now support Go's `fmt.Sprintf` formatting alongside variable expansion:

```go
// Using Go formatting with variable expansion
gexe.SetVar("name", "Alice")
gexe.Run("echo Hello %s, your home is ${HOME}", "World")

// File operations with formatting
gexe.FileWrite("/tmp/log_%s.txt", time.Now().Format("2006-01-02"))

// Variable setting with formatting
gexe.SetVar("message", "User %s logged in at %s", username, timestamp)
```

The formatting is applied intelligently - if no format verbs are detected in the string, the arguments are ignored, maintaining backward compatibility.

## Examples
Find more examples [here](./examples/)!

### Building project `$gexe` with `gexe`
This example shows how `gexe` can be used to build Go project binaries for multiple
platforms and OSes. Note the followings:
* The command string is naturally expressed as you would in a shell.
* The use of variable expansion in the commands.

```go
func main() {
	for _, arch := range []string{"amd64"} {
		for _, opsys := range []string{"darwin", "linux"} {
			gexe.SetVar("arch", arch).SetVar("os", opsys)
			gexe.SetVar("binpath", fmt.Sprintf("build/%s/%s/mybinary", arch, opsys))
			result := gexe.Envs("CGO_ENABLED=0 GOOS=$os GOARCH=$arch").Run("go build -o $binpath .")
			if result != "" {
				fmt.Printf("Build for %s/%s failed: %s\n", arch, opsys, result)
				os.Exit(1)
			}
			fmt.Printf("Build %s/%s: %s OK\n", arch, opsys, gexe.Eval("$binpath"))
		}
	}
}
```
> See [./examples/build/main.go](./examples/build/main.go)

### Long-running process
This example shows how `gexe` can be used to launch a long-running process and stream
its output. The code invokes the `ping` command, streams its output, displays the result,
and then kills the process after 5 seconds.

```go
func main() {
	execTime := time.Second * 5
	fmt.Println("ping golang.org...")

	p := gexe.StartProc("ping golang.org")

	if p.Err() != nil {
		fmt.Println("ping failed:", p.Err())
		os.Exit(1)
	}

	go func() {
		if _, err := io.Copy(os.Stdout, p.StdOut()); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	<-time.After(execTime)
	p.Kill()
	fmt.Printf("Pinged golang.org for %s\n", execTime)
}
```

### Using a shell
This example uses the `git` command to print logs and commit info by using `/bin/sh` to start a shell for command piping:

```go
func main() {
	cmd := `/bin/sh -c "git log --reverse --abbrev-commit --pretty=oneline | cut -d ' ' -f1"`
	for _, p := range strings.Split(gexe.Run(cmd), "\n") {
		gexe.SetVar("patch", p)
		cmd := `/bin/sh -c "git show --abbrev-commit -s --pretty=format:'%h %s (%an) %n' ${patch}"`
		fmt.Println(gexe.Run(cmd))
	}
}
```

# Project Name Change
Originally this project was named `echo`.  However, another Go project by that name has gotten really popular.
So this project was renamed `gexe` (pronounced Jesse).
# License
MIT