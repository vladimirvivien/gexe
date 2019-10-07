package main

import (
	"fmt"

	"github.com/vladimirvivien/echo"
)

func main() {
	e := echo.New()
	e.Var("MSG", "Hello,World,!")
	fmt.Println(e.Eval("I am in $PWD"))
	for _, msg := range e.Arr("$MSG", ",") {
		fmt.Println(msg)
	}
}
