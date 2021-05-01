# `echo` 
OS interaction wrapped in the security and type safety of the Go programming language!

The goal of `echo` is to make it easy to write code that interacts with the OS (or other infrastructure components) 
using the security and type safety of the Go programming language. 

## What `echo` is
* Not a tool for shell-scripting Go code
* Designed to be used as idiomatic Go
* Rich types for easy interactions with OS (exec, IO, etc)
* Programs can be used as pre-compiled binaries or with `go run`
   

## Using `echo`
The `echo` package comes with several functions ready to be used. For instance
to execute an external command, you can use the following:"
```
echo.Run(`echo "Hello World!"`)
```

Alternatively, you can create your own echo session with:
```
e := echo.New()
```
Then run the echo methods available:
```
e.Run(...)
```

### Building Go with `echo`
This example shows how `echo` can be used to build Go project binaries for multiple
platforms and OSes.
```go
func main() {
	for _, arch := range []string{"amd64"} {
		for _, opsys := range []string{"darwin", "linux"} {
			echo.SetVar("arch", arch).SetVar("os", opsys)
			echo.SetVar("binpath", fmt.Sprintf("build/%s/%s/mybinary", arch, opsys))
			result := echo.Envs("CGO_ENABLED=0 GOOS=$os GOARCH=$arch").Run("go build -o $binpath .")
			if result != "" {
				fmt.Printf("Build for %s/%s failed: %s\n", arch, opsys, result)
				os.Exit(1)
			}
			fmt.Printf("Build %s/%s: %s OK\n", arch, opsys, echo.Eval("$binpath"))
		}
	}
}
```
> See [./examples/build/main.go](./examples/build/main.go)

### Example of a long running process
The following shows how `echo` can be used to launch a long running process and stream 
its output. The code invokese the `ping` command, streams its output, displays the result,
and then kills the process after 5 seconds.

```go
func main() {
	execTime := time.Second * 5
	fmt.Println("ping golang.org...")

	p := echo.StartProc("ping golang.org")

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
	fmt.Printf("Pingged golang.org for %s\n", execTime)
}
```

### Example using shell 
This example uses git to print logs and commit info.
The code invokes git as a sub-command of `/bin/sh` to start a shell 
for more complex functionalities (such as piping).
```go
func main() {
	e := echo.New()
	cmd := `/bin/sh -c "git log --reverse --abbrev-commit --pretty=oneline | cut -d ' ' -f1"`
	for _, p := range strings.Split(e.Run(cmd), "\n") {
		e.SetVar("patch", p)
		cmd := `/bin/sh -c "git show --abbrev-commit -s --pretty=format:'%h %s (%an) %n' ${patch}"`
		fmt.Println(e.Run(cmd))
	}
}
```
## License
MIT