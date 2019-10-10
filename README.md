# `echo` (v0.0.1-alpha.0)
A Go pacakge with script-like functionalities!

## Start
Create a session:
```
e := echo.New()
```
Then configure your session:
```
e.Conf.SetPanicOnError(true)
```

## Variables

### Var, SetVar
Echo supports storing values that can be accessed using method `e.Var()` for the duration of a session:
```
e.Var("Foo=Bar")
e.Var("Fuzz=${Foo} Buzz=Bazz") 
```
Method `e.SetVar(name, value string)` saves a named value one at a time.

### Env, SetEnv
Values can be made visible as environment variables for externally launced commands using method `e.Env()`. Both methods support value expansion as shown below:
```
e.Env("Foo=Bar")
e.Env("Fuzz=${Foo} Buzz=Bazz")
e.Env("BAZZ=$HOME")
```
Method `e.SetEnv(name, value string)` sets environment variables one value at a time.

### Expansion
All `echo` methods support variable value expansion using `$name` or `${name}` which are automatically replaced with the value of the named variable.

## Slices

### Split
A space-separated list can be turned into a native Go `[]string` using the `e.Split()` method as shown:
```
e.SetVar("list", "item3 item4")
for _, val := range e.Split("item0 item1 item2 $list"){  
    ...
}
```

An additional separator value may be provided to `e.Split()`:
```
e.SetVar("list", "item3;item4")
for _, val := range e.Arr("Hello;World;!;$list", ";"){  
    fmt.Println(val)
}
```

### Glob
Method `e.Glob()` expands the provided shell path pattern into a slice of matching file/directory names:
```
for _, f := range e.Glob("$HOME/go/src/*.com") {
  fmt.Println(f)
}
```

## External processes

### Run
External commands can be launched using the `e.Run()` method by passing a command and its arguments as a single string:
```
fmt.Println(e.Var("greeting=Hello name=Vladimir").Run("echo $greeting $name"))
```

Method `e.Runout` automatically prints the result to sdtdout for you:
```
e.Runout("echo no place like $HOME")
```

### Process run-related methods
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
e.Ext()
e.PathJoin()
e.PathMatched()
e.IsExit()
e.IsReg()
e.IsDir()
e.Mkdirs()
e.Rmdirs()
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

## Flagset
...

## More to come

## License
MIT