func QueryDynamo(ctx context.Context, client *dynamodb.Client, tableName string) ([]map[string]types.AttributeValue, error) {
    var allItems []map[string]types.AttributeValue

    input := &dynamodb.QueryInput{
        TableName: aws.String(tableName),
        KeyConditionExpression: aws.String("planID = :planVal AND begins_with(buildID, :buildVal)"),
        ExpressionAttributeValues: map[string]types.AttributeValue{
            ":planVal": &types.AttributeValueMemberS{Value: "service:app"},
            ":buildVal": &types.AttributeValueMemberS{Value: "branch"},
        },
    }

    for {
        result, err := client.Query(ctx, input)
        if err != nil {
            return nil, fmt.Errorf("failed to query items, %w", err)
        }

        allItems = append(allItems, result.Items...)

        if result.LastEvaluatedKey == nil {
            break
        }

        input.ExclusiveStartKey = result.LastEvaluatedKey
    }

    return allItems, nil
}
