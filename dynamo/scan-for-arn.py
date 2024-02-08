import boto3

# Initialize a DynamoDB client
dynamodb = boto3.resource('dynamodb')
table = dynamodb.Table('YourTableName')

# Scan the table
response = table.scan()

# Check each item for the "arn" substring in any attribute
for item in response['Items']:
    for key, value in item.items():
        if isinstance(value, str) and "arn" in value:
            print(f'Found "arn" in item with key {key}: {value}')
