aws dynamodb query \
    --profile your-profile-name \
    --region ap-southeast-2 \
    --table-name your-table-name \
    --key-condition-expression "partitionKey = :planID and begins_with(sortKey, :buildID)" \
    --expression-attribute-values '{":planID": {"S": "cloud:arturocdc"}, ":buildID": {"S": "iam"}}' \
    --query "Items[*]"
