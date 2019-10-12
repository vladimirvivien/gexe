## TODO
## v0.0.1-alpha.2
- Fix nested quoted string parsing (i.e. '"value"' will fail)

## Upcoming
- Exit(statcode int, msg...string)
- Store result of previous command in session (using special chars as keys, i.e. @?)
- Ability to pipe result from previous cmd i.e.:
  e.Pipe(e.Run(cmd), e.Run(cmd), e.Run(cmd))
- Wrap command in a shell
  e.Shell(e.Bash, "command")
  e.Shellout(e.Bash, "command")
- Item iterator:
  e.ForEach(list interface{}, func(item))
- Files iterator:
  e.ForPath(paths interface{}, func(item)))