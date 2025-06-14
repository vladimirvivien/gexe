# gexe

The goal of project `gexe` is to make it simple to write code for system operations and task automation using a script-like API that offers the security and type safety of the Go programming language.

## The gexe session
To get access to `gexe`'s functionalities, you must create a session. This is done as follows:

```go
g := gexe.New()
```

Variable `g` starts a `gexe` session (type `gexe.Session`) that provides access to the following:
* Access to all of `gexe` functionalities
* Session variables
* Reported errors

For instance, we can use the previously created session to run an OS command:

```go
g := gexe.New()
fmt.Println(g.Run(`echo "Hello World!"`))
```

The top-level `gexe` package offers another level of convenience by exposing a default session called `gexe.DefaultSession` that is created when the `gexe` package is initialized. So the previous program can be rewritten as:

```go
fmt.Println(gexe.Run(`echo "Hello World!"`))
```

As mentioned earlier, the `gexe` session provides shortcut access to all functionalities implemented by its subpackages. The following shows a short snippet of the functions that you can access using a `gexe` session:

```go
gexe.Run(`echo "Hello World!"`)
gexe.Pipe(`echo "Hello World!"`, "wc -l")
gexe.PathExists("/tmp/myfile")
gexe.Mkdir("/tmp/myapp")
gexe.FileWrite("/tmp/myapp/myfile.txt").String(gexe.Run(`echo "Hello World!"`))
gexe.Get("https://somehost/someresource").String()
// ... etc
```
See all top-level functions in [functions.go](../functions.go).

These functions make use of the functionalities that are implemented in the packages outlined below.

---

## Package exec
Package `exec` can be used to launch and manage external processes by wrapping the `os/exec` types. The package exposes functionalities such as:
* External processes control (launch, kill, etc)
* Retrieve process information
* Launch and manage multiple processes (serially or concurrently)

Package `exec` exposes type `Proc` which provides an entry point for process management. The following example shows how the package works:

```go
p := exec.NewProc(`echo "Hello World!"`)
p.Start()
if p.IsSuccess() {
    fmt.Println("processes launched okay")
}
```

The previous can be shortened using convenient functions found in the `exec` package:

```go
msg := exec.Run(`echo "Hello World!"`)
fmt.Println(msg)
```

The `gexe` package provides functions to access `exec` types to make it easy to launch and manage processes. The previous examples can be done using the `gexe` package:

```go
gexe.SetVar("msg", "World!")
message := gexe.Run(`echo "Hello $msg"`)
fmt.Print(message)
```

### Exec builders
Package `exec` also exposes builders that are designed to launch and manage multiple external processes at once. For instance, the following uses the exec builder to download three files at once by launching three processes concurrently:

```go
cmds := exec.Commands(
    "wget https://text/file1",
    "wget https://text/file2",
    "wget https://text/file3",
).WithPolicy(exec.ConcurrentExecPolicy).Run()

if len(cmds.Errs()) > 0 {
    fmt.Println("There were errors")
}
```

The previous can be simplified as follows:

```go
cmds := exec.Commands(
    "wget https://text/file1",
    "wget https://text/file2",
    "wget https://text/file3",
).Concurr()

if len(cmds.Errs()) > 0 {
    fmt.Println("There were errors")
}
```
As before, functionalities from the `exec` package are exposed via the `gexe` package. The previous example can be achieved using the `gexe` package as follows:

```go
cmds := gexe.RunConcur(
    "wget https://text/file1",
    "wget https://text/file2",
    "wget https://text/file3",
)

if len(cmds.Errs()) > 0 {
    fmt.Println("There were errors")
}
```

---

## Package fs

This package exposes functionalities to create filesystem artifacts such as directories and files. The package makes it easy to create files and add content from different sources including string, bytes, other files, etc.

### Path and directories
Package `fs` exposes type `Path` which provides the entry point for accessing an OS path. A path can point to a directory or an actual file. Once a path is referenced, the API provides additional functionalities to create or access path resources. For instance, the following example uses the `fs` package to create a directory:

```go
dir := fs.Path("/path/to/dir").MkDir(0744)
if dir.Err() == nil {
    fmt.Printf("Created dir %s", dir.Path())
}
```

Or, you can do the same using `gexe` top-level package

```go
dir := gexe.MkDir("/path/to/dir", 0744)
if dir.Err() == nil {
    fmt.Printf("Created dir %s", dir.Path())
}
```

### Reading files
The `fs` package provides access to functionalities to read content of existing files as shown in the following example:

```go
data := fs.Read("/path/to/file")
if data.Err() != nil {
    fmt.Println("Unable to read file:", data.Err())
}
fmt.Println("File content:", data.String())
```

The previous functionality can be accessed from the `gexe` package as shown below:
```go
data := gexe.ReadFile("/path/to/file")
if data.Err() != nil {
    fmt.Println("Unable to read file:", data.Err())
}
fmt.Println("File content:", data.String())
```

### Writing files
The `fs` package also makes it easy to write content, from different sources, to a file as shown in the example below:

```go
get, _ := http.Get("https://test/file1")
fs.Write("path/to/file").From(get.Body())
```

Alternatively, the same functionality can be achieved using the `gexe` package:

```go
get, _ := http.Get("https://test/file1")
gexe.FileWrite("path/to/file").From(get.Body())
```

---

## Package http

This package uses the gexe programming model to make it easy to interact with HTTP servers.

### Reading remote resources
Reading a remote resource using the `http` package can be easily done as shown in the following:

```go
f1 := http.Get("https://test/file1").String()
fmt.Println(f1)
```
Similarly, package `http` makes it easy to post data to a remote HTTP server as shown in the following example:

```go
http.Post("https://test/message").String("Hello")
```

As before, functionalities from the `http` package are exposed as functions in the `gexe` package for easier access as shown below:

```go
gexe.FileWrite("/tmp/eapv2.txt").String(gexe.GetUrl("https://test/file1").String())
```

---

## Package vars

Package `vars` exposes type `Variables` which lets you store named values that can be retrieved programmatically or using string expansion methods.

### Type vars.Variables

Type `vars.Variables` is the main type that is used to manage variables. It exposes several methods including:

* Set environment variables that are accessible to forked processes
* Set session variables that are accessible to a running gexe session
* Evaluate a string with embedded variable expansion values

The following shows how to create a variable and use it

```go
v := vars.New()
v.SetVar("msg", "World!")
v.Eval(`Hello ${msg}!`)
```

A `gexe` session maintains an instance of type `vars.Variables` which is used to store session variable values at runtime as shown below:

```go
gexe.SetVar("msg", "World")
gexe.Run(`echo "Hello ${msg}!"`)
```
### Setting session variables

You can use several methods from type `vars.Variables` to set variables.

**Setting session variables using string literals:**

```go
gexe.Vars(`MSG1="Hello"`, `MSG2=World!`)
gexe.Run(`echo "${MSG1} $MSG2"`)
```
**You can set variables using the `vars.Variables.SetVar` method:**

```go
gexe.SetVar("MSG1","Hello").SetVar("MSG2","World!")
gexe.Run(`echo "${MSG1} $MSG2"`)
```
**Sometimes, you may want to clear a previously set variable:**

```go
gexe.SetVar("MSG1","Hello").SetVar("MSG2","World!")
gexe.Run(`echo "${MSG1} $MSG2"`)
gexe.UnsetVar("MSG1").UnsetVar("MSG2")
```
### Setting environment variables

You can use several methods from the `vars.Variables` type to create environment variables that can be accessed by processes launched by `gexe`.

**Setting environment variables using string literals:**

```go
gexe.Envs(`GOOS="linux"`, `GOARCH=amd64`)
gexe.Run(`go build .`)
```
In the previous snippet, the `go build` command will have access to the environment variables `GOOS` and `GOARCH` at process runtime.

**Setting environment variables using key/value with method `vars.Variables.SetEnv`:**

```go
gexe.SetEnv("GOOS","linux").SetEnv("GOARCH","amd64")
gexe.Run(`go build .`)
```
### Accessing variables

There are a couple of ways you can access your session or environment variables once they are set.

**You can access your variables using variable expansion:**

```go
gexe.SetEnv("GOOS","linux").SetEnv("GOARCH","amd64")
gexe.Run(`echo "Building OS: $GOOS; ARC: $GOARCH"`)
```
**Additionally, you can use `Variables` method `Val` to access variable values directly:**

```go
gexe.SetEnv("GOOS","linux").SetEnv("GOARCH","amd64")
gexe.Run(fmt.Sprintf(`echo "Building OS: %s ARC: %s"`, gexe.Val("GOOS"), gexe.Val("GOARCH")))
```