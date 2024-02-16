import boto3

def get_all_accounts(org_client):
    accounts = []
    paginator = org_client.get_paginator('list_accounts')

    for page in paginator.paginate():
        accounts.extend(page['Accounts'])

    return accounts

# Example usage
client = boto3.client('organizations')
all_accounts = get_all_accounts(client)
print(all_accounts)
