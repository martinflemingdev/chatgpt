package main

import (
	"fmt"
	"regexp"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// CapitalizeBetweenColons takes a string and capitalizes the substring between the 3rd and 4th colon
func CapitalizeBetweenColons(input string) (string, error) {
	re := regexp.MustCompile(`^([^:]*:[^:]*:[^:]*:)([^:]*)(:.*)$`)
	matches := re.FindStringSubmatch(input)

	if len(matches) != 4 {
		return "", fmt.Errorf("input string does not match the expected format")
	}

	caser := cases.Title(language.English)
	return matches[1] + caser.String(matches[2]) + matches[3], nil
}

func main() {
	p := "arn:aws:syd:non:cloud:cdc:comp:4:s3:"
	modified, err := CapitalizeBetweenColons(p)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Modified string:", modified)
	}
}
