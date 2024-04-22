package main

import (
	"fmt"
	"regexp"
)

func normalizeIAMPolicyResources(policyJSON string) (string, error) {
	// Regex to find "Resource": followed by a quoted string (not an array)
	pattern := regexp.MustCompile(`("Resource":\s*)"([^"\[\]]+)"`)

	// Replacement pattern to convert the found string into an array
	replacement := `${1}["${2}"]`

	// Use regex to replace all occurrences where "Resource" is not an array
	normalizedJSON := pattern.ReplaceAllString(policyJSON, replacement)

	return normalizedJSON, nil
}

func main() {
	serializedPolicy := `
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "s3:ListBucket",
            "Resource": "arn:aws:s3:::example_bucket"
        },
        {
            "Effect": "Allow",
            "Action": "s3:GetObject",
            "Resource": ["arn:aws:s3:::example_bucket/*"]
        }
    ]
}
`

	normalizedPolicy, err := normalizeIAMPolicyResources(serializedPolicy)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Normalized Policy JSON:")
		fmt.Println(normalizedPolicy)
	}
}
