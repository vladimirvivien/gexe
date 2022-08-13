package http

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// HttpWriter represents types and methods used to post data to an HTTP server
type HttpWriter struct {
	client  *http.Client
	err     error
	url     string
	headers http.Header
	data    io.Reader
	res     *HttpResponse
}

// Post starts a "POST" HTTP operation to the provided resource.
func Post(resource string) *HttpWriter {
	return &HttpWriter{url: resource, client: &http.Client{}, headers: make(http.Header)}
}

// Err returns the last known error for the post operation
func (w *HttpWriter) Err() error {
	return w.err
}

// Do is a terminal method that completes the post request of data to the HTTP server.
func (w *HttpWriter) Do() *HttpWriter {
	req, err := http.NewRequest("POST", w.url, w.data)
	if err != nil {
		w.err = err
		w.res = &HttpResponse{}
		return w
	}

	// set headers
	req.Header = w.headers

	// post request
	res, err := w.client.Do(req)
	if err != nil {
		w.err = err
		w.res = &HttpResponse{}
		return w
	}

	w.res = &HttpResponse{stat: res.Status, statCode: res.StatusCode, body: res.Body}

	return w
}

// WithHeaders sets all headers for the post operation
func (w *HttpWriter) WithHeaders(h http.Header) *HttpWriter {
	w.headers = h
	return w
}

// AddHeader is a convenience method to add a single header
func (w *HttpWriter) AddHeader(key, value string) *HttpWriter {
	w.headers.Add(key, value)
	return w
}

// SetHeader is a convenience method to sets a specific header
func (w *HttpWriter) SetHeader(key, value string) *HttpWriter {
	w.headers.Set(key, value)
	return w
}

// String posts the string value as content to the server
func (w *HttpWriter) String(val string) *HttpWriter {
	w.data = strings.NewReader(val)
	return w.Do()
}

// Bytes posts the slice of bytes as content to the server
func (w *HttpWriter) Bytes(val []byte) *HttpWriter {
	w.data = bytes.NewReader(val)
	return w.Do()
}

// Body provides an io reader to stream content to the server
func (w *HttpWriter) Body(val io.Reader) *HttpWriter {
	w.data = val
	return w.Do()
}

// FormData posts form-encoded data as content to the server
func (w *HttpWriter) FormData(val map[string][]string) *HttpWriter {
	w.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	formData := url.Values(val)
	w.data = strings.NewReader(formData.Encode())
	return w.Do()
}
