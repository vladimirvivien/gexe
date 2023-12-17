package http

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHttpReader_String(t *testing.T) {
	tests := []struct {
		name string
		data []byte
	}{
		{name: "no data"},
		{name: "string0", data: []byte("Hello World!")},
		{name: "string1", data: []byte("Hello\nWorld!")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, string(test.data))
			}))
			defer ts.Close()

			r := Get(ts.URL).Do()
			if r.Err() != nil {
				t.Fatal(r.Err())
			}

			if strings.TrimSpace(r.String()) != string(test.data) {
				t.Fatal("unexpected server response: ", r.String())
			}
		})
	}
}

func TestHttpReader_Bytes(t *testing.T) {
	tests := []struct {
		name string
		data string
	}{
		{name: "no data"},
		{name: "string0", data: "Hello World!"},
		{name: "string1", data: "Hello\nWorld!"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, string(test.data))
			}))
			defer ts.Close()

			r := Get(ts.URL).Do()
			if r.Err() != nil {
				t.Fatal(r.Err())
			}

			data := r.Bytes()
			dataLen := len(data)

			if dataLen != len(test.data) {
				t.Fatalf("unexpected server response length: %d ", dataLen)
			}

			if strings.TrimSpace(string(data)) != test.data {
				t.Fatalf("unexpected server response: %s", r.String())
			}
		})
	}
}

func TestHttpReader_Body(t *testing.T) {
	tests := []struct {
		name string
		data []byte
	}{
		{name: "no data"},
		{name: "string0", data: []byte("Hello World!")},
		{name: "string1", data: []byte("Hello\nWorld!")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, string(test.data))
			}))
			defer ts.Close()

			r := Get(ts.URL).Do()
			if r.Err() != nil {
				t.Fatal(r.Err())
			}

			var result bytes.Buffer
			if _, err := io.Copy(&result, r.Body()); err != nil {
				t.Fatal(err)
			}
			if err := r.Body().Close(); err != nil {
				t.Fatal(err)
			}
			if strings.TrimSpace(result.String()) != string(test.data) {
				t.Fatal("unexpected server response: ", r.String())
			}
		})
	}
}

func TestHttpReader_Do(t *testing.T) {
	tests := []struct {
		name   string
		status int
		data   []byte
	}{
		{name: "no data", status: 200},
		{name: "normal request", data: []byte("Hello World!"), status: 200},
		{name: "missing resource", status: 404},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if test.status > 299 {
					w.WriteHeader(test.status)
					return
				}
				fmt.Fprint(w, string(test.data))
			}))
			defer ts.Close()

			r := Get(ts.URL).Do()
			if r.Err() != nil {
				t.Fatal(r.Err())
			}

			dataLen := len(r.Bytes())
			if dataLen != len(test.data) {
				t.Fatal("unexpected server response length: ", dataLen)
			}

			if r.StatusCode() != test.status {
				t.Fatal("got unexpected status code: ", r.StatusCode())
			}
		})
	}
}
