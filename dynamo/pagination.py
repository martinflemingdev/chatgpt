import boto3

# Create a DynamoDB client
dynamodb = boto3.client('dynamodb')

# Define the table name and the query parameters
table_name = 'YourTableName'
query_params = {
    'TableName': table_name,
    'KeyConditionExpression': 'YourKeyConditionExpression',
    # Add other query parameters as needed
}

# Initialize an empty list to store all the items
all_items = []

# Query the table and handle pagination
while True:
    response = dynamodb.query(**query_params)
    all_items.extend(response.get('Items', []))

    # Check if there are more items to fetch
    if 'LastEvaluatedKey' in response:
        query_params['ExclusiveStartKey'] = response['LastEvaluatedKey']
    else:
        break

# Now, all_items contains all the items from the query
print(all_items)
