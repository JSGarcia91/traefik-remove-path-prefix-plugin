package traefik_remove_path_prefix_plugin

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

type Config struct {
}

// CreateConfig creates and initializes the plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// RemovePathPrefix a plugin.
type RemovePathPrefix struct {
	next http.Handler
	name string
}

// New created a new plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &RemovePathPrefix{
		name: name,
		next: next,
	}, nil
}

func (e *RemovePathPrefix) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// Remove the first path argument of the URL.
	var separator byte = '/'
	var sanitazeUrl = strings.Trim(req.URL.Path, "/")
	_, path := split(sanitazeUrl, separator, false)

	// Change the URL original req for the URL without prefix.
	req.URL = &url.URL{Path: path}

	e.next.ServeHTTP(rw, req)
}

// split slices s into two substrings separated by the first occurrence of
// sep. If cutc is true then sep is excluded from the second substring.
// If sep does not occur in s then s and the empty string is returned.
func split(s string, sep byte, cutc bool) (string, string) {
	i := strings.IndexByte(s, sep)
	if i < 0 {
		return s, ""
	}
	if cutc {
		return s[:i], s[i+1:]
	}
	return s[:i], s[i:]
}
