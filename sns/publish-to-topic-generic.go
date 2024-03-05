package main

import (
    "context"
    "fmt"
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/sns"
)

// PublishMessage publishes a message to an SNS topic and returns the message ID.
func PublishMessage(ctx context.Context, topicArn, message string) (string, error) {
    // Load the AWS default configuration.
    cfg, err := config.LoadDefaultConfig(ctx)
    if err != nil {
        return "", fmt.Errorf("error loading AWS configuration: %w", err)
    }

    // Create an SNS client.
    snsClient := sns.NewFromConfig(cfg)

    // Create the input for the Publish call.
    input := &sns.PublishInput{
        Message:  aws.String(message),
        TopicArn: aws.String(topicArn),
    }

    // Publish the message.
    result, err := snsClient.Publish(ctx, input)
    if err != nil {
        return "", fmt.Errorf("error publishing message to SNS topic: %w", err)
    }

    return *result.MessageId, nil
}

func main() {
    ctx := context.TODO()
    topicArn := "arn:aws:sns:us-west-2:123456789012:MyTopic"
    message := "Hello, world!"

    messageId, err := PublishMessage(ctx, topicArn, message)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("Message published successfully. Message ID: %s\n", messageId)
}
