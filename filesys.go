package gexe

import (
	"context"
	"os"

	"github.com/vladimirvivien/gexe/fs"
)

// PathExists returns true if path exists.
// All errors causes to return false.
func (e *Session) PathExists(path string, args ...interface{}) bool {
	path = applyFmt(path, args...)
	return fs.PathWithVars(path, e.vars).Exists()
}

// MkDir creates a directory at specified path with mode value.
// FSInfo contains information about the path or error if occured
func (e *Session) MkDir(path string, mode os.FileMode, args ...interface{}) *fs.FSInfo {
	path = applyFmt(path, args...)
	p := fs.PathWithVars(path, e.vars)
	return p.MkDir(mode)
}

// RmPath removes specified path (dir or file).
// Error is returned FSInfo.Err()
func (e *Session) RmPath(path string, args ...interface{}) *fs.FSInfo {
	path = applyFmt(path, args...)
	p := fs.PathWithVars(path, e.vars)
	return p.Remove()
}

// PathInfo
func (e *Session) PathInfo(path string, args ...interface{}) *fs.FSInfo {
	path = applyFmt(path, args...)
	return fs.PathWithVars(path, e.vars).Info()
}

// FileReadWithContext uses specified context to provide methods to read file
// content at path.
func (e *Session) FileReadWithContext(ctx context.Context, path string, args ...interface{}) *fs.FileReader {
	path = applyFmt(path, args...)
	return fs.ReadWithContextVars(ctx, path, e.vars)
}

// FileRead provides methods to read file content
func (e *Session) FileRead(path string, args ...interface{}) *fs.FileReader {
	path = applyFmt(path, args...)
	return fs.ReadWithContextVars(context.Background(), path, e.vars)
}

// FileWriteWithContext uses context ctx to create a fs.FileWriter to write content to provided path
func (e *Session) FileWriteWithContext(ctx context.Context, path string, args ...interface{}) *fs.FileWriter {
	path = applyFmt(path, args...)
	return fs.WriteWithContextVars(ctx, path, e.vars)
}

// FileWrite creates a fs.FileWriter to write content to provided path
func (e *Session) FileWrite(path string, args ...interface{}) *fs.FileWriter {
	path = applyFmt(path, args...)
	return fs.WriteWithContextVars(context.Background(), path, e.vars)
}

// FileAppend creates a new fs.FileWriter to append content to provided path
func (e *Session) FileAppendWithContext(ctx context.Context, path string, args ...interface{}) *fs.FileWriter {
	path = applyFmt(path, args...)
	return fs.AppendWithContextVars(ctx, path, e.vars)
}

// FileAppend creates a new fs.FileWriter to append content to provided path
func (e *Session) FileAppend(path string, args ...interface{}) *fs.FileWriter {
	path = applyFmt(path, args...)
	return fs.AppendWithContextVars(context.Background(), path, e.vars)
}
