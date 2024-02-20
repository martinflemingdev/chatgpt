package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Define a struct to hold your item's attributes
type Item struct {
	PlanID  string `dynamodbav:"planID"`
	BuildID string `dynamodbav:"buildID"`
	// Add other attributes as needed
}

// QueryDynamo queries a DynamoDB table using a partition key 'planID' and a sort key 'buildID'
func QueryDynamo(ctx context.Context, client *dynamodb.Client, tableName string) (*dynamodb.QueryOutput, error) {
	input := &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		KeyConditionExpression: aws.String("planID = :planVal AND begins_with(buildID, :buildVal)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":planVal":  &types.AttributeValueMemberS{Value: "service:app"},
			":buildVal": &types.AttributeValueMemberS{Value: "branch"},
		},
	}

	result, err := client.Query(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to query items, %w", err)
	}

	return result, nil
}

// UnmarshalQueryResult takes the result from a DynamoDB query and unmarshals it into a slice of Items
func UnmarshalQueryResult(result *dynamodb.QueryOutput) ([]Item, error) {
	var items []Item
	for _, item := range result.Items {
		var i Item
		err := attributevalue.UnmarshalMap(item, &i)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal item, %w", err)
		}
		items = append(items, i)
	}
	return items, nil
}

func main() {
	// Example usage
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic(fmt.Sprintf("Failed to load configuration: %v", err))
	}
	client := dynamodb.NewFromConfig(cfg)

	queryResult, err := QueryDynamo(ctx, client, "YourTableName")
	if err != nil {
		fmt.Println("Error querying DynamoDB:", err)
		return
	}

	items, err := UnmarshalQueryResult(queryResult)
	if err != nil {
		fmt.Println("Error unmarshaling query result:", err)
		return
	}

	fmt.Println("Queried Items:", items)
}
