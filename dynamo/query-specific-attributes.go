package main

import (
    "context"
    "fmt"
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// buildProjectionExpression builds the ProjectionExpression and ExpressionAttributeNames for a DynamoDB query
func buildProjectionExpression(attributes []string) (string, map[string]string) {
    projExpression := ""
    expressionAttributeNames := make(map[string]string)
    for i, attr := range attributes {
        if i > 0 {
            projExpression += ", "
        }
        placeholder := fmt.Sprintf("#attr%d", i)
        projExpression += placeholder
        expressionAttributeNames[placeholder] = attr
    }
    return projExpression, expressionAttributeNames
}

// QueryDynamoDBForAttributesWithBeginsWith queries a DynamoDB table for specific attributes based on partition key and a sort key that begins with a specified value
func QueryDynamoDBForAttributesWithBeginsWith(ctx context.Context, dynamoClient *dynamodb.Client, tableName string, partitionKeyName string, partitionKeyValue string, sortKeyName string, sortKeyPrefix string, attributes []string) ([]map[string]types.AttributeValue, error) {
    keyConditionExpression := fmt.Sprintf("%s = :pkval AND begins_with(%s, :skprefix)", partitionKeyName, sortKeyName)
    expressionAttributeValues := map[string]types.AttributeValue{
        ":pkval":    &types.AttributeValueMemberS{Value: partitionKeyValue},
        ":skprefix": &types.AttributeValueMemberS{Value: sortKeyPrefix},
    }

    projExpression, expressionAttributeNames := buildProjectionExpression(attributes)

    input := &dynamodb.QueryInput{
        TableName:                 aws.String(tableName),
        KeyConditionExpression:    aws.String(keyConditionExpression),
        ExpressionAttributeValues: expressionAttributeValues,
        ProjectionExpression:      aws.String(projExpression),
        ExpressionAttributeNames:  expressionAttributeNames,
    }

    result, err := dynamoClient.Query(ctx, input)
    if err != nil {
        return nil, err
    }

    return result.Items, nil
}

func main() {
    ctx := context.TODO()

    cfg, err := config.LoadDefaultConfig(ctx)
    if err != nil {
        fmt.Println("error loading AWS config:", err)
        return
    }

    dynamoClient := dynamodb.NewFromConfig(cfg)

    tableName := "YourTableName"
    partitionKeyName := "YourPartitionKeyName"
    partitionKeyValue := "YourPartitionKeyValue"
    sortKeyName := "YourSortKeyName"
    sortKeyPrefix := "YourSortKeyPrefix"
    attributes := []string{"atr1", "atr2", "atr3"}

    items, err := QueryDynamoDBForAttributesWithBeginsWith(ctx, dynamoClient, tableName, partitionKeyName, partitionKeyValue, sortKeyName, sortKeyPrefix, attributes)
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