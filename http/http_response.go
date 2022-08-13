package http

import "io"

// HttpResponse stores high level metadata and response body from the server
type HttpResponse struct {
	stat     string
	statCode int
	body     io.ReadCloser
}

// Status returns the standard lib http.Response.Status value from the server
func (res *HttpResponse) Status() string {
	return res.stat
}

// StatusCode returns the standard lib http.Response.StatusCode value from the server
func (res *HttpResponse) StatusCode() int {
	return res.statCode
}

// Body is io.ReadCloser stream to the content from serve.
// NOTE: ensure to call Close() if used directly.
func (res *HttpResponse) Body() io.ReadCloser {
	return res.body
}
