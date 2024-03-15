package main

import (
    "fmt"
    "strconv"
)

// Define your struct
type MyStruct struct {
    ActiveIds []string
    ReleaseId int64
}

// Function to check if the releaseId is present in activeIds
func (s *MyStruct) IsReleaseIdActive() bool {
    // Convert the int64 releaseId to a string
    releaseIdStr := strconv.FormatInt(s.ReleaseId, 10)

    // Iterate over the slice of strings
    for _, id := range s.ActiveIds {
        if id == releaseIdStr {
            return true // Found a match
        }
    }

    return false // No match found
}

func main() {
    // Example usage
    exampleStruct := MyStruct{
        ActiveIds: []string{"11", "12", "13"},
        ReleaseId: 12,
    }

    if exampleStruct.IsReleaseIdActive() {
        fmt.Println("ReleaseId is active.")
    } else {
        fmt.Println("ReleaseId is not active.")
    }
}
