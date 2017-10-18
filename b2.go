package main

import (
	"errors"
	"log"
	"net/url"
	"strings"
)

// GenerateB2URL generates a streamable Backblaze B2 URL under the assumption that disk paths directly correlate to B2 paths, based on a given number of path components to strip
func GenerateB2URL(absPath string) (*url.URL, error) {
	// NOTE: string(os.PathSeparator) makes more sense but in dev I'm running the sonarr server on a Linux machine while developing on a windows one
	pathComponents := strings.Split(absPath, "/")
	strippedComponents := 1 + config.B2.StrippedComponents
	if len(pathComponents) < strippedComponents {
		log.Println("GenerateB2URL encountered a path too short to handle", pathComponents)
		return nil, errors.New("encountered a path too short to handle")
	}
	// see previous note
	relPath := strings.Join(pathComponents[strippedComponents:], "/")

	// very bad way of unescaping the path separators
	urlRelPath := strings.Replace(url.QueryEscape(relPath), "%2F", "/", -1)
	return url.Parse(config.B2.BaseURL + urlRelPath)
}
