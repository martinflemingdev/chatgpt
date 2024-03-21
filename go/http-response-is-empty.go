package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    url := "https://api.example.com/data"
    if isEmptyJSON(url) {
        fmt.Println("The response is an empty JSON object.")
    } else {
        fmt.Println("The response contains data.")
    }
}

// isEmptyJSON makes a GET request to the specified URL and checks if the response is an empty JSON object.
func isEmptyJSON(url string) bool {
    // Make the GET request
    resp, err := http.Get(url)
    if err != nil {
        fmt.Printf("Error making the request: %s\n", err)
        return false
    }
    defer resp.Body.Close()

    // Read the body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("Error reading the response body: %s\n", err)
        return false
    }

    // Unmarshal into a map to check if it's empty
    var data map[string]interface{}
    err = json.Unmarshal(body, &data)
    if err != nil {
        fmt.Printf("Error unmarshaling the response: %s\n", err)
        return false
    }

    // Check if the map is empty
    return len(data) == 0
}
