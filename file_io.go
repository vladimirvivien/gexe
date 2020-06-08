package echo

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

type File struct {
	filename string
	err      error
}

// OpenFile saves the filename for IO operation
func (e *Echo) OpenFile(filename string) *File {
	return &File{filename: filename}
}

// Read opens an *os.File and reads its content as string
func (f *File) Read() string {
	file, err := os.Open(f.filename)
	if err != nil {
		f.err = err
		return ""
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		f.err = err
		return ""
	}

	return buf.String()
}

// ReadLines opens an *os.File and read its content as []string
func (f *File) ReadLines() []string {
	file, err := os.Open(f.filename)
	if err != nil {
		f.err = err
		return nil
	}
	var lines []string
	scnr := bufio.NewScanner(file)

	for scnr.Scan() {
		lines = append(lines, scnr.Text())
	}

	if scnr.Err() != nil {
		f.err = scnr.Err()
		return nil
	}

	return lines
}

// ReadBytes opens an *os.File and read its content as []byte
func (f *File) ReadBytes() []byte {
	file, err := os.Open(f.filename)
	if err != nil {
		f.err = err
		return nil
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		f.err = err
		return nil
	}

	return buf.Bytes()
}

// Reader opens an *os.File and returns it as io.ReadCloser
// Ensure to call io.ReadCloser.Close() after use.
func (f *File) Reader() io.ReadCloser {
	file, err := os.Open(f.filename)
	if err != nil {
		f.err = err
		return nil
	}
	return file
}

// Write opens an os.File and writes data into it
func (f *File) Write(data string) {

}

// WriteLines opens an os.File and writes data into it
func (f *File) WriteLines(data []string){

}

// WriteBytes opens an os.File and writes data into it
func (f *File) WriteBytes(data []byte){

}

// Writer opens an os.File and returns as io.WriteCloser
// Ensure to call io.WriterCloser.Close() after use
func (f *File) Writer () io.WriteCloser {
	return nil
}

// Err returns any execution error
func (f *File) Err() error {
	return f.err
}
