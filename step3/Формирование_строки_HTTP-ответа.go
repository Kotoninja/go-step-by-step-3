package main

import (
	"fmt"
)

func BuildHTTPResponse(statusLine, headers, body string) string {
	result := fmt.Sprintf("%s\r\n%s\r\n%s", statusLine, headers, body)
	return result
}
