package http

import (
	"net/http"
	"time"

	"github.com/vladimirvivien/gexe/vars"
)

// ResourceReader provides types and methods to read content of resources from a server using HTTP
type ResourceReader struct {
	client *http.Client
	err    error
	url    string
	vars   *vars.Variables
}

// Get initiates a "GET" operation for the specified resource
func Get(url string) *ResourceReader {
	return &ResourceReader{url: url, client: &http.Client{}, vars: &vars.Variables{}}
}

// Get initiates a "GET" operation and sets session variables
func GetWithVars(url string, variables *vars.Variables) *ResourceReader {
	r := Get(variables.Eval(url))
	r.vars = variables
	return r
}

// SetVars sets session variables for ResourceReader
func (r *ResourceReader) SetVars(variables *vars.Variables) *ResourceReader {
	r.vars = variables
	return r
}

// Err returns the last known error
func (r *ResourceReader) Err() error {
	return r.err
}

// WithTimeout sets the HTTP reader's timeout
func (r *ResourceReader) WithTimeout(to time.Duration) *ResourceReader {
	r.client.Timeout = to
	return r
}

// Do is a terminal method that actually retrieves the HTTP resource from the server.
// It returns a gexe/http/*Response instance that can be used to access the result.
func (r *ResourceReader) Do() *Response {
	res, err := r.client.Get(r.url)
	if err != nil {
		return &Response{err: err}
	}
	return &Response{stat: res.Status, statCode: res.StatusCode, body: res.Body}
}
