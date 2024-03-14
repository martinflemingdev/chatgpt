package main

import (
    "fmt"
    "regexp"
    "strings"
)

func substituteVariables(input string) (string, error) {
    // Regex to find ${VariableName} placeholders
    placeholderRegex := regexp.MustCompile(`\$\{([^}]+)\}`)
    // Find all placeholders
    placeholders := placeholderRegex.FindAllStringSubmatch(input, -1)

    for _, placeholder := range placeholders {
        // Extract the variable name from the placeholder
        variableName := placeholder[1]

        // Build a regex pattern to find "VariableName": "VariableValue"
        variablePattern := fmt.Sprintf(`"%s":\s*"([^"]+)"`, variableName)
        variableRegex := regexp.MustCompile(variablePattern)

        // Find the variable value
        matches := variableRegex.FindStringSubmatch(input)
        if len(matches) < 2 {
            return "", fmt.Errorf("variable %s not found", variableName)
        }
        variableValue := matches[1]

        // Replace the placeholder with the variable value
        input = strings.Replace(input, placeholder[0], variableValue, -1)
    }

    return input, nil
}

func main() {
    input := `{
  "Resource": [
    {
      "Fn::Sub": [
        "arn:aws:s3:::${BucketName}",
        {
          "BucketName": "Mybucket"
        }
      ]
    }
  ]
}`

    result, err := substituteVariables(input)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Result:", result)
}
