package echo

import (
	"github.com/vladimirvivien/echo/prog"
)

func (e *Echo) Prog() *prog.ProgInfo {
	return e.prog
}

