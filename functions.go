package gexe

import (
	"context"
	"io"
	"os"

	"github.com/vladimirvivien/gexe/exec"
	"github.com/vladimirvivien/gexe/fs"
	"github.com/vladimirvivien/gexe/http"
	"github.com/vladimirvivien/gexe/prog"
	"github.com/vladimirvivien/gexe/str"
	"github.com/vladimirvivien/gexe/vars"
)

// Variables returns variable map for DefaultEcho session
func Variables() *vars.Variables {
	return DefaultSession.Variables()
}

// Envs declares environment variables using
// a multi-line space-separated list:
//
//	Envs("GOOS=linux GOARCH=amd64")
//
// Environment vars can be used in string values
// using Eval("building for os=$GOOS")
func Envs(val ...string) *Session {
	return DefaultSession.Envs(val...)
}

// SetEnv sets a process environment variable.
func SetEnv(name, value string, args ...interface{}) *Session {
	return DefaultSession.SetEnv(name, value, args...)
}

// Vars declares multiple session-scope variables using
// string literals:
//
//	Envs("foo=bar", "platform=amd64", `"data="info ${platform}"`)
//
// Note that session vars are only available
// for the running process.
func Vars(variables ...string) *Session {
	return DefaultSession.Vars(variables...)
}

// SetVar declares a session variable.
func SetVar(name, value string, args ...interface{}) *Session {
	return DefaultSession.SetVar(name, value, args...)
}

// Val retrieves a session or environment variable
func Val(name string) string {
	return DefaultSession.Val(name)
}

// Eval returns the string str with its content expanded
// with variable values i.e. Eval("I am $HOME") returns
// "I am </user/home/path>"
func Eval(str string, args ...interface{}) string {
	return DefaultSession.Eval(str, args...)
}

// NewProcWithContext setups a new process with specified context and command cmdStr and returns immediately
// without starting. Information about the running process is stored in *exec.Proc.
func NewProcWithContext(ctx context.Context, cmdStr string, args ...interface{}) *exec.Proc {
	return DefaultSession.NewProcWithContext(ctx, cmdStr, args...)
}

// NewProc setups a new process with specified command cmdStr and returns immediately
// without starting. Information about the running process is stored in *exec.Proc.
func NewProc(cmdStr string, args ...interface{}) *exec.Proc {
	return DefaultSession.NewProcWithContext(context.Background(), cmdStr, args...)
}

// StartProcWith executes the command in cmdStr with the specified contex and returns immediately
// without waiting. Information about the running process is stored in *exec.Proc.
func StartProcWithContext(ctx context.Context, cmdStr string, args ...interface{}) *exec.Proc {
	return DefaultSession.StartProcWithContext(ctx, cmdStr, args...)
}

// StartProc executes the command in cmdStr and returns immediately
// without waiting. Information about the running process is stored in *exec.Proc.
func StartProc(cmdStr string, args ...interface{}) *exec.Proc {
	return DefaultSession.StartProc(cmdStr, args...)
}

// RunProcWithContext executes command in cmdStr, with specified ctx, and waits for the result.
// It returns a *Proc with information about the executed process.
func RunProcWithContext(ctx context.Context, cmdStr string, args ...interface{}) *exec.Proc {
	return DefaultSession.RunProcWithContext(ctx, cmdStr, args...)
}

// RunProc executes command in cmdStr and waits for the result.
// It returns a *Proc with information about the executed process.
func RunProc(cmdStr string, args ...interface{}) *exec.Proc {
	return DefaultSession.RunProc(cmdStr, args...)
}

// RunWithContext executes cmdStr, with specified context, and waits for completion.
// It returns the result as a string.
func RunWithContext(ctx context.Context, cmdStr string, args ...interface{}) string {
	return DefaultSession.RunWithContext(ctx, cmdStr, args...)
}

// Run executes cmdStr, waits, and returns the result as a string.
func Run(cmdStr string, args ...interface{}) string {
	return DefaultSession.Run(cmdStr, args...)
}

// Runout executes command cmdStr and prints out the result
func Runout(cmdStr string, args ...interface{}) {
	DefaultSession.Runout(cmdStr, args...)
}

// Commands returns a *exe.CommandBuilder to build a multi-command execution flow.
func Commands(cmdStrs ...string) *exec.CommandBuilder {
	return DefaultSession.Commands(cmdStrs...)
}

// StartAll starts the exection of each command sequentially and
// does not wait for their completion.
func StartAll(cmdStrs ...string) *exec.CommandResult {
	return DefaultSession.StartAll(cmdStrs...)
}

// RunAll executes each command, in cmdStrs, successively and wait for their
// completion.
func RunAll(cmdStrs ...string) *exec.CommandResult {
	return DefaultSession.RunAll(cmdStrs...)
}

// StartConcur starts the exection of each command concurrently and
// does not wait for their completion.
func StartConcur(cmdStrs ...string) *exec.CommandResult {
	return DefaultSession.StartConcur(cmdStrs...)
}

// RunConcur executes each command, in cmdStrs, concurrently and waits
// their completion.
func RunConcur(cmdStrs ...string) *exec.CommandResult {
	return DefaultSession.RunConcur(cmdStrs...)
}

// Pipe executes each command, in cmdStrs, by piping the result
// of the previous command as input to the next command until done.
func Pipe(cmdStrs ...string) *exec.PipedCommandResult {
	return DefaultSession.Pipe(cmdStrs...)
}

// PathExists returns true if specified path exists.
// Any error will cause it to return false.
func PathExists(path string, args ...interface{}) bool {
	return DefaultSession.PathExists(path, args...)
}

// PathInfo returns information for specified path (i.e. size, etc)
func PathInfo(path string, args ...interface{}) *fs.FSInfo {
	return DefaultSession.PathInfo(path, args...)
}

// MkDirs creates one or more directories along the specified path
func MkDirs(path string, mode os.FileMode, args ...interface{}) *fs.FSInfo {
	return DefaultSession.MkDir(path, mode, args...)
}

// MkDir creates a directory with default mode 0744
func MkDir(path string, args ...interface{}) *fs.FSInfo {
	return DefaultSession.MkDir(path, 0744, args...)
}

// RmPath removes files or directories along specified path
func RmPath(path string, args ...interface{}) *fs.FSInfo {
	return DefaultSession.RmPath(path, args...)
}

// FileRead uses context ctx to read file content from path
func FileReadWithContext(ctx context.Context, path string, args ...interface{}) *fs.FileReader {
	return DefaultSession.FileReadWithContext(ctx, path, args...)
}

// FileRead provides methods to read file content from path
func FileRead(path string, args ...interface{}) *fs.FileReader {
	return DefaultSession.FileReadWithContext(context.Background(), path, args...)
}

// FileWriteWithContext uses context ctx to write file content to path
func FileWriteWithContext(ctx context.Context, path string, args ...interface{}) *fs.FileWriter {
	return DefaultSession.FileWriteWithContext(ctx, path, args...)
}

// FileWrite provides methods to write file content to path
func FileWrite(path string, args ...interface{}) *fs.FileWriter {
	return DefaultSession.FileWriteWithContext(context.Background(), path, args...)
}

// HttpGetWithContext uses context ctx to start an HTTP GET operation to retrieve resource at URL/path
func HttpGetWithContext(ctx context.Context, url string, paths ...string) *http.ResourceReader {
	return DefaultSession.HttpGetWithContext(ctx, url, paths...)
}

// HttpGet starts an HTTP GET operation to retrieve resource at URL/path
func HttpGet(url string, paths ...string) *http.ResourceReader {
	return DefaultSession.HttpGetWithContext(context.Background(), url, paths...)
}

// Get is a convenient alias for HttpGet that retrieves specified resource at given URL/path
func Get(url string, paths ...string) *http.Response {
	return DefaultSession.Get(url, paths...)
}

// HttpPostWithContext uses context ctx to start an HTTP POST operation to post resource to URL/path
func HttpPostWithContext(ctx context.Context, url string, paths ...string) *http.ResourceWriter {
	return DefaultSession.HttpPostWithContext(ctx, url, paths...)
}

// HttpPost starts an HTTP POST operation to post resource to URL/path
func HttpPost(url string, paths ...string) *http.ResourceWriter {
	return DefaultSession.HttpPostWithContext(context.Background(), url, paths...)
}

// Post is a convenient alias for HttpPost to post data at specified URL
func Post(data []byte, url string) *http.Response {
	return DefaultSession.Post(data, url)
}

// Prog returns program information via *prog.Info
func Prog() *prog.Info {
	return DefaultSession.Prog()
}

// ProgAvail returns the full path of the program if available.
func ProgAvail(program string, args ...interface{}) string {
	return DefaultSession.ProgAvail(program, args...)
}

// Workdir returns the current program's working directory
func Workdir() string {
	return DefaultSession.Workdir()
}

// AddExecPath adds an executable path to PATH
func AddExecPath(execPath string) {
	DefaultSession.AddExecPath(execPath)
}

func String(s string, args ...interface{}) *str.Str {
	return DefaultSession.String(s, args...)
}

// Print outputs the formatted string to stdout without a newline.
// The string supports both Go's fmt.Sprintf formatting and gexe variable expansion.
// Variables are expanded using ${VAR} or $VAR syntax.
//
// Example:
//
//	gexe.Print("Processing %s in ${HOME}", filename)
func Print(format string, args ...interface{}) *Session {
	return DefaultSession.Print(format, args...)
}

// Println outputs the formatted string to stdout with a trailing newline.
// The string supports both Go's fmt.Sprintf formatting and gexe variable expansion.
// Variables are expanded using ${VAR} or $VAR syntax.
//
// Example:
//
//	gexe.Println("User ${USER} logged in at %s", timestamp)
func Println(format string, args ...interface{}) *Session {
	return DefaultSession.Println(format, args...)
}

// PrintTo outputs the formatted string to the specified writer.
// The string supports both Go's fmt.Sprintf formatting and gexe variable expansion.
// Variables are expanded using ${VAR} or $VAR syntax.
//
// Example:
//
//	var buf bytes.Buffer
//	gexe.PrintTo(&buf, "Log: ${USER} performed %s", action)
func PrintTo(w io.Writer, format string, args ...interface{}) *Session {
	return DefaultSession.PrintTo(w, format, args...)
}

// Error creates and returns an error with the formatted string.
// The string supports both Go's fmt.Sprintf formatting and gexe variable expansion.
// Variables are expanded using ${VAR} or $VAR syntax.
//
// Example:
//
//	err := gexe.Error("Failed to process %s in ${HOME}", filename)
//	return err
func Error(format string, args ...interface{}) error {
	return DefaultSession.Error(format, args...)
}

// Join uses strings.Join to join arbitrary strings with a separator
// while applying gexe variable expansion to each element.
//
// Example:
//
//	gexe.SetVar("HOME", "/home/user")
//	result := gexe.Join(",", "${HOME}", "documents", "file.txt")
//	// Returns: "/home/user,documents,file.txt"
func Join(sep string, elem ...string) string {
	return DefaultSession.Join(sep, elem...)
}

// JoinPath uses filepath.Join to join file paths using OS-specific path separators
// while applying gexe variable expansion to each element.
//
// Example:
//
//	gexe.SetVar("HOME", "/home/user")
//	path := gexe.JoinPath("${HOME}", "documents", "file.txt")
//	// Returns: "/home/user/documents/file.txt" (Unix) or "C:\Users\user\documents\file.txt" (Windows)
func JoinPath(elem ...string) string {
	return DefaultSession.JoinPath(elem...)
}
