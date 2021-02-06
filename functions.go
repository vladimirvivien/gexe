package echo

import (
	"github.com/vladimirvivien/echo/exec"
)

// Package level functions that wraps Echo methods
// with support for variable expansions.


// StartProc executes the command in cmdStr and returns immediately
// without waiting. Information about the running process is stored in *Proc.
func StartProc(cmdStr string) *exec.Proc {
	return DefaultEcho.StartProc(cmdStr)
}

// RunProc executes command in cmdStr and waits for the result.
// It returns a *Proc with information about the executed process.
func RunProc(cmdStr string) *exec.Proc {
	return DefaultEcho.RunProc(cmdStr)
}

// Run executes cmdStr, waits, and returns the result as a string.
func Run(cmdStr string) string {
	return DefaultEcho.Run(cmdStr)
}

// Runout executes command cmdStr and prints out the result
func Runout(cmdStr string) {
	DefaultEcho.Runout(cmdStr)
}

// Abs returns the absolute representation for path.
// Returns empty string if it fails for any reason.
func Abs(path string) string {
	return DefaultEcho.Abs(path)
}

// Rel returns path that is relative to base/target.
// Returns an empty path if error.
func Rel(base string, target string) string {
	return DefaultEcho.Rel(base, target)
}

// Base returns the last portion (or name) of a path.
func Base(path string) string {
	return DefaultEcho.Base(path)
}

// Dir returns parent path portion of a path (without the Base).
func Dir(path string) string {
	return DefaultEcho.Dir(path)
}

// PathSplit splits an OS-specific list of path into []string
func PathSplit(path string) []string {
	return DefaultEcho.PathSplit(path)
}

// PathSym returns the evaluated symbolic link for path.
// Returns empty if symbolic evaluation fails.
func PathSym(path string) string {
	return DefaultEcho.PathSym(path)
}

// Ext returns extension part of path
func Ext(path string) string {
	return DefaultEcho.Ext(path)
}

// PathJoin collate individual paths together for a longer path
func PathJoin(paths ...string) string {
	return DefaultEcho.PathJoin(paths...)
}

// PathMatched returns true if path matches shell file pattern
func PathMatched(pattern, path string) bool {
	return DefaultEcho.PathMatched(pattern, path)
}

// IsAbs returns true if path is an absolute path
func IsAbs(path string) bool {
	return DefaultEcho.IsAbs(path)
}

// IsExist returns true if file is 1) accessible 2) exists.
// All other conditions returns false.
func IsExist(path string) bool {
	return DefaultEcho.IsExist(path)
}

// IsReg returns true if path is a regular file.
// All other cases (errors) will return false.
func IsReg(path string) bool {
	return DefaultEcho.IsReg(path)
}

// IsDir returns true if path is a directory.
// All other cases (errors) will return false.
func IsDir(path string) bool {
	return DefaultEcho.IsDir(path)
}

// Mkdirs creates one or more space-separated directories using the optional filemods.
// If no mode is provided, it defaults to 0777. When mode is provided only
// modes[0] is applied.
func Mkdirs(paths string, modes ...uint32) {
	DefaultEcho.Mkdirs(paths,modes...)
}

// Rmdirs removes one or more space-separated directories
func Rmdirs(paths string) {
	DefaultEcho.Rmdirs(paths)
}

// AreSame tests files at path0 and path to check if they are the same
func AreSame(path0, path1 string) bool {
	return DefaultEcho.AreSame(path0, path1)
}
