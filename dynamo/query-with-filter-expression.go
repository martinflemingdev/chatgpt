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
	PartitionKey string `dynamodbav:"PartitionKey"`
	SortKey      string `dynamodbav:"SortKey"`
	ARN          string `dynamodbav:"ARN"`
	// Add other attributes as needed
}

func main() {
	// Load the AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2")) // Change to your region
	if err != nil {
		panic(fmt.Sprintf("Failed to load configuration: %v", err))
	}

	// Create a DynamoDB client
	svc := dynamodb.NewFromConfig(cfg)

	// Define your query input
	input := &dynamodb.QueryInput{
		TableName: aws.String("YourTableName"),
		KeyConditionExpression: aws.String("PartitionKey = :partitionKey AND begins_with(SortKey, :sortKeyPrefix)"),
		FilterExpression:       aws.String("ARN = :arnValue"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":partitionKey": &types.AttributeValueMemberS{Value: "YourPartitionKeyValue"},
			":sortKeyPrefix": &types.AttributeValueMemberS{Value: "YourSortKeyPrefix"},
			":arnValue":      &types.AttributeValueMemberS{Value: "YourARNValue"},
		},
	}

	// Execute the query
	resp, err := svc.Query(context.TODO(), input)
	if err != nil {
		panic(fmt.Sprintf("Failed to query items: %v", err))
	}

	// Process the query results
	for _, item := range resp.Items {
		var i Item
		err = attributevalue.UnmarshalMap(item, &i)
		if err != nil {
			panic(fmt.Sprintf("Failed to unmarshal item: %v", err))
		}
		fmt.Printf("Item: %+v\n", i)
	}
}
