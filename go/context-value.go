package main

import (
    "context"
    "fmt"
)

// Define a key type to avoid key collisions in context
type keyType int

// Define a key of the custom type
const key keyType = iota

func main() {
    // Create a context
    ctx := context.Background()

    // Add a value to the context
    val := "example value"
    ctx = context.WithValue(ctx, key, val)

    // Retrieve the value from the context
    if retrievedVal, ok := ctx.Value(key).(string); ok {
        fmt.Println("Retrieved value:", retrievedVal)
    } else {
        fmt.Println("Failed to retrieve value")
    }
}
