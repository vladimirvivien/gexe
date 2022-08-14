package http

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func runServer(t *testing.T, expected []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != string(expected) {
			t.Fatal("server received unexpected data:", string(data))
		}
	}))
}

func TestHttpWriter_String(t *testing.T) {
	tests := []struct {
		name string
		data []byte
	}{
		{name: "no data"},
		{name: "simple data", data: []byte("Hello World!!!")},
		{name: "multilines data", data: []byte("Hello\nWorld!!!")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := runServer(t, test.data)
			defer server.Close()

			w := Post(server.URL)
			if w.Err() != nil {
				t.Fatal(w.Err())
			}
			w.String(string(test.data))
		})
	}
}

func TestHttpWriter_Bytes(t *testing.T) {
	tests := []struct {
		name string
		data []byte
	}{
		{name: "no data"},
		{name: "simple data", data: []byte("Hello World!!!")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := runServer(t, test.data)
			defer server.Close()

			w := Post(server.URL)
			if w.Err() != nil {
				t.Fatal(w.Err())
			}
			w.Bytes(test.data)
		})
	}
}

func TestHttpWriter_Body(t *testing.T) {
	tests := []struct {
		name string
		data []byte
	}{
		{name: "no data"},
		{name: "simple data", data: []byte("Hello World!!!")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := runServer(t, test.data)
			defer server.Close()

			w := Post(server.URL)
			if w.Err() != nil {
				t.Fatal(w.Err())
			}
			w.Body(bytes.NewReader(test.data))
		})
	}
}
