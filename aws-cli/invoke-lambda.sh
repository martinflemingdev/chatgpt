aws lambda invoke \
--function-name YourLambdaFunctionName \
--invocation-type Event \
--payload file://payload.json \
outputfile.txt

# Base64 encode the payload
PAYLOAD=$(base64 payload.json)

# Invoke the Lambda function with the base64-encoded payload
aws lambda invoke \
--function-name YourLambdaFunctionName \
--invocation-type Event \
--payload $PAYLOAD \
outputfile.txt

aws lambda invoke \
--function-name YourLambdaFunctionName \
--invocation-type Event \
--payload '{"Records":[{"messageId":"example-message-id","receiptHandle":"example-receipt-handle","body":"Hello from SQS!","attributes":{"ApproximateReceiveCount":"1","SentTimestamp":"1523232000000","SenderId":"123456789012","ApproximateFirstReceiveTimestamp":"1523232000001"},"messageAttributes":{},"md5OfBody":"example-md5-of-body","eventSource":"aws:sqs","eventSourceARN":"arn:aws:sqs:us-east-1:123456789012:MyQueue","awsRegion":"us-east-1"}]}' \
outputfile.txt
