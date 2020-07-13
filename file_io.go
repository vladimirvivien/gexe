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
	if _, err := buf.ReadFrom(file); err != nil {
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

	if _, err := buf.ReadFrom(file); err != nil {
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

// Write creates/truncates an os.File (mode 0666) and writes data into it
func (f *File) Write(data string) {
	file, err := os.Create(f.filename)
	if err != nil {
		f.err = err
		return
	}
	defer file.Close()

	if _, err := file.Write([]byte(data)); err != nil {
		f.err = err
		return
	}
}

// Append appends to an existing os.File (mode 0644) and writes data into it
func (f *File) Append(data string) {
	file, err := os.OpenFile(f.filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		f.err = err
		return
	}
	defer file.Close()
	if _, err := file.Write([]byte(data)); err != nil {
		f.err = err
		return
	}
}

// WriteLines opens an os.File and writes data into it
func (f *File) WriteLines(lines []string) {
	file, err := os.Create(f.filename)
	if err != nil {
		f.err = err
		return
	}
	defer file.Close()

	for _, line := range lines {
		if _, err := file.Write([]byte(line)); err != nil {
			f.err = err
			return
		}
	}
}

// AppendLines appends lines to an existing os.File (mode 0644)
func (f *File) AppendLines(lines []string) {
	file, err := os.OpenFile(f.filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		f.err = err
		return
	}
	defer file.Close()

	for _, line := range lines {
		if _, err := file.Write([]byte(line)); err != nil {
			f.err = err
			return
		}
	}
}

// WriteBytes opens an os.File and writes data into it
func (f *File) WriteBytes(data []byte) {
	file, err := os.Create(f.filename)
	if err != nil {
		f.err = err
		return
	}
	defer file.Close()

	if _, err := file.Write(data); err != nil {
		f.err = err
		return
	}
}

// AppendBytes appends bytes to an existing os.File (mode 0644)
func (f *File) AppendBytes(data []byte) {
	file, err := os.OpenFile(f.filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		f.err = err
		return
	}
	defer file.Close()

	if _, err := file.Write(data); err != nil {
		f.err = err
		return
	}
}

// Writer opens an os.File and returns it as io.WriteCloser
// Ensure to call io.WriterCloser.Close() after use
func (f *File) Writer() io.WriteCloser {
	file, err := os.Open(f.filename)
	if err != nil {
		f.err = err
		return nil
	}
	return file
}

// StreamFrom creates a new file for f and streams bytes from r into it
func (f *File) StreamFrom(r io.Reader) {
	file, err := os.Create(f.filename)
	if err != nil {
		f.err = err
		return
	}
	defer file.Close()

	if _, err := io.Copy(file, r); err != nil {
		f.err = err
		return
	}
}

// StreamTo streams content of file f into w
func (f *File) StreamTo(w io.Writer) {
	file, err := os.Open(f.filename)
	if err != nil {
		f.err = err
		return
	}
	defer file.Close()

	if _, err := io.Copy(w, file); err != nil {
		f.err = err
		return
	}
}

// Err returns any execution error
func (f *File) Err() error {
	return f.err
}
