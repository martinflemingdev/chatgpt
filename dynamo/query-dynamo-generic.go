package main

import (
    "context"
    "fmt"
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// QueryDynamoDB is a generic function to query a DynamoDB table.
func QueryDynamoDB(ctx context.Context, svc *dynamodb.Client, tableName string, keyConditionExpression string, filterExpression string, expressionAttributeValues map[string]types.AttributeValue) (*dynamodb.QueryOutput, error) {
    // Construct the query input
    input := &dynamodb.QueryInput{
        TableName:                 aws.String(tableName),
        KeyConditionExpression:    aws.String(keyConditionExpression),
        FilterExpression:          aws.String(filterExpression),
        ExpressionAttributeValues: expressionAttributeValues,
    }

    // Execute the query
    result, err := svc.Query(ctx, input)
    if err != nil {
        return nil, fmt.Errorf("failed to query items: %w", err)
    }

    return result, nil
}

func main() {
    // Initialize a DynamoDB client
    cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
        panic("configuration error, " + err.Error())
    }

    svc := dynamodb.NewFromConfig(cfg)

    // Example usage
    tableName := "YourTableName"
    keyConditionExpression := "YourPartitionKey = :v1"
    filterExpression := "YourFilterExpression"
    expressionAttributeValues := map[string]types.AttributeValue{
        ":v1": &types.AttributeValueMemberS{Value: "YourKeyValue"},
    }

    ctx := context.TODO()
    result, err := QueryDynamoDB(ctx, svc, tableName, keyConditionExpression, filterExpression, expressionAttributeValues)
    if err != nil {
        fmt.Println("Query failed:", err)
        return
    }

    fmt.Println("Query succeeded:", result)
}
