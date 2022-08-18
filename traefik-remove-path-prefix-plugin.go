package traefik_remove_path_prefix_plugin

import (
	"context"
	"net/http"
	"strings"
)

type Config struct {
	ForceSlash bool `ForceSlash:"rewrites,omitempty"`
}

// CreateConfig creates and initializes the plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// removePathPrefix is a middleware used to remove the prefix from an URL request.
type removePathPrefix struct {
	next       http.Handler
	forceSlash bool
	name       string
}

// New created a new removePathPrefix plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &removePathPrefix{
		forceSlash: config.ForceSlash,
		name:       name,
		next:       next,
	}, nil
}

func (r *removePathPrefix) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// Remove the first path argument of the URL.
	var separator byte = '/'
	var sanitazeUrl = strings.TrimPrefix(req.URL.Path, string(separator))
	_, urlPath := split(sanitazeUrl, separator, false)

	// Checks if the slash should be forced at the end of the path.
	if r.forceSlash {
		if urlPath == "" {
			urlPath = "/"
		} else {
			urlPath = "/" + strings.TrimPrefix(urlPath, string(separator)) + "/"
		}
	}

	// Change the URL original req for the URL without prefix.
	req.URL.Path = urlPath
	if req.URL.RawPath != "" {
		req.URL.RawPath = urlPath
	}
	req.RequestURI = req.URL.RequestURI()

	r.next.ServeHTTP(rw, req)
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
