package main

import (
    "encoding/base64"
    "encoding/json"
    "reflect"
    "testing"

    "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// TestConvertDynamoDBItemToJSON tests the conversion of DynamoDB items to JSON.
func TestConvertDynamoDBItemToJSON(t *testing.T) {
    // Mock DynamoDB item
    dynamoItem := map[string]types.AttributeValue{
        "StringField": &types.AttributeValueMemberS{Value: "StringValue"},
        "NumberField": &types.AttributeValueMemberN{Value: "123"},
        "BinaryField": &types.AttributeValueMemberB{Value: []byte("BinaryData")},
        "BoolField":   &types.AttributeValueMemberBOOL{Value: true},
        "NullField":   &types.AttributeValueMemberNULL{Value: true},
        // Add other fields as necessary for your test
    }

    // Expected JSON representation
    expectedJSON := map[string]interface{}{
        "StringField": "StringValue",
        "NumberField": "123",
        "BinaryField": base64.StdEncoding.EncodeToString([]byte("BinaryData")),
        "BoolField":   true,
        "NullField":   true,
        // Add other fields to match your test
    }

    // Convert DynamoDB item to JSON
    jsonData, err := convertDynamoDBItemToJSON(dynamoItem)
    if err != nil {
        t.Fatalf("convertDynamoDBItemToJSON() error = %v", err)
    }

    // Unmarshal the JSON data into a map for comparison
    var resultJSON map[string]interface{}
    if err := json.Unmarshal(jsonData, &resultJSON); err != nil {
        t.Fatalf("Error unmarshaling JSON data: %v", err)
    }

    // Compare the expected and result JSON
    if !reflect.DeepEqual(expectedJSON, resultJSON) {
        t.Errorf("convertDynamoDBItemToJSON() = %v, want %v", resultJSON, expectedJSON)
    }
}

// Include your convertDynamoDBItemToJSON function here
