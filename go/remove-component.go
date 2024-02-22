package main

import (
    "fmt"
    "strings"
)

func removeLastComponent(arn string) string {
    components := strings.Split(arn, ":")
    if len(components) > 1 {
        components = components[:len(components)-1] // Remove the last component
    }
    return strings.Join(components, ":")
}

func main() {
    arn := "arn:aws:region:account:service:app:branch:build:component:"
    modifiedArn := removeLastComponent(arn)
    fmt.Println(modifiedArn)
}
