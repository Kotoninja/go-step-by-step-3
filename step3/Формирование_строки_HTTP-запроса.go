package main

import (
	"fmt"
)

func BuildHTTPRequest(method, url, host, headers, body string) string {
	result := fmt.Sprintf("%s %s HTTP/1.1\r\nHost: %s\r\n%s\r\n%s", method, url, host, headers, body)
	return result
}
