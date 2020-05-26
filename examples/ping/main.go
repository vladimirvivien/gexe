package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/vladimirvivien/echo"
)

func main() {
	execTime := time.Second * 5
	fmt.Println("ping golang.org...")

	e := echo.New()
	p := e.StartProc("ping golang.org")

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
	p.Kill()
	p.Wait()
	fmt.Printf("Pingged golang.org for %s\n", execTime)
}
