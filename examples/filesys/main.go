package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/vladimirvivien/gexe"
	"github.com/vladimirvivien/gexe/str"
)

// This example uses local git to create a file with commit logs.
func main() {
	buf := new(bytes.Buffer)
	cmd := `/bin/sh -c "git log --reverse --abbrev-commit --pretty=oneline | cut -d ' ' -f1"`
	for _, p := range str.SplitLines(gexe.Run(cmd)) {
		gexe.SetVar("patch", p)
		cmd := `/bin/sh -c "git show --abbrev-commit -s --pretty=format:'%h %s (%an) %n' ${patch}"`
		buf.WriteString(fmt.Sprintf("%s\n", gexe.Run(cmd)))
	}

	gitfile := "./gitlog.txt"

	if w := gexe.Write(gitfile).From(buf); w.Err() != nil {
		fmt.Println(w.Err())
		os.Exit(1)
	}

	// read the file and print
	fmt.Println(gexe.Read(gitfile).String())

	if err := os.Remove(gitfile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
