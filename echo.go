package echo

type echo struct {
	vars map[string]string
}

// New creates a new session
func New() *echo {
	return &echo{
		vars: make(map[string]string),
	}
}
