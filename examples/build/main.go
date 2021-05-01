package main

import (
	"fmt"
	"os"

	"github.com/vladimirvivien/echo"
)

// This example shows how echo can be used in a CI
// pipeline to build Go project binaries for multiple
// platforms and OSes.
func main() {
	for _, arch := range []string{"amd64"} {
		for _, opsys := range []string{"darwin", "linux"} {
			echo.SetVar("arch", arch).SetVar("os", opsys)
			echo.SetVar("binpath", fmt.Sprintf("build/%s/%s/mybinary", arch, opsys))
			result := echo.Envs("CGO_ENABLED=0 GOOS=$os GOARCH=$arch").Run("go build -o $binpath .")
			if result != "" {
				fmt.Printf("Build for %s/%s failed: %s\n", arch, opsys, result)
				os.Exit(1)
			}
			fmt.Printf("Build %s/%s: %s OK\n", arch, opsys, echo.Eval("$binpath"))
		}
	}
}
