package main

import (
    "context"
    "fmt"
    "log"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/sqs"
)

// sendToSQS sends a message to the specified SQS queue.
func sendToSQS(queueURL string, messageBody string) error {
    // Load the Shared AWS Configuration (~/.aws/config)
    cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
        return fmt.Errorf("error loading AWS configuration: %w", err)
    }

    // Create an SQS client
    client := sqs.NewFromConfig(cfg)

    // Create the message
    input := &sqs.SendMessageInput{
        MessageBody: aws.String(messageBody),
        QueueUrl:    aws.String(queueURL),
    }

    // Send the message
    _, err = client.SendMessage(context.TODO(), input)
    if err != nil {
        return fmt.Errorf("error sending message to SQS: %w", err)
    }

    log.Printf("Message sent to SQS queue: %s", queueURL)
    return nil
}

func main() {
    queueURL := "https://sqs.us-east-1.amazonaws.com/123456789012/MyQueue" // Replace with your Queue URL
    messageBody := "Hello from Lambda!"

    if err := sendToSQS(queueURL, messageBody); err != nil {
        log.Fatalf("Failed to send message: %s", err)
    }
}
