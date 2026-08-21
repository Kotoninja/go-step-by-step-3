package main

import (
	"strconv"
	"strings"
)

func ParseHTTPStatus(statusLine string) (code int, reason string) {
	statusLine = statusLine[strings.Index(statusLine, " ")+1:]

	separater := strings.Index(statusLine, " ")

	number, _ := strconv.Atoi(statusLine[:separater])
	return number, statusLine[separater+1:]
}
