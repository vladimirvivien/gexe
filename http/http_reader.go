package http

import (
	"io"
	"net/http"
)

// HttpReader provides types and methods to read content from a server using HTTP
type HttpReader struct {
	client *http.Client
	err    error
	url    string
	res    *HttpResponse
}

// Get initiates a "GET" operation for the specified resource
func Get(url string) *HttpReader {
	return &HttpReader{url: url, client: &http.Client{}}
}

// Err returns the last known error
func (r *HttpReader) Err() error {
	return r.err
}

// Response returns the server's response info
func (r *HttpReader) Response() *HttpResponse {
	return r.res
}

// Bytes returns the server response as a []byte
func (b *HttpReader) Bytes() []byte {
	if b.Do().Err() != nil {
		return nil
	}
	return b.read()
}

// String returns the server response as a string
func (b *HttpReader) String() string {
	if b.Do().Err() != nil {
		return ""
	}
	return string(b.read())
}

// Body returns an io.ReadCloser to stream the server response.
// NOTE: ensure to close the stream when finished.
func (r *HttpReader) Body() io.ReadCloser {
	if r.Do().Err() != nil {
		return nil
	}
	return r.res.body
}

// Do invokes the client.Get to "GET" the content from server
func (r *HttpReader) Do() *HttpReader {
	res, err := r.client.Get(r.url)
	if err != nil {
		r.err = err
		r.res = &HttpResponse{}
		return r
	}
	r.res = &HttpResponse{stat: res.Status, statCode: res.StatusCode, body: res.Body}
	return r
}

// read reads the content of the response body and returns a []byte
func (r *HttpReader) read() []byte {
	if r.res.body == nil {
		return nil
	}

	data, err := io.ReadAll(r.res.body)
	defer r.res.body.Close()
	if err != nil {
		r.err = err
		return nil
	}
	return data
}
