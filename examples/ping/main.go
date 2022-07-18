package main

import (
	"fmt"
	"os"
	"time"

	"github.com/vladimirvivien/gexe"
)

// This example shows how gexe can be used to launch and stream
// the output of the process as it happens. The following code
// starts a `ping` command, streams the output, displays the result,
// then kill the process after 5 seconds.
func main() {
	execTime := time.Second * 5
	fmt.Println("ping golang.org...")

	p := gexe.NewProc("ping golang.org")
	p.SetStdout(os.Stdout)

	if p.Start().Err() != nil {
		fmt.Println("ping failed to start:", p.Err())
		os.Exit(1)
	}

	<-time.After(execTime)
	if err := p.Kill().Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Pingged golang.org for %s\n", execTime)
}
