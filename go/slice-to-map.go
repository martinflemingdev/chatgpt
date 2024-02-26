package main

import (
    "fmt"
)

// Define the Item struct
type Item struct {
    // Assuming 'arn' is of type string and other fields as needed
    arn string
    // other fields...
}

func main() {
    // Example slice of Item structs
    items := []Item{
        {arn: "arn1", /* other field values */},
        {arn: "arn2", /* other field values */},
        // Add more Items as needed
    }

    // Convert the slice to a map
    itemsMap := make(map[string]Item)
    for _, item := range items {
        itemsMap[item.arn] = item
    }

    // Print the map to verify
    fmt.Println(itemsMap)

    // To access a specific item by its ARN
    specificItem, exists := itemsMap["arn1"]
    if exists {
        fmt.Println("Found item:", specificItem)
    } else {
        fmt.Println("Item not found")
    }
}
