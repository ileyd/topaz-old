package utils

import "net/url"

func GenerateB2URL(relPath string) string {
	baseUrl := "https://f001.backblazeb2.com/file/testing-content/"
	urlRelPath := url.PathEscape(relPath)

	return baseUrl + urlRelPath
}
