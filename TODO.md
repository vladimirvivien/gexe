## TODO
## v0.0.1-alpha.2
- Fix nested quoted string parsing (i.e. '"value"' will fail)

## v0.0.1-alpha.3
* [ ] Enhance quotation support for Run


## v0.0.1-alpha.4
* [ ] Default shell: sets a default shell for method Run (`e.Conf.Shell()`).
* [ ] Store process information for child process
* [ ] Store process information for Echo

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