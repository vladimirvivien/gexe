package fs

import (
	"bufio"
	"bytes"
	"io"
	"os"

	"github.com/vladimirvivien/gexe/vars"
)

type FileReader struct {
	err     error
	path    string
	info    os.FileInfo
	mode    os.FileMode
	vars    *vars.Variables
	content *bytes.Buffer
}

// Read reads the file at path and creates new FileReader.
// Access file content using FileReader methods.
func Read(path string) *FileReader {
	info, err := os.Stat(path)
	if err != nil {
		return &FileReader{err: err, path: path}
	}

	fileData, err := os.ReadFile(path)
	if err != nil {
		return &FileReader{err: err, path: path}
	}

	return &FileReader{
		path:    path,
		info:    info,
		mode:    info.Mode(),
		content: bytes.NewBuffer(fileData),
	}
}

// ReadWithVars creates a new FileReader and sets the reader's session variables
func ReadWithVars(path string, variables *vars.Variables) *FileReader {
	reader := Read(variables.Eval(path))
	reader.vars = variables
	return reader
}

// SetVars sets the FileReader's session variables
func (fr *FileReader) SetVars(variables *vars.Variables) *FileReader {
	fr.vars = variables
	return fr
}

// Err returns an operation error during file read.
func (fr *FileReader) Err() error {
	return fr.err
}

// Info surfaces the os.FileInfo for the associated file
func (fr *FileReader) Info() os.FileInfo {
	return fr.info
}

// String returns the content of the file as a string value
func (fr *FileReader) String() string {
	return fr.content.String()
}

// Lines returns the content of the file as slice of string
func (fr *FileReader) Lines() []string {
	if fr.err != nil {
		return []string{}
	}

	var lines []string
	scnr := bufio.NewScanner(fr.content)

	for scnr.Scan() {
		lines = append(lines, scnr.Text())
	}

	// err should never happen, but capture it anyway
	if scnr.Err() != nil {
		fr.err = scnr.Err()
		return []string{}
	}

	return lines
}

// Bytes returns the content of the file as []byte
func (fr *FileReader) Bytes() []byte {
	if fr.err != nil {
		return []byte{}
	}

	return fr.content.Bytes()
}

// Into reads the content of the file and writes
// it into the specified Writer
func (fr *FileReader) Into(w io.Writer) *FileReader {
	if fr.err != nil {
		return fr
	}

	if _, err := io.Copy(w, fr.content); err != nil {
		fr.err = err
	}
	return fr
}
