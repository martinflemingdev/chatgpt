package main

import (
	"fmt"
	"regexp"
	"strings"
)

// CapitalizeBetweenColons takes a string and capitalizes the substring between the 3rd and 4th colon
func CapitalizeBetweenColons(input string) (string, error) {
	re := regexp.MustCompile(`^([^:]*:[^:]*:[^:]*:)([^:]*)(:.*)$`)
	matches := re.FindStringSubmatch(input)

	if len(matches) != 4 {
		return "", fmt.Errorf("input string does not match the expected format")
	}

	return matches[1] + strings.Title(matches[2]) + matches[3], nil
}

func main() {
	p := "arn:aws:region:non:cloud:cdc:comp:4:s3:"
	modified, err := CapitalizeBetweenColons(p)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Modified string:", modified)
	}
}
