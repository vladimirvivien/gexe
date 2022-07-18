package exec

import (
	"sync"
)

type CommandPolicy byte

const (
	ExitOnErrPolicy CommandPolicy = 1 << iota
	ConcurrentExecPolicy
)

// CommandResult stores results of executed commands using the CommandBuilder
type CommandResult struct {
	mu       sync.RWMutex
	workChan chan *Proc
	procs    []*Proc
	errProcs []*Proc
}

func (cr *CommandResult) Procs() []*Proc {
	cr.mu.RLock()
	defer cr.mu.RUnlock()
	return cr.procs
}
func (cr *CommandResult) ErrProcs() []*Proc {
	cr.mu.RLock()
	defer cr.mu.RUnlock()
	return cr.errProcs
}

type PipedCommandResult struct {
	procs    []*Proc
	errProcs []*Proc
	lastProc *Proc
}

func (cr *PipedCommandResult) Procs() []*Proc {
	return cr.procs
}
func (cr *PipedCommandResult) ErrProcs() []*Proc {
	return cr.errProcs
}
func (cr *PipedCommandResult) LastProc() *Proc {
	procLen := len(cr.procs)
	if procLen == 0 {
		return nil
	}
	return cr.procs[procLen-1]
}

type CommandBuilder struct {
	cmdPolicy CommandPolicy
	procs     []*Proc
}

// Commands creates a *CommandBuilder used to collect
// command strings to be executed.
func Commands(cmds ...string) *CommandBuilder {
	cb := new(CommandBuilder)
	for _, cmd := range cmds {
		cb.procs = append(cb.procs, NewProc(cmd))
	}
	return cb
}

// WithPolicy sets one or more command policy mask values, i.e. (CmdOnErrContinue | CmdExecConcurrent)
func (cb *CommandBuilder) WithPolicy(policyMask CommandPolicy) *CommandBuilder {
	cb.cmdPolicy = policyMask
	return cb
}

// Add adds a new command string to the builder
func (cb *CommandBuilder) Add(cmds ...string) *CommandBuilder {
	for _, cmd := range cmds {
		cb.procs = append(cb.procs, NewProc(cmd))
	}
	return cb
}

// Run executes all commands successively and waits for all of the result. The result of each individual
// command can be accessed from CommandResult.Procs[] after the execution completes. If policy == ExitOnErrPolicy, the
// execution will stop on the first error encountered, otherwise it will continue. Processes with errors can be accessed
// from CommandResult.ErrProcs.
func (cb *CommandBuilder) Run() *CommandResult {
	var result CommandResult
	for _, p := range cb.procs {
		result.procs = append(result.procs, p)
		if err := cb.runCommand(p); err != nil {
			result.errProcs = append(result.errProcs, p)
			if hasPolicy(cb.cmdPolicy, ExitOnErrPolicy) {
				break
			}
			continue
		}
	}

	return &result
}

// Start starts all processes sequentially by default, or concurrently if ConcurrentExecPolicy is set, and does not wait for the commands
// to complete. Use CommandResult.Wait to wait for the processes to complete. Then, the result of each command can be accessed
// from CommandResult.Procs[] or CommandResult.ErrProcs to access failed processses. If policy == ExitOnErrPolicy, the execution will halt
// on the first error encountered, otherwise it will continue.
func (cb *CommandBuilder) Start() *CommandResult {
	result := &CommandResult{workChan: make(chan *Proc, len(cb.procs))}
	go func(builder *CommandBuilder, cr *CommandResult) {
		defer close(cr.workChan)

		// start with concurrently
		if hasPolicy(builder.cmdPolicy, ConcurrentExecPolicy) {
			var gate sync.WaitGroup
			for _, proc := range builder.procs {
				cr.mu.Lock()
				cr.procs = append(cr.procs, proc)
				cr.mu.Unlock()
				proc.cmd.Stdout = proc.result
				proc.cmd.Stderr = proc.result

				gate.Add(1)
				go func(conProc *Proc, conResult *CommandResult) {
					conResult.mu.Lock()
					defer conResult.mu.Unlock()
					defer gate.Done()
					if err := conProc.Start().Err(); err != nil {
						cr.errProcs = append(cr.errProcs, conProc)
						return
					}
					conResult.workChan <- conProc
				}(proc, cr)
			}
			gate.Wait()
			return
		}

		// start sequentially
		for _, proc := range builder.procs {
			cr.mu.Lock()
			cr.procs = append(cr.procs, proc)
			cr.mu.Unlock()
			proc.cmd.Stdout = proc.result
			proc.cmd.Stderr = proc.result

			// start sequentially
			if err := proc.Start().Err(); err != nil {
				cr.mu.Lock()
				cr.errProcs = append(cr.errProcs, proc)
				cr.mu.Unlock()
				if hasPolicy(builder.cmdPolicy, ExitOnErrPolicy) {
					break
				}
				continue
			}

			cr.workChan <- proc
		}
	}(cb, result)

	return result
}

// Concurr starts all processes concurrently and does not wait for the commands
// to complete. It is equivalent to Commands(...).WithPolicy(ConcurrentExecPolicy).Start().
func (cb *CommandBuilder) Concurr() *CommandResult {
	cb.cmdPolicy = ConcurrentExecPolicy
	return cb.Start()
}

// Pipe executes each command serially chaining the combinedOutput of previous command to the inputPipe of next command.
func (cb *CommandBuilder) Pipe() *PipedCommandResult {
	var result PipedCommandResult
	procLen := len(cb.procs)
	if procLen == 0 {
		return nil
	}

	last := procLen - 1
	result.lastProc = cb.procs[last]
	result.lastProc.cmd.Stdout = result.lastProc.result
	result.lastProc.cmd.Stderr = result.lastProc.result

	if procLen > 1 {
		result.lastProc.cmd.Stdout = result.lastProc.result
		// connect input/output between commands
		for i, p := range cb.procs[:last] {
			// link proc.Output to proc[next].Input
			cb.procs[i+1].SetStdin(p.GetOutputPipe())
			p.cmd.Stderr = p.result
		}
	}

	// start each process (but, not wait for result)
	// to ensure data flow between successive processes start
	for _, p := range cb.procs {
		result.procs = append(result.procs, p)
		if err := p.Start().Err(); err != nil {
			result.errProcs = append(result.errProcs, p)
			return &result
		}
	}

	// wait and access processes result
	for _, p := range cb.procs {
		if err := p.Wait().Err(); err != nil {
			result.errProcs = append(result.errProcs, p)
			break
		}
	}

	return &result
}

// Start starts running the registered procs serially and returns immediately.
// This should be followed by a call to Wait to retrieve results.
//func (cb *CommandBuilder) Start() *CommandBuilder {
//	if len(cb.procs) == 0 {
//		return cb
//	}
//
//	cb.procChan = make(chan *Proc, len(cb.procs))
//	switch {
//	case hasPolicy(cb.cmdPolicy, CmdExecConcurrent):
//		// launch each command in its own goroutine
//		go func() {
//			defer close(cb.procChan)
//			var gate sync.WaitGroup
//			for _, proc := range cb.procs {
//				gate.Add(1)
//				go func(wg *sync.WaitGroup, ch chan<- *Proc, p *Proc) {
//					defer wg.Done()
//					ch <- p.Start()
//				}(&gate, cb.procChan, proc)
//			}
//			// wait for procs to launch
//			gate.Wait()
//		}()
//
//	case hasPolicy(cb.cmdPolicy, CmdExecPipe):
//		// execute each proc, then pipe result into successive command
//		go func(ch chan<- *Proc) {
//			defer close(cb.procChan)
//			if len(cb.procs) == 0 {
//				return
//			}
//			if len(cb.procs) == 1 {
//				ch <- cb.procs[0].Start().Wait()
//				return
//			}
//
//			// execute each proc and chain combinedOutput into next inputPipe
//			prev := new(bytes.Buffer)
//
//			var i int
//			for i = 0; i < len(cb.procs); i++ {
//				p := cb.procs[i]
//
//				// execute first proc
//				if i == 0 {
//					if p.Start().Err() != nil && hasPolicy(cb.cmdPolicy, CmdOnErrExit) {
//						break
//					}
//					// save combinedOutput for next inputPipe
//					//if _, err := prev.ReadFrom(p.Out()); err != nil && hasPolicy(cb.cmdPolicy, CmdOnErrExit) {
//					//	break
//					//}
//					if p.Err() != nil && hasPolicy(cb.cmdPolicy, CmdOnErrExit) {
//						break
//					}
//					continue
//				}
//
//				// execute remaining procs
//				if p.Start().Err() != nil && hasPolicy(cb.cmdPolicy, CmdOnErrExit) {
//					break
//				}
//
//				// copy proc[i-1].out -> proc.in
//				go func(cw io.WriteCloser, cr io.Reader) {
//					io.Copy(cw, cr)
//				}(p.Input(), prev)
//
//				p.Wait()
//
//				// save wait for and save combinedOutput for next command
//				if _, err := prev.ReadFrom(p.StdOut()); err != nil && hasPolicy(cb.cmdPolicy, CmdOnErrExit) {
//					break
//				}
//
//				// check for err
//				if p.Err() != nil && hasPolicy(cb.cmdPolicy, CmdOnErrExit) {
//					break
//				}
//			}
//
//			// send proc results
//			for j := 0; j <= i; j++ {
//				ch <- cb.procs[j]
//			}
//		}(cb.procChan)
//	default:
//		// launch all procs (serially), return immediately
//		go func(ch chan<- *Proc) {
//			defer close(cb.procChan)
//			for _, proc := range cb.procs {
//				ch <- proc.Start()
//			}
//		}(cb.procChan)
//	}
//	return cb
//}
//
//func (cb *CommandBuilder) Wait() CommandProcs {
//	if len(cb.procs) == 0 || cb.procChan == nil {
//		return CommandProcs{procs: []*Proc{}}
//	}
//
//	var result CommandProcs
//	for proc := range cb.procChan {
//		result.procs = append(result.procs, proc)
//
//		switch {
//		case hasPolicy(cb.cmdPolicy, CmdExecPipe):
//			// piped cmds already processed, nothing else to do
//		default:
//			// check for start err
//			if proc.Err() != nil {
//				if hasPolicy(cb.cmdPolicy, CmdOnErrExit) {
//					break
//				}
//			}
//
//			// wait for command to complete
//			if err := proc.Wait().Err(); err != nil {
//				if hasPolicy(cb.cmdPolicy, CmdOnErrExit) {
//					break
//				}
//			}
//		}
//
//	}
//	return result
//}

func (cp *CommandBuilder) runCommand(proc *Proc) error {
	// setup combined output for reach proc
	proc.cmd.Stdout = proc.result
	proc.cmd.Stderr = proc.result

	if err := proc.Start().Err(); err != nil {
		return err
	}

	if err := proc.Wait().Err(); err != nil {
		return err
	}
	return nil
}

func (cr *CommandResult) Wait() *CommandResult {
	for proc := range cr.workChan {
		if err := proc.Wait().Err(); err != nil {
			cr.errProcs = append(cr.errProcs, proc)
		}
	}
	return cr
}

func hasPolicy(mask, pol CommandPolicy) bool {
	return (mask & pol) != 0
}
