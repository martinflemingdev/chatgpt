package main

import (
    "bytes"
    "encoding/json"
    "log"
    "net/http"
)

type MyPayload struct {
    Name    string `json:"name"`
    Age     int    `json:"age"`
    Country string `json:"country"`
}

func main() {
    payload := MyPayload{
        Name:    "John Doe",
        Age:     30,
        Country: "Neverland",
    }

    jsonData, err := json.Marshal(payload)
    if err != nil {
        log.Fatalf("Error marshaling JSON: %v", err)
    }

    reader := bytes.NewReader(jsonData)
    req, err := http.NewRequest("POST", "http://example.com/api", reader)
    if err != nil {
        log.Fatalf("Error creating request: %v", err)
    }

    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatalf("Error sending request: %v", err)
    }
    defer resp.Body.Close()

    // Process the response...
}
