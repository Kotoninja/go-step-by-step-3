package main

import (
	"fmt"
	"strings"
)

func MakeCurlCommand(method, url, headers, body string) string {
	result := "curl "

	if method != "GET" {
		result += fmt.Sprintf("-X %s ", method)
	}

	if headers != "" {
		for _, head := range strings.Split(headers, "\n")[:len(strings.Split(headers, "\n"))-1] {
			result += fmt.Sprintf("-H '%s' ", head)
		}
	}

	if body != "" {
		result += fmt.Sprintf("--data %s", body)
	}

	result += url
	return strings.TrimSpace(result)
}

// func main() {
// 	fmt.Println(MakeCurlCommand(
// 		"POST",
// 		"https://example.com/api/users",
// 		"Content-Type: application/json\nAuthorization: Bearer abc123\n",
// 		`{"name":"John Doe","email":"johndoe@example.com","password":"123456"}`))
// }
