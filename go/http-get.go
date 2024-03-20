package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "net/url"
)

func main() {
    baseURL := "http://example.com/api"
    paramName := "paramName"
    paramValue := "paramValue"

    parsedURL := PrepareQueryParameters(baseURL, paramName, paramValue)
    req := CreateHTTPRequest(parsedURL)
    resp := SendHTTPRequest(req)
    uuid, err := ExtractUUIDFromResponse(resp)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("Uuid: %s\n", uuid)
}

func PrepareQueryParameters(baseURL, paramName, paramValue string) string {
    parsedURL, err := url.Parse(baseURL)
    if err != nil {
        fmt.Printf("Error parsing URL: %v\n", err)
        return ""
    }

    params := url.Values{}
    params.Add(paramName, paramValue)
    parsedURL.RawQuery = params.Encode()

    return parsedURL.String()
}

func CreateHTTPRequest(parsedURL string) *http.Request {
    req, err := http.NewRequest("GET", parsedURL, nil)
    if err != nil {
        fmt.Printf("Error creating request: %v\n", err)
        return nil
    }
    // Optionally, set any headers here
    // req.Header.Set("Key", "Value")
    return req
}

func SendHTTPRequest(req *http.Request) *http.Response {
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Printf("Error sending request: %v\n", err)
        return nil
    }
    return resp
}

func ExtractUUIDFromResponse(resp *http.Response) (string, error) {
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", fmt.Errorf("error reading response body: %v", err)
    }

    var result map[string]interface{}
    if err := json.Unmarshal(body, &result); err != nil {
        return "", fmt.Errorf("error parsing JSON response: %v", err)
    }

    uuid, ok := result["Uuid"].(string)
    if !ok {
        return "", fmt.Errorf("Uuid field is missing or not a string")
    }

    return uuid, nil
}
