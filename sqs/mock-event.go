package main

import (
    "context"
    "fmt"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
)

// Your Lambda handler
func HandleRequest(ctx context.Context, sqsEvent events.SQSEvent) error {
    for _, message := range sqsEvent.Records {
        fmt.Printf("The message body: %s\n", message.Body)
        // Process message...
    }
    return nil
}

func main() {
    lambda.Start(HandleRequest)
}

// Test function
func mockSQSEvent() events.SQSEvent {
    return events.SQSEvent{
        Records: []events.SQSMessage{
            {
                MessageId:     "1",
                ReceiptHandle: "abc",
                Body:          "Message 1 body",
                Attributes: map[string]string{
                    "ApproximateReceiveCount":          "1",
                    "SentTimestamp":                    "1523232000000",
                    "SenderId":                         "123456789012",
                    "ApproximateFirstReceiveTimestamp": "1523232000001",
                },
                MessageAttributes: map[string]events.SQSMessageAttribute{
                    "Attribute1": {
                        StringValue:      aws.String("Value1"),
                        BinaryValue:      nil,
                        StringListValues: nil,
                        BinaryListValues: nil,
                        DataType:         "String",
                    },
                },
                MD5OfBody:            "7b270e59b47ff90a553787216d55d91d",
                EventSource:          "aws:sqs",
                EventSourceARN:       "arn:aws:sqs:us-east-1:123456789012:MyQueue",
                AwsRegion:            "us-east-1",
            },
            // Add more messages as needed
        },
    }
}
