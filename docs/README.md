# `gexe` API Reference

## Package `gexe`
The `gexe` top-level package provides functions that wraps access to other top-level packages via a package instance of type `Echo` to make access API functionalities easier. Package `gexe` exposes an instance of `Echo` called `DefaultEcho` that is used to provide access to an `Echo` values and methods.

### `gexe.Echo`
The `Echo` type provides access to a session that exposes all of the functionalities of the API. Echo tracks the followings:
* Error
* Session variables
* Program information

You can create a new Echo session using the `gexe.New()` method:

```go
e := gexe.New()
e.Run(`echo "Hello World!"`)
```

When using the `gexe` package instance of `Echo`, the previous example can be written as follows:

```go
gexe.Run(`echo "Hello World!"`)
```
## Package `exec`
Package `exec` can be used to launch and manage external processes by wraping the `os/exec` types.  The package exposes functionalities such as the folliwngs:
* Lauhch / kill externaal processes
* Retrieve process information
* Launch and manage multiple processes

Package `exec` exposes type `Proc` which provides an entry point for process management. The following example shows how the package works:

```go
p := exec.NewProc(`echo "Hello World!"`)
p.Start()
if p.IsSuccess() {
    fmt.Println("processess launched okay")
}
```

The previous can be shortened using convenient functions found in the `exec` package:

```go
msg := exec.Run(`echo "Hello World!"`)
fmt.Println(msg)
```

The `gexe` package provides functions to access `exec` types to make it easy to launch and mange processes. The previous examples can be done using the `gexe` package:

```go
gexe.SetVar("msg", "World!")
message = gexe.Run(`echo "Hello $msg"`)
fmt.Print(message)
```

### Exec builders
Package `exec` also exposes builders that are designed to launch and manage multiple external processes at once. For instance, the following uses the exec builder to download three files at once by launching three processes concurrently:

```go
cmds := exec.Commands(
    "wget https://text/file1",
    "wget https://text/file2",
    "wget https://text/file3",
).WithPolicy(exec.ConcerrentExecPolicy).Run()

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
As before, functionalities from the `exec` package is exposed via the `gexe` package. The previous example can be achieved using the `gexe` package as follows:

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

## Package `fs`
This package exposes functionalities to create filesystem artifacts such as directories and files. The package makes it easy to create files and add content from different sources including string, bytes, other files, etc.

### Path and directories
Package `fs` exposes type `Path` which provides the entry point for access an OS path. A path can point to a directory or an actual file. Once a path is referenced, the API provides additional functionalities to create or access path resources.  For instance, the following examples uses the `fs` package to create a directory:

```go
dir := fs.Path("/path/to/dir").MkDir(0744)
if dir.Err() == nil {
    fmt.Printf("Created dir %s", dir.Path())
}
```

Or, you can do the the same using `gexe` top-level package

```go
dir := gexe.MkDir("/path/to/dir", 0744)
if dir.Err() == nil {
    fmt.Printf("Created dir %s", dir.Path())
}
```

### Reading files
The `fs` package provides access to functionalities to read content of existing files as shown in the follwoing example :

```go
data := fs.Read("/path/to/file")
if data.Err() != nil{
    fmt.Println("Unable to read file:", data.Err())
}
fmt.Println("File content:", data.String())
```

The previous functionality can be accessed from the `gexe` package as shown below:
```go
data := gexe.ReadFile("/path/to/file")
if data.Err() != nil{
    fmt.Println("Unable to read file:", data.Err())
}
fmt.Println("File content:", data.String())
```

### Writing files
The `fs` package also makes it eash to write content, from different sources, to a file as shown in the example below:

```go
get, _ := http.Get("https://test/file1")
fs.Write("path/to/file").From(get.Body())
```

Alternatively, the same functionality can be achieved using the `gexe` package:

```go
get, _ := http.Get("https://test/file1")
gexe.FileWrite("path/to/file").From(get.Body())
```

## Package `http`
This package uses the Gexe programming model to make it easy to programmatically interact with HTTP servers.

### Reading remote resources
Reading a remote resource using the `http` package can be easily done as shown the following:

```go
f1 := http.Get("https://test/file1").String()
fmt.Println(f1)
```
Similarly, package `http` makes it easy to post data to a remote HTTP server as shown in the following example:

```go
http.Post("https://test/message").String("Hello)
```

As before, functionalities from the `http` package are exposed as functions in the `gexe` package for easier access as shown below:

```go
gexe.FileWrite("/tmp/eapv2.txt").String(gexe.GetUrl("https://test/file1").String())
```

## Package `vars`
Package `vars` exposes type `Variables` which lets you store named values that can be retrieved programmatically or using string expansion methods.

### `vars.Variables`
Type `Variables` is the main type that is used to manage variables.  It exposes several methods including
* Set environment variables that are accessible to forked processes
* Set session variables that are accessible to a running gexe session
* Evaluate a string with embedded variable expansion values

The following shows how to create a variable and use it

```go
v := vars.New()
v.SetVar("msg", "World!")
v.Eval(`Hello ${msg}!)
```

The `gexe.Echo` type maintains an instance of type `vars.Variables` which are used to store values for a gexe session as shown in the following exeample:

```go
gexe.SetVar("msg", "World")
gexe.Run(`echo "Hello ${msg}!"`)
```
