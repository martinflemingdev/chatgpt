package main

import (
    "encoding/json"
    "fmt"
    "log"

    "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// MarshalDynamoDBItemsToJson takes DynamoDB items and marshals them into a JSON []byte
func MarshalDynamoDBItemsToJson(items []map[string]types.AttributeValue) ([]byte, error) {
    jsonData, err := json.Marshal(items)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal items to JSON: %w", err)
    }
    return jsonData, nil
}

func main() {
    // Example usage
    // Assuming you have a 'result' variable which is a dynamodb.QueryOutput
    // result := ... 

    jsonData, err := MarshalDynamoDBItemsToJson(result.Items)
    if err != nil {
        log.Fatalf("Error marshaling DynamoDB items to JSON: %v", err)
    }

    // Now jsonData contains the JSON representation of the DynamoDB items
    // This can be passed to your PutObjectOnS3 function
}
