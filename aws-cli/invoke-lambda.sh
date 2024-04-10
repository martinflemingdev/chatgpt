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