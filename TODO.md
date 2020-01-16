## TODO

## v0.0.1-alpha.3
* [ ] Refactor configuration storage
* [ ] Store and expose proc information for child processes

## v0.0.1-alpha.4
* [ ] Default shell: sets a default shell for method Run (`e.Conf.Shell()`).
* [ ] Enhanced text parsing with parameter expansion escape

## v0.0.1-alpha.5
* [ ] Exit - `e.Exit(code, msgs...string)` (exit prog)
* [ ] Name - `e.Name()` (prog name)
* [ ] Args - `e.Args()` (slice of prog arguments)
* [ ] Pid  - `e.Pid()` (proc id)

## v0.0.1-alpha.6
* [ ] Hostname  - `e.Hostname()`
* [ ] Pwd       - `e.Pwd()`
* [ ] Which     - `e.Which(path)`

## Upcoming
- Ability to pipe result from previous cmd i.e.:
  e.Pipe(e.Run(cmd), e.Run(cmd), e.Run(cmd))
- Item iterator:
  e.ForEach(list interface{}, func(item))
- Files iterator:
  e.ForPath(paths interface{}, func(item)))

## Future
- Windows support?