package echo

import (
	"os"
	"path/filepath"
)

// Abs returns the absolute representation for path.
// Returns empty string if it fails for any reason.
func (e *Echo) Abs(path string) string {
	abs, err := filepath.Abs(os.Expand(path, e.Val))
	if err != nil {
		e.shouldPanic(err.Error())
		return ""
	}
	return abs
}

// Rel returns path that is relative to base/target.
// Returns an empty path if error.
func (e *Echo) Rel(base string, target string) string {
	abs, err := filepath.Rel(os.Expand(base, e.Val), os.Expand(target, e.Val))
	if err != nil {
		e.shouldPanic(err.Error())
		return ""
	}
	return abs
}

// Base returns the last portion (or name) of a path.
func (e *Echo) Base(path string) string {
	return filepath.Base(os.Expand(path, e.Val))
}

// Dir returns parent path portion of a path (without the Base).
func (e *Echo) Dir(path string) string {
	return filepath.Dir(os.Expand(path, e.Val))
}

// PathSplit splits an OS-specific list of path into []string
func (e *Echo) PathSplit(path string) []string {
	return filepath.SplitList(os.Expand(path, e.Val))
}

// PathSym returns the evaluated symbolic link for path.
// Returns empty if symbolic evaluation fails.
func (e *Echo) PathSym(path string) string {
	link, err := filepath.EvalSymlinks(os.Expand(path, e.Val))
	if err != nil {
		e.shouldPanic(err.Error())
		return ""
	}
	return link
}

// Ext returns extension part of path
func (e *Echo) Ext(path string) string {
	return filepath.Ext(os.Expand(path, e.Val))
}

// PathJoin collate individual paths together for a longer path
func (e *Echo) PathJoin(paths ...string) string {
	for i, path := range paths {
		paths[i] = os.Expand(path, e.Val)
	}
	return filepath.Join(paths...)
}

// PathMatched returns true if path matches shell file pattern
func (e *Echo) PathMatched(pattern, path string) bool {
	matched, err := filepath.Match(pattern, os.Expand(path, e.Val))
	if err != nil {
		return false
	}
	return matched
}

// IsAbs returns true if path is an absolute path
func (e *Echo) IsAbs(path string) bool {
	return filepath.IsAbs(os.Expand(path, e.Val))
}

// IsExist returns true if file is 1) accessible 2) exists.
// All other conditions returns false.
func (e *Echo) IsExist(path string) bool {
	if _, err := os.Stat(os.Expand(path, e.Val)); err != nil {
		if !os.IsNotExist(err) {
			e.shouldPanic(err.Error())
		}
		return false
	}
	return true
}

// IsReg returns true if path is a regular file.
// All other cases (errors) will return false.
func (e *Echo) IsReg(path string) bool {
	info, err := os.Stat(os.Expand(path, e.Val))
	if err != nil {
		return false
	}

	return info.Mode().IsRegular()
}

// IsDir returns true if path is a directory.
// All other cases (errors) will return false.
func (e *Echo) IsDir(path string) bool {
	info, err := os.Stat(os.Expand(path, e.Val))
	if err != nil {
		return false
	}
	return info.Mode().IsDir()
}

// Mkdirs creates one or more space-separated directories using the optional filemods.
// If no mode is provided, it defaults to 0777. When mode is provided only
// modes[0] is applied.
func (e *Echo) Mkdirs(paths string, modes ...uint32) {
	for _, path := range e.Split(paths) {
		var mode uint32
		switch {
		case len(modes) >= 1:
			mode = modes[0]
		default:
			mode = 0777
		}

		if err := os.MkdirAll(e.Eval(path), os.FileMode(mode)); err != nil {
			e.shouldPanic(err.Error())
			e.shouldLog(err.Error())
		}
	}
}

// Rmdirs removes one or more space-separated directories
func (e *Echo) Rmdirs(paths string) {
	for _, path := range e.Split(paths) {
		if err := os.RemoveAll(path); err != nil {
			e.shouldPanic(err.Error())
			e.shouldLog(err.Error())
		}
	}
}

// AreSame tests files at path0 and path to check if they are the same
func (e *Echo) AreSame(path0, path1 string) bool {
	info0, err := os.Stat(path0)
	if err != nil {
		e.shouldPanic(err.Error())
		e.shouldLog(err.Error())
	}

	info1, err := os.Stat(path1)
	if err != nil {
		e.shouldPanic(err.Error())
		e.shouldLog(err.Error())
	}

	return os.SameFile(info0, info1)
}
