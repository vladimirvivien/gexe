package main

import (
	"fmt"
	"os"

	"github.com/vladimirvivien/gexe"
)

func main() {
	fmt.Println("Download and saving text")

	// The following downloads a large text and saves the result
	url := "https://www.gutenberg.org/cache/epub/2148/pg2148.txt"

	if w := gexe.FileWrite("/tmp/eapv2.txt").From(gexe.GetUrl(url).Body()); w.Err() != nil {
		fmt.Println(w.Err())
		os.Exit(1)
	}

	if err := os.Remove("/tmp/eapv2.txt"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
