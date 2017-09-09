package utils

import "net/url"

func GenerateB2URL(relPath string) string {
	baseURL := "https://f001.backblazeb2.com/file/testing-content/"
	urlRelPath := url.PathEscape(relPath)

	return baseURL + urlRelPath
}
