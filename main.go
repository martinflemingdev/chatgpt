package main

import (
    "fmt"
    "regexp"
    "strings"
)

func transformInput(input string) (string, error) {
    // Regex to match the "Fn::Sub" sections including the variable mappings
    fnSubRegex := regexp.MustCompile(`(\{\s*"Fn::Sub":\s*\[\s*"([^"]+)",\s*\{\s*"([^"]+)":\s*"([^"]+)"\s*\}\s*\]\s*\})`)
    
    // Find all matches and their indexes
    matches := fnSubRegex.FindAllStringSubmatchIndex(input, -1)
    if len(matches) == 0 {
        return "", fmt.Errorf("no Fn::Sub structures found")
    }

    // Process matches in reverse to avoid index shifting issues
    for i := len(matches) - 1; i >= 0; i-- {
        match := matches[i]
        // Extracting the template, variable name, and variable value
        template := input[match[4]:match[5]]
        variableName := input[match[6]:match[7]]
        variableValue := input[match[8]:match[9]]

        // Perform the substitution
        substituted := strings.Replace(template, fmt.Sprintf("${%s}", variableName), variableValue, -1)

        // Replace the original Fn::Sub section with the substituted value in the input string
        input = input[:match[0]] + fmt.Sprintf(`"%s"`, substituted) + input[match[1]:]
    }

    return input, nil
}

func main() {
    input := `{
  "Resources": [
    {
      "Fn::Sub": [
        "arn:aws:s3:::${BucketName}",
        {
          "BucketName": "Mybucket1"
        }
      ]
    },
    {
      "Fn::Sub": [
        "arn:aws:s3:::${AnotherBucketName}",
        {
          "AnotherBucketName": "Mybucket2"
        }
      ]
    }
  ]
}`

    output, err := transformInput(input)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Transformed Output:", output)
}
