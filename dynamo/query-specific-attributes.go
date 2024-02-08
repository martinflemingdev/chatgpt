package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// QueryDynamoDBForAttributes queries a DynamoDB table for specific attributes based on partition and sort key
func QueryDynamoDBForAttributes(ctx context.Context, dynamoClient *dynamodb.Client, tableName string, partitionKeyName string, partitionKeyValue string, sortKeyName string, sortKeyValue string, attributes []string) ([]map[string]types.AttributeValue, error) {
	// Construct the query input
	keyConditionExpression := fmt.Sprintf("%s = :pkval AND %s = :skval", partitionKeyName, sortKeyName)
	expressionAttributeValues := map[string]types.AttributeValue{
		":pkval": &types.AttributeValueMemberS{Value: partitionKeyValue},
		":skval": &types.AttributeValueMemberS{Value: sortKeyValue},
	}

	projExpression := ""
	for i, attr := range attributes {
		if i > 0 {
			projExpression += ", "
		}
		projExpression += fmt.Sprintf("#attr%d", i)
		expressionAttributeValues[fmt.Sprintf(":attr%dval", i)] = &types.AttributeValueMemberS{Value: attr}
	}
	expressionAttributeNames := make(map[string]string)
	for i, attr := range attributes {
		expressionAttributeNames[fmt.Sprintf("#attr%d", i)] = attr
	}

	input := &dynamodb.QueryInput{
		TableName:                 aws.String(tableName),
		KeyConditionExpression:    aws.String(keyConditionExpression),
		ExpressionAttributeValues: expressionAttributeValues,
		ProjectionExpression:      aws.String(projExpression),
		ExpressionAttributeNames:  expressionAttributeNames,
	}

	// Execute the query
	result, err := dynamoClient.Query(ctx, input)
	if err != nil {
		return nil, err
	}

	return result.Items, nil
}

func main() {
	ctx := context.TODO()

	// Load the AWS default config
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		fmt.Println("error loading AWS config:", err)
		return
	}

	// Create a DynamoDB client
	dynamoClient := dynamodb.NewFromConfig(cfg)

	// Example usage
	tableName := "YourTableName"
	partitionKeyName := "YourPartitionKeyName"
	partitionKeyValue := "YourPartitionKeyValue"
	sortKeyName := "YourSortKeyName"
	sortKeyValue := "YourSortKeyValue"
	attributes := []string{"atr1", "atr2", "atr3"} // Attributes you are interested in

	items, err := QueryDynamoDBForAttributes(ctx, dynamoClient, tableName, partitionKeyName, partitionKeyValue, sortKeyName, sortKeyValue, attributes)
	if err != nil {
		fmt.Println("Query failed:", err)
		return
	}

	fmt.Println("Query succeeded:", items)
}

// Important Points:
// The ProjectionExpression is used to specify the attributes you want to get. However, to avoid conflicts with DynamoDB reserved words, 
// it's often paired with ExpressionAttributeNames to map attribute names.
// This example uses the KeyConditionExpression to filter items based on partition and sort key. Adjust the logic as necessary for your specific use case, especially if your query conditions differ.
// Ensure you replace placeholders like YourTableName, YourPartitionKeyName, etc., with actual values from your DynamoDB setup.
// This function provides a basic template to start querying your DynamoDB table for specific attributes. Depending on your actual data model and access patterns, you might need to adjust the query parameters or handle pagination if the response exceeds the 1 MB limit.