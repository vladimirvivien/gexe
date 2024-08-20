package gexe

import (
	"strings"

	"github.com/vladimirvivien/gexe/http"
)

// Get uses an HTTP GET operation to retrieve server resources.
func (e *Echo) Get(url string, paths ...string) *http.Response {
	var exapandedUrl strings.Builder
	exapandedUrl.WriteString(e.vars.Eval(url))
	for _, path := range paths {
		exapandedUrl.WriteString(e.vars.Eval(path))
	}
	return http.GetWithVars(exapandedUrl.String(), e.vars).Do()
}

// Post uses an HTTP POST operation to post data to the server.
func (e *Echo) Post(data []byte, url string, paths ...string) *http.Response {
	var exapandedUrl strings.Builder
	exapandedUrl.WriteString(e.vars.Eval(url))
	for _, path := range paths {
		exapandedUrl.WriteString(e.vars.Eval(path))
	}
	return http.PostWithVars(exapandedUrl.String(), e.vars).Bytes(data).Do()
}
