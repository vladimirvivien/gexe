package main

import (
	"fmt"
	"os"

	"github.com/vladimirvivien/gexe"
)

// This examples shows simple how to store value in
// local or environment variables used in expansion
// at runtime.
func main() {
	e := gexe.New()
	e.Vars("MYUSERNAME=$USER")
	if e.Eval("$MYUSERNAME") == "" {
		fmt.Println("You were not found")
		os.Exit(1)
	}

	fmt.Printf("Hello %s !!! \n", e.Eval("$MYUSERNAME"))
}
