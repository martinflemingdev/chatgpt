package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Define the Item struct
type Item struct {
	PartitionKey string `dynamodbav:"PartitionKey"`
	SortKey      string `dynamodbav:"SortKey"`
	ARN          string `dynamodbav:"ARN"`
	// Add other attributes as needed
}

// UnmarshalDynamoDBItem takes a DynamoDB item and unmarshals it into an Item struct
func UnmarshalDynamoDBItem(dynamoDBItem map[string]types.AttributeValue) (*Item, error) {
	var item Item
	err := attributevalue.UnmarshalMap(dynamoDBItem, &item)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal DynamoDB item: %w", err)
	}
	return &item, nil
}

func main() {
	// Example usage
	dynamoDBItem := map[string]types.AttributeValue{
		"PartitionKey": &types.AttributeValueMemberS{Value: "ExamplePartitionKey"},
		"SortKey":      &types.AttributeValueMemberS{Value: "ExampleSortKey"},
		"ARN":          &types.AttributeValueMemberS{Value: "ExampleARN"},
	}

	item, err := UnmarshalDynamoDBItem(dynamoDBItem)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Unmarshaled Item: %+v\n", item)
	}
}
