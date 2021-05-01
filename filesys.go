package echo

import (
	"github.com/vladimirvivien/echo/fs"
)

func (e *Echo) Read(path string) fs.FileReader {
	return fs.Read(e.Eval(path))
}

func (e *Echo) Write(path string) fs.FileWriter {
	return fs.Write(e.Eval(path))
}