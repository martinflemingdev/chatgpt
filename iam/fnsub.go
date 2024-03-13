package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Define a struct that matches the overall structure of your IAM policy document.
type PolicyDocument struct {
	Statement []StatementEntry `json:"Statement"`
}

type StatementEntry struct {
	Resource []interface{} `json:"Resource"`
}

// Define a struct for the special Resource field that uses Fn::Sub
type FnSubResource struct {
	FnSub []interface{} `json:"Fn::Sub"`
}

func main() {
	// Example JSON string
	jsonString := `{
        "Statement": [
            {
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
            }
        ]
    }`

	var doc PolicyDocument
	err := json.Unmarshal([]byte(jsonString), &doc)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Process each statement
	for i, stmt := range doc.Statement {
		for j, res := range stmt.Resource {
			// Attempt to unmarshal the Resource into the FnSubResource struct
			resBytes, err := json.Marshal(res)
			if err != nil {
				fmt.Println("Error marshalling Resource:", err)
				continue
			}

			var fnSubRes FnSubResource
			if err := json.Unmarshal(resBytes, &fnSubRes); err == nil {
				// Successfully unmarshalled into FnSubResource, now substitute
				substitutedResource, err := substituteFnSub(fnSubRes)
				if err != nil {
					fmt.Println("Error substituting Fn::Sub:", err)
					continue
				}
				// Replace the original resource with the substituted value
				doc.Statement[i].Resource[j] = substitutedResource
			}
		}
	}

	// Marshal back to JSON to see the result
	modifiedJSON, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling modified document:", err)
		return
	}

	fmt.Println(string(modifiedJSON))
}

// substituteFnSub performs the substitution defined in the FnSubResource struct.
func substituteFnSub(fnSubRes FnSubResource) (string, error) {
	if len(fnSubRes.FnSub) != 2 {
		return "", fmt.Errorf("Fn::Sub array does not have 2 elements")
	}

	template, ok := fnSubRes.FnSub[0].(string)
	if !ok {
		return "", fmt.Errorf("first element of Fn::Sub is not a string")
	}

	variables, ok := fnSubRes.FnSub[1].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("second element of Fn::Sub is not an object")
	}

	for varName, varValue := range variables {
		varValueStr, ok := varValue.(string)
		if !ok {
			return "", fmt.Errorf("variable value is not a string")
		}
		placeholder := fmt.Sprintf("${%s}", varName)
		template = strings.ReplaceAll(template, placeholder, varValueStr)
	}

	return template, nil
}
