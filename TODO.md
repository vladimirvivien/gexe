# TODO
## v0.2.0
* [ ] Map program flags (#20)
* [ ] Piping/chaining OS exec commands (#29)
* [ ] Support for scatter/gather exec commands
* [ ] Support for concurrent exec of os commands 
* [ ] Introduce Windows support (#27)




## Previous releases
#### v0.0.1-alpha.3
* [x] Refactor configuration storage
* [x] Store and expose proc information for child processes

#### v0.0.1-alpha.4
* [x] Add Proc namespace with process-related methods
* `Proc.ID()`
* `Proc.Exited()`
* `Proc.ExitCode()`
* `Proc.IsSuccess()`
* `Proc.SysTime()`
* `Proc.UserTime()`
* [x] Add Prog namespace for methods related to running program
* Exit  - `e.Prog.Exit(code, msgs...string)` (exit prog)
* Name  - `e.Prog.Name()` (prog name)
* Args  - `e.Prog.Args()` (slice of prog arguments)
* Pid   - `e.Prog.Pid()` (proc id)
* Pwd   - `e.Prog.Workdir()`
* Avail - `e.Prog.Avail(path)`
* [x] Doc update

#### v0.0.1-alpha.5
* [x] Namespace all methods
* [x] Add package level shortcut funcs for popular methods

#### v0.0.1-alpha.6
* [ ] ~Default shell: sets a default shell for method Run (`e.Conf.Shell()`)~.
* [x] Enhanced text parsing with parameter expansion escape


## ~Upcoming~
- Ability to pipe result from previous cmd i.e.:
  e.Pipe(e.Run(cmd), e.Run(cmd), e.Run(cmd))
- Item iterator:
  e.ForEach(list interface{}, func(item))
- Files iterator:
  e.ForPath(paths interface{}, func(item)))

## Future
- Windows support?