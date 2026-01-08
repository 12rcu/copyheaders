// Package copyheaders plugin middleware to copy a header value to another header.
package copyheaders

import (
	"context"
	"net/http"
)

// HeaderConfig define to and from header so to copy.
type HeaderConfig struct {
	From   string `json:"from,omitempty"`
	To     string `json:"to,omitempty"`
	Prefix string `json:"prefix,omitempty"`
}

// Config the plugin configuration.
type Config struct {
	Headers []HeaderConfig `json:"headers,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// RewriteHeader structure for copying the headers.
type RewriteHeader struct {
	next          http.Handler
	name          string
	replaceHeader []HeaderConfig
}

// New created a new RewriteHeader plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &RewriteHeader{
		next:          next,
		name:          name,
		replaceHeader: config.Headers,
	}, nil
}

func (a *RewriteHeader) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	for _, config := range a.replaceHeader {
		headerValue := req.Header.Get(config.From)
		if len(headerValue) == 0 {
			continue
		}
		req.Header.Set(config.To, config.Prefix+headerValue)
	}
	a.next.ServeHTTP(rw, req)
}
