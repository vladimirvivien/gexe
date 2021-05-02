package fs

import (
	"io"
	"os"
)

type FileReader  interface{
	Err() error
	Info() os.FileInfo
	String() string
	Lines() [] string
	Bytes() [] byte
	WriteTo(io.Writer) FileReader
}

type FileWriter interface {
	Err() error
	Info() os.FileInfo
	String(string) FileWriter
	Lines([]string) FileWriter
	Bytes([]byte) FileWriter
	ReadFrom(io.Reader) FileWriter
}

type FileAppender interface {
	FileWriter
}


