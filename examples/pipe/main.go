package main

import (
	"fmt"
	"log"

	"github.com/vladimirvivien/gexe"
)

// This example shows how to use gexe to pipe commands.
// The first curl command will download a text by W. E. Du Bois, from Gutenberg,
// and pipe it to the second command which returns the number of lines in that text

func main() {
	pipe := gexe.Pipe(
		"curl https://www.gutenberg.org/cache/epub/15265/pg15265.txt",
		"wc -l",
	)

	if len(pipe.ErrProcs()) > 0 {
		log.Fatalf("failed to download file")
	}
	fmt.Printf("The Quest of the Silver Fleece, by W. E. B. Du Bois: has %s lines\n", pipe.LastProc().Result())
}
