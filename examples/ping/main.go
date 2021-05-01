package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/vladimirvivien/echo"
)

// This example shows how echo can be used to launch and stream
// the output of the process as it happens. The following code
// starts a `ping` command, streams the output, displays the result,
// then kill the process after 5 seconds.
func main() {
	execTime := time.Second * 5
	fmt.Println("ping golang.org...")

	p := echo.StartProc("ping golang.org")

	if p.Err() != nil {
		fmt.Println("ping failed:", p.Err())
		os.Exit(1)
	}

	go func() {
		if _, err := io.Copy(os.Stdout, p.StdOut()); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	<-time.After(execTime)
	if err := p.Kill().Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Pingged golang.org for %s\n", execTime)
}
