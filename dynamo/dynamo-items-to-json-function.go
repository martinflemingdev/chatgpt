package main

import (
    "encoding/base64"
    "encoding/json"
    "fmt"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func convertDynamoDBItemToJSON(dynamoItem map[string]types.AttributeValue) ([]byte, error) {
    jsonItem := make(map[string]interface{})

    for key, attributeValue := range dynamoItem {
        switch v := attributeValue.(type) {
        case *types.AttributeValueMemberS: // String
            jsonItem[key] = v.Value
        case *types.AttributeValueMemberN: // Number
            jsonItem[key] = v.Value
        case *types.AttributeValueMemberB: // Binary (Base64 Encoded)
            jsonItem[key] = base64.StdEncoding.EncodeToString(v.Value)
        case *types.AttributeValueMemberBOOL: // Boolean
            jsonItem[key] = v.Value
        case *types.AttributeValueMemberNULL: // Null
            jsonItem[key] = v.Value
        case *types.AttributeValueMemberL: // List
            list := make([]interface{}, len(v.Value))
            for i, av := range v.Value {
                var err error
                list[i], err = convertDynamoDBItemToJSON(map[string]types.AttributeValue{"value": av})
                if err != nil {
                    return nil, err
                }
            }
            jsonItem[key] = list
        case *types.AttributeValueMemberM: // Map
            m, err := convertDynamoDBItemToJSON(v.Value)
            if err != nil {
                return nil, err
            }
            jsonItem[key] = m
        case *types.AttributeValueMemberSS: // String Set
            jsonItem[key] = v.Value
        case *types.AttributeValueMemberNS: // Number Set
            jsonItem[key] = v.Value
        case *types.AttributeValueMemberBS: // Binary Set
            bs := make([]string, len(v.Value))
            for i, b := range v.Value {
                bs[i] = base64.StdEncoding.EncodeToString(b)
            }
            jsonItem[key] = bs
        // Add other types as necessary
        }
    }

    return json.Marshal(jsonItem)
}

func main() {
    // Mock DynamoDB item
    dynamoItem := map[string]types.AttributeValue{
        "Id":     &types.AttributeValueMemberS{Value: "123"},
        "Name":   &types.AttributeValueMemberS{Value: "Example"},
        "Count":  &types.AttributeValueMemberN{Value: "10"},
        // Include other types for demonstration
    }

    jsonData, err := convertDynamoDBItemToJSON(dynamoItem)
    if err != nil {
        fmt.Printf("Error converting DynamoDB item to JSON: %s\n", err)
        return
    }

    fmt.Printf("JSON output: %s\n", jsonData)
}
