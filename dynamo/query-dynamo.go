package main

import (
    "context"
    "fmt"
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// QueryDynamo queries a DynamoDB table using a partition key 'planID' and a sort key 'buildID'
func QueryDynamo(ctx context.Context, client *dynamodb.Client, tableName string) (*dynamodb.QueryOutput, error) {
    input := &dynamodb.QueryInput{
        TableName: aws.String(tableName),
        KeyConditionExpression: aws.String("planID = :planVal AND begins_with(buildID, :buildVal)"),
        ExpressionAttributeValues: map[string]types.AttributeValue{
            ":planVal": &types.AttributeValueMemberS{Value: "service:app"},
            ":buildVal": &types.AttributeValueMemberS{Value: "branch"},
        },
    }

    result, err := client.Query(ctx, input)
    if err != nil {
        return nil, fmt.Errorf("failed to query items, %w", err)
    }

    return result, nil
}

func main() {
    // Initialize DynamoDB client and context
    // ...

    // Call QueryDynamo with the client, table name
    // ...
}
