package main

import (
    "fmt"
    "strings"
    "unicode"
)

// Capitalize capitalizes the first letter of the given string.
func Capitalize(s string) string {
    if s == "" {
        return ""
    }
    return strings.ToUpper(string(s[0])) + s[1:]
}

// CapitalizeFirstOnly capitalizes only the first letter of the entire string,
// ensuring the rest of the string is in lowercase.
func CapitalizeFirstOnly(s string) string {
    if s == "" {
        return ""
    }
    // Convert the first character to uppercase and the rest to lowercase.
    return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}

// CapitalizeFirstRune capitalizes the first letter of the string, supporting Unicode.
// This is a more correct approach if you're dealing with multi-byte characters.
func CapitalizeFirstRune(s string) string {
    for _, v := range s {
        // Convert the first rune to uppercase and concatenate with the rest of the string.
        return string(unicode.ToUpper(v)) + s[len(string(v)):]
    }
    return ""
}

func main() {
    testString := "hello world"
    fmt.Println(Capitalize(testString))        // Output: Hello world
    fmt.Println(CapitalizeFirstOnly(testString)) // Output: Hello world
    fmt.Println(CapitalizeFirstRune("привет мир")) // Output: Привет мир (demonstrating Unicode support)
}
