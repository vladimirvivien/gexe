package main

import (
	"fmt"
	"os"

	"github.com/vladimirvivien/gexe"
)

// This example shows how to use the fs package to create/write
// files with simplicity.
func main() {
	fmt.Println("Download and saving text to /tmp/thenegro.txt")

	// the following downloads a large W. E. Du Bois text from Gutenburg and writes locally.
	cmd := "wget -O /tmp/thenegro.txt https://www.gutenberg.org/cache/epub/15359/pg15359.txt"

	if w := gexe.FileWrite("/tmp/thenegro.txt").From(gexe.RunProc(cmd).Out()); w.Err() != nil {
		fmt.Println(w.Err())
		os.Exit(1)
	}

	if err := os.Remove("/tmp/thenegro.txt"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
