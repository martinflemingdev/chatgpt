package main

import (
	"fmt"
	"strings"
)

func removeStringBetweenColons(input string, from, to int) string {
	components := strings.Split(input, ":")
	if len(components) > to {
		components[from] = "" // Remove the string between the specified positions
	}
	return strings.Join(components, ":")
}

func main() {
	input := "arn:aws:region:account:service:app:branch:build:component:"
	result := removeStringBetweenColons(input, 7, 8)
	fmt.Println(result)
}
