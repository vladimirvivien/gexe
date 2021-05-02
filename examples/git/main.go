package main

import (
	"fmt"
	"strings"

	"github.com/vladimirvivien/gexe"
)

// This example uses local git to print logs and commit info.
// Notice the use of /bin/sh to start a shell for more complex
// commands (such as piping).
func main() {
	cmd := `/bin/sh -c "git log --reverse --abbrev-commit --pretty=oneline | cut -d ' ' -f1"`
	for _, p := range strings.Split(gexe.Run(cmd), "\n") {
		gexe.SetVar("patch", p)
		cmd := `/bin/sh -c "git show --abbrev-commit -s --pretty=format:'%h %s (%an) %n' ${patch}"`
		fmt.Println(gexe.Run(cmd))
	}
}
