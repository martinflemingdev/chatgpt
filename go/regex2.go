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

    // Variable to keep track of how much the string has shifted
    shift := 0

    for _, match := range matches {
        // Extracting the full match and the groups
        fullMatch := input[match[0]:match[1]]
        template := input[match[4]:match[5]]
        variableName := input[match[6]:match[7]]
        variableValue := input[match[8]:match[9]]

        // Perform the substitution
        substituted := strings.Replace(template, fmt.Sprintf("${%s}", variableName), variableValue, -1)

        // Replace the original Fn::Sub section with the substituted value in the input string
        before := input[:match[0]+shift]
        after := input[match[1]+shift:]
        input = before + fmt.Sprintf(`"%s"`, substituted) + after

        // Calculate the new shift based on the difference in length between the original and substituted
        shiftDiff := len(fmt.Sprintf(`"%s"`, substituted)) - len(fullMatch)

        // Update the shift for the next iteration
        shift += shiftDiff
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
