import boto3
import json

def lambda_handler(event, context):
    # Parameters
    assumed_role_arn = 'arn:aws:iam::<ACCOUNT_ID>:role/<ROLE_NAME>'
    dynamodb_table = '<DYNAMODB_TABLE_NAME>'
    partition_key = '<PARTITION_KEY>'
    partition_value = '<PARTITION_VALUE>'
    sort_key = '<SORT_KEY>'
    sort_key_prefix = '<SORT_KEY_PREFIX>'
    s3_bucket = '<S3_BUCKET_NAME>'
    s3_key = '<S3_OBJECT_KEY>'

    # Assume role
    sts_client = boto3.client('sts')
    assumed_role_object = sts_client.assume_role(
        RoleArn=assumed_role_arn,
        RoleSessionName='AssumeRoleSession'
    )
    credentials = assumed_role_object['Credentials']

    # Create a DynamoDB client with the assumed role
    dynamodb = boto3.client(
        'dynamodb',
        aws_access_key_id=credentials['AccessKeyId'],
        aws_secret_access_key=credentials['SecretAccessKey'],
        aws_session_token=credentials['SessionToken'],
    )

    # Query DynamoDB
    response = dynamodb.query(
        TableName=dynamodb_table,
        KeyConditionExpression=f"{partition_key} = :partition_value AND begins_with({sort_key}, :sort_key_prefix)",
        ExpressionAttributeValues={
            ':partition_value': {'S': partition_value},
            ':sort_key_prefix': {'S': sort_key_prefix}
        }
    )

    # Process DynamoDB items
    items = response.get('Items', [])

    # Create an S3 client with the lambda's role
    s3 = boto3.client('s3')

    # Write to S3
    s3.put_object(
        Bucket=s3_bucket,
        Key=s3_key,
        Body=json.dumps(items)
    )

    return {
        'statusCode': 200,
        'body': json.dumps('Items written to S3')
    }
