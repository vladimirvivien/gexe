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

// ResourceWriter represents types and methods used to post resource data to an HTTP server
type ResourceWriter struct {
	client  *http.Client
	err     error
	url     string
	headers http.Header
	data    io.Reader
	res     *Response
	vars    *vars.Variables
}

// Post starts a "POST" HTTP operation to the provided resource.
func Post(resource string) *ResourceWriter {
	return &ResourceWriter{url: resource, client: &http.Client{}, headers: make(http.Header), vars: &vars.Variables{}}
}

// PostWithVars sets up a "POST" operation and sets its session variables
func PostWithVars(resource string, variables *vars.Variables) *ResourceWriter {
	w := Post(variables.Eval(resource))
	w.vars = variables
	return w
}

// SetVars sets session variables for the ResourceWriter
func (w *ResourceWriter) SetVars(variables *vars.Variables) *ResourceWriter {
	w.vars = variables
	return w
}

// WithTimeout sets the HTTP client's timeout
func (w *ResourceWriter) WithTimeout(to time.Duration) *ResourceWriter {
	w.client.Timeout = to
	return w
}

// Err returns the last known error for the post operation
func (w *ResourceWriter) Err() error {
	return w.err
}

// Do is a terminal method that completes the post request of data to the HTTP server.
func (w *ResourceWriter) Do() *ResourceWriter {
	req, err := http.NewRequest("POST", w.url, w.data)
	if err != nil {
		w.err = err
		w.res = &Response{}
		return w
	}

	// set headers
	req.Header = w.headers

	// post request
	res, err := w.client.Do(req)
	if err != nil {
		w.err = err
		w.res = &Response{}
		return w
	}

	w.res = &Response{stat: res.Status, statCode: res.StatusCode, body: res.Body}

	return w
}

// WithHeaders sets all headers for the post operation
func (w *ResourceWriter) WithHeaders(h http.Header) *ResourceWriter {
	w.headers = h
	return w
}

// AddHeader is a convenience method to add a single header
func (w *ResourceWriter) AddHeader(key, value string) *ResourceWriter {
	w.headers.Add(w.vars.Eval(key), w.vars.Eval(value))
	return w
}

// SetHeader is a convenience method to sets a specific header
func (w *ResourceWriter) SetHeader(key, value string) *ResourceWriter {
	w.headers.Set(w.vars.Eval(key), w.vars.Eval(value))
	return w
}

// String posts the string value as content to the server
func (w *ResourceWriter) String(val string) *ResourceWriter {
	w.data = strings.NewReader(w.vars.Eval(val))
	return w.Do()
}

// Bytes posts the slice of bytes as content to the server
func (w *ResourceWriter) Bytes(val []byte) *ResourceWriter {
	w.data = bytes.NewReader(val)
	return w.Do()
}

// Body provides an io reader to stream content to the server
func (w *ResourceWriter) Body(val io.Reader) *ResourceWriter {
	w.data = val
	return w.Do()
}

// FormData posts form-encoded data as content to the server
func (w *ResourceWriter) FormData(val map[string][]string) *ResourceWriter {
	w.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	formData := url.Values(val)
	w.data = strings.NewReader(formData.Encode())
	return w.Do()
}
