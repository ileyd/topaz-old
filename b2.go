package main

import (
	"net/url"
	"os"
	"strings"
)

// GenerateB2URL generates a streamable Backblaze B2 URL under the assumption that disk paths directly correlate to B2 paths, based on a given number of path components to strip
func GenerateB2URL(absPath string) (*url.URL, error) {
	pathComponents := strings.Split(absPath, string(os.PathSeparator))
	strippedComponents := config.B2.StrippedComponents
	relPath := strings.Join(pathComponents[(2*strippedComponents):], string(os.PathSeparator))

	// very bad way of unescaping the path separators
	urlRelPath := strings.Replace(url.QueryEscape(relPath), "%2F", "/", -1)
	return url.Parse(config.B2.BaseURL + urlRelPath)
}
