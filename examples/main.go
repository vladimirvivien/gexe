package main

import (
	"fmt"

	"github.com/vladimirvivien/echo"
)

var (
	e = echo.New()
)

func init() {
	e.Conf.SetPanicOnErr(false)
}

func main() {
	e.Var("MYHOME=$HOME")
	if e.IsExist("$MYHOME") {
		fmt.Println(e.Eval("I found my $HOME"))
	}
}
