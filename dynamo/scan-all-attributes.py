import boto3
from botocore.exceptions import NoCredentialsError

def assume_role_with_profile(profile_name, role_arn, session_name):
    try:
        session = boto3.Session(profile_name=profile_name)
        sts_client = session.client('sts')
        assumed_role_object = sts_client.assume_role(
            RoleArn=role_arn,
            RoleSessionName=session_name
        )
        credentials = assumed_role_object['Credentials']
        return credentials
    except NoCredentialsError:
        print("Credentials not found for profile:", profile_name)
        raise

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

# Replace with your AWS profile name
profile_name = 'your_profile_name'

# Replace with your role ARN and a session name
role_arn = 'your_role_arn'
session_name = 'your_session_name'

# Replace with your table name
table_name = 'your_table_name'

# Assume the role using the AWS profile
try:
    credentials = assume_role_with_profile(profile_name, role_arn, session_name)

    # Get DynamoDB resource using the assumed role credentials
    dynamodb_resource = get_dynamodb_resource(credentials)

    # Get all attributes
    all_attributes = get_all_attributes(table_name, dynamodb_resource)
    print(all_attributes)
except NoCredentialsError:
    print("Failed to assume role.")
