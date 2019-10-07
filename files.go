package echo

import (
	"os"
	"path/filepath"
)

// PathAbs returns the absolute representation for path.
// Returns empty string if it fails for any reason.
func (e *echo) PathAbs(path string) string {
	abs, err := filepath.Abs(os.Expand(path, e.Val))
	if err != nil {
		return ""
	}
	return abs
}

// PathRel returns path that is relative to base/target.
// Returns an empty path if error.
func (e *echo) PathRel(base string, target string) string {
	abs, err := filepath.Rel(os.Expand(base, e.Val), os.Expand(target, e.Val))
	if err != nil {
		return ""
	}
	return abs
}

// PathBase returns the last portion (or name) of a path.
func (e *echo) PathBase(path string) string {
	return filepath.Base(os.Expand(path, e.Val))
}

// PathDir returns parent path portion of a path (without the Base).
func (e *echo) PathDir(path string) string {
	return filepath.Dir(os.Expand(path, e.Val))
}

// PathSplit splits an OS-specific list of path into []string
func (e *echo) PathSplit(path string) []string {
	return filepath.SplitList(os.Expand(path, e.Val))
}

// PathSym returns the evaluated symbolic link for path.
// Returns empty if symbolic evaluation fails.
func (e *echo) PathSym(path string) string {
	link, err := filepath.EvalSymlinks(os.Expand(path, e.Val))
	if err != nil {
		return ""
	}
	return link
}

// PathExt returns extension part of path
func (e *echo) PathExt(path string) string {
	return filepath.Ext(os.Expand(path, e.Val))
}

// JoinPaths collate individual paths together for a longer path
func (e *echo) PathJoin(paths ...string) string {
	for i, path := range paths {
		paths[i] = os.Expand(path, e.Val)
	}
	return filepath.Join(paths...)
}

// PathMatched returns true if path matches shell file pattern
func (e *echo) PathMatched(pattern, path string) bool {
	matched, err := filepath.Match(pattern, os.Expand(path, e.Val))
	if err != nil {
		return false
	}
	return matched
}

// IsAbs returns true if path is an absolute path
func (e *echo) IsPathAbs(path string) bool {
	return filepath.IsAbs(os.Expand(path, e.Val))
}

// IsPathExist returns true if file is 1) accessible 2) exists.
// All other conditions returns false.
func (e *echo) IsPathExist(path string) bool {
	if _, err := os.Stat(os.Expand(path, e.Val)); err != nil {
		return false
	}
	return true
}

// IsPathReg returns true if path is a regular file.
// All other cases (errors) will return false.
func (e *echo) IsPathReg(path string) bool {
	info, err := os.Stat(os.Expand(path, e.Val))
	if err != nil {
		return false
	}

	return info.Mode().IsRegular()
}

// IsPathDir returns true if path is a directory.
// All other cases (errors) will return false.
func (e *echo) IsPathDir(path string) bool {
	info, err := os.Stat(os.Expand(path, e.Val))
	if err != nil {
		return false
	}
	return info.Mode().IsDir()
}
