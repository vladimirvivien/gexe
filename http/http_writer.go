package http

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/vladimirvivien/gexe/vars"
)

// ResourceWriter provides types and methods used to post resource data to an HTTP server
type ResourceWriter struct {
	client  *http.Client
	err     error
	url     string
	headers http.Header
	data    io.Reader
	vars    *vars.Variables
}

// Post starts an HTTP "POST" operation using the provided URL.
func Post(resource string) *ResourceWriter {
	return &ResourceWriter{url: resource, client: &http.Client{}, headers: make(http.Header), vars: &vars.Variables{}}
}

// PostWithVars starts an HTTP "POST" operation and sets the provided gexe session variables
func PostWithVars(resource string, variables *vars.Variables) *ResourceWriter {
	w := Post(variables.Eval(resource))
	w.vars = variables
	return w
}

// SetVars sets gexe session variables to be used by the post operation
func (w *ResourceWriter) SetVars(variables *vars.Variables) *ResourceWriter {
	w.vars = variables
	return w
}

// WithTimeout sets the HTTP client's timeout used for the post operation
func (w *ResourceWriter) WithTimeout(to time.Duration) *ResourceWriter {
	w.client.Timeout = to
	return w
}

// Err returns the last error generated by the ResourceWriter
func (w *ResourceWriter) Err() error {
	return w.err
}

// WithHeaders sets all HTTP headers to be used by the post operation
func (w *ResourceWriter) WithHeaders(h http.Header) *ResourceWriter {
	w.headers = h
	return w
}

// AddHeader is a convenience method to add a single header for the post operation
func (w *ResourceWriter) AddHeader(key, value string) *ResourceWriter {
	w.headers.Add(w.vars.Eval(key), w.vars.Eval(value))
	return w
}

// SetHeader is a convenience method to sets a specific header for the post operation
func (w *ResourceWriter) SetHeader(key, value string) *ResourceWriter {
	w.headers.Set(w.vars.Eval(key), w.vars.Eval(value))
	return w
}

// String posts the string value content to the server
// and returns a gexe/http/*Response
func (w *ResourceWriter) String(val string) *ResourceWriter {
	w.data = strings.NewReader(w.vars.Eval(val))
	return w
}

// Bytes posts the slice of bytes as content to the server
func (w *ResourceWriter) Bytes(val []byte) *ResourceWriter {
	w.data = bytes.NewReader(val)
	return w
}

// Body provides an io reader to stream content to the server
func (w *ResourceWriter) Body(val io.Reader) *ResourceWriter {
	w.data = val
	return w
}

// FormData posts form-encoded data as content to the server
func (w *ResourceWriter) FormData(val map[string][]string) *ResourceWriter {
	w.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	formData := url.Values(val)
	w.data = strings.NewReader(formData.Encode())
	return w
}

// Do is a terminal method that actually posts the HTTP request to the server.
// It returns a gexe/http/*Response instance that can be used to access post result.
func (w *ResourceWriter) Do() *Response {
	req, err := http.NewRequest("POST", w.url, w.data)
	if err != nil {
		return &Response{err: err}
	}

	// set headers
	req.Header = w.headers

	// post request
	res, err := w.client.Do(req)
	if err != nil {
		return &Response{err: err}
	}

	return &Response{stat: res.Status, statCode: res.StatusCode, body: res.Body}
}
