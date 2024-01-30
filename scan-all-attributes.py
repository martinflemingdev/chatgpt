import boto3

def assume_role(role_arn, session_name, aws_access_key_id, aws_secret_access_key):
    sts_client = boto3.client(
        'sts',
        aws_access_key_id=aws_access_key_id,
        aws_secret_access_key=aws_secret_access_key
    )
    assumed_role_object = sts_client.assume_role(
        RoleArn=role_arn,
        RoleSessionName=session_name
    )
    credentials = assumed_role_object['Credentials']
    return credentials

def get_dynamodb_resource(credentials):
    return boto3.resource(
        'dynamodb',
        aws_access_key_id=credentials['AccessKeyId'],
        aws_secret_access_key=credentials['SecretAccessKey'],
        aws_session_token=credentials['SessionToken'],
    )

def get_all_attributes(table_name, dynamodb_resource):
    table = dynamodb_resource.Table(table_name)

    attributes = set()
    scan = table.scan()

    for item in scan['Items']:
        attributes.update(item.keys())

    while 'LastEvaluatedKey' in scan:
        scan = table.scan(ExclusiveStartKey=scan['LastEvaluatedKey'])
        for item in scan['Items']:
            attributes.update(item.keys())

    return attributes

# Replace with your static IAM user credentials
aws_access_key_id = 'your_access_key_id'
aws_secret_access_key = 'your_secret_access_key'

# Replace with your role ARN and a session name
role_arn = 'your_role_arn'
session_name = 'your_session_name'

# Replace with your table name
table_name = 'your_table_name'

# Assume the role using static IAM user credentials
credentials = assume_role(role_arn, session_name, aws_access_key_id, aws_secret_access_key)

# Get DynamoDB resource using the assumed role credentials
dynamodb_resource = get_dynamodb_resource(credentials)

# Get all attributes
all_attributes = get_all_attributes(table_name, dynamodb_resource)
print(all_attributes)
