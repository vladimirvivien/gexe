package main

import (
	"fmt"
	"log"
	"os"

	"github.com/vladimirvivien/gexe"
)

// This example shows how to execute OS commands concurrently and wait for the result.
func main() {

	fmt.Println("Downloading 3 books concurrently from Gutenberg")

	// The following launches each command concurrently to download long text from Gutenberg.
	// The download should be faster than if ran sequentially.
	result := gexe.RunConcur(
		"wget -O /tmp/thenegro.txt https://www.gutenberg.org/cache/epub/15359/pg15359.txt",
		"wget -O /tmp/fleece.txt https://www.gutenberg.org/cache/epub/15265/pg15265.txt",
		"wget -O /tmp/conversation.txt https://www.gutenberg.org/cache/epub/31254/pg31254.txt",
	)

	// inspect result or check for errors.
	if len(result.ErrProcs()) > 0 {
		log.Println("One or more commands failed")
	}

	for _, path := range []string{"/tmp/thenegro.txt", "/tmp/fleece.txt", "/tmp/conversation.txt"} {
		if _, err := os.Stat(path); err == nil {
			fmt.Printf("file %s downloaded OK\n", path)
			os.RemoveAll(path)
		}
	}

}
