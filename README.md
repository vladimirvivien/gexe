# Echo (v0.0.1-alpha.0)
An abstraction layer for the Go exec package to script shell-like.

## Start
Create a session:
```
e := echo.New()
```

## Variables
Local variables used in current `echo` session can be declared using `e.Var()`. Values can be made visible to launched external process using `e.Export()`. Both methods support value expansion as shown below:
```
e.Var("Foo", "Bar")
e.Var("Fuzz", "${Foo}")
e.Export("Bazz", "$Fuzz") 
```

## Array
A space-separated list can be turned into a native Go `[]string` using the `e.Arr()` method as shown:
```
for _, val := range e.Arr("val0 val1 val2 val3"){  
    ...
}
```

An additional separator string may be provided to `e.Arr()`:
```
for _, val := range e.Arr("Hello,World,!", ","){  
    fmt.Println(val)
}
```
Method `e.Glob()` expands the provided shell file pattern into a slice of matching file names:
```
for _, f := range e.Glob("$HOME/go/src/*.com") {
  fmt.Println(f)
}
```

## Run
External commands can be launched using the `e.Run()` method by passing a command string with its arguments as would be done in a shell.  The `Run` method supports string expansion as shown below:
```
e.Var("salutation", "Hello")
e.Var("name", "Vladimir")
e.Var("MSG", "$salutation $name")
result := e.Run("echo $MSG")
```

If the result must be printed to `stdout` immediately, use `e.Runout(...)` instead:
```
e.Runout("We are all done here!")
```

### Other Run-related Methods
The followings returns Process information after each run:
```
e.RExitCode() int
e.RStatus() 
e.RExited() bool
e.RPid() int
e.ROkay() bool
e.RSystime() string
e.RUsertime() string
```

## Strings
```
e.Empty(string) bool
e.Lower(string) string
e.Upper(string) string
e.Streq(string, string) bool // string equal
e.Trim(string)string 
```

## Files
File operation methods:

```
e.Abs()
e.Rel()
e.Base()
e.Dir()
e.PathSym()
e.PathExt()
e.PathJoin()
e.PathMatched()
e.IsPathExit()
e.IsPathReg()
e.IsPathDir()
e.Mkdir()
e.Chown()
e.Chmod()
e.AreSame() // Are files equal
```

## Program 
```
e.Name()
e.Args()
e.Exit(code)
e.Hostname()
e.Grps()[]int
e.Pid()
e.PPid()
e.Pwd()
e.Which(path)
```

## User

```
e.Username()
e.Home()
e.Gid()
e.Uid()
```