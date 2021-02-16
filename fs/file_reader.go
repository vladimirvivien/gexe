package fs

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

type fileReader struct {
	err error
	path string
	finfo os.FileInfo
}

// Read creates a file reader; it stat(path)
// to ensure it already exists, if not FileReader.err is filled.
func Read(path string) FileReader {
	fr := &fileReader{path:path}
	info, err := os.Stat(fr.path)
	if err != nil {
		fr.err = err
		return fr
	}
	fr.finfo = info
	return fr
}

func (fr *fileReader) Err() error {
	return fr.err
}

func (fr *fileReader) Info() os.FileInfo {
	return fr.finfo
}


func (fr *fileReader) String() string {
	file, err := os.Open(fr.path)
	if err != nil {
		fr.err = err
		return ""
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(file); err != nil {
		fr.err = err
		return ""
	}

	return buf.String()
}

func (fr *fileReader) Lines() []string {
	file, err := os.Open(fr.path)
	if err != nil {
		fr.err = err
		return []string{}
	}
	var lines []string
	scnr := bufio.NewScanner(file)

	for scnr.Scan() {
		lines = append(lines, scnr.Text())
	}

	if scnr.Err() != nil {
		fr.err = scnr.Err()
		return []string{}
	}

	return lines
}

func (fr *fileReader) Bytes() []byte {
	file, err := os.Open(fr.path)
	if err != nil {
		fr.err = err
		return []byte{}
	}
	defer file.Close()

	buf := new(bytes.Buffer)

	if _, err := buf.ReadFrom(file); err != nil {
		fr.err = err
		return []byte{}
	}

	return buf.Bytes()
}

func (fr *fileReader) WriteTo(w io.Writer) FileReader {
	file, err := os.Open(fr.path)
	if err != nil {
		fr.err = err
		return fr
	}
	defer file.Close()
	if _, err := io.Copy(w, file); err != nil {
		fr.err = err
	}
	return fr
}