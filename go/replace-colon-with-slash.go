package main

import (
	"fmt"
	"strings"
)

// ReplaceLastColonWithSlash takes a string and replaces the last colon with a forward slash, if the last character is a colon
func ReplaceLastColonWithSlash(input string) string {
	if strings.HasSuffix(input, ":") {
		return strings.TrimSuffix(input, ":") + "/"
	}
	return input
}

func main() {
	p := "arn:aws:syd:Non:cloud:cdc:comp:4:s3:"
	modified := ReplaceLastColonWithSlash(p)
	fmt.Println("Modified string:", modified)
}
