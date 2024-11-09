package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/vladimirvivien/gexe"
)

// This examples demonstrates how to use gexe to run
// subprocesses with different userid/groupid.
// Note: setting userid/gid for subprocesses requires
// elevated privilege such as using sudo.

func main() {
	p := gexe.NewProc(`echo "Hello World!`)
	var uid string

	switch runtime.GOOS {
	case "windows":
	case "darwin":
		uid = gexe.Run(`id -u`)
	case "linux":
		uid = gexe.Run(`id -l`)
	}

	if uid != "" {
		if err := p.SetUserid(uid).Err(); err != nil {
			fmt.Println("Failed to set userid: ", err)
			os.Exit(1)
		}
	}

	result := p.Run()
	if err := result.Err(); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println("Running process with userid:", uid)
	fmt.Println(result.Result())
}
