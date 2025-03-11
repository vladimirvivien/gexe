//go:build windows

package exec

func (p *Proc) setUser() {
	// Windows does not support setting user/group IDs directly in the same way as Unix
	// This is a no-op implementation for Windows
	return
}
