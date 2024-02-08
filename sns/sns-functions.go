package main

import (
    "context"
    "fmt"
    "log"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/sns"
)

// NewSNSClient creates a new SNS client from the provided AWS config.
func NewSNSClient(cfg aws.Config) *sns.Client {
    return sns.NewFromConfig(cfg)
}

// PublishMessage publishes a message to the specified SNS topic.
func PublishMessage(ctx context.Context, snsClient *sns.Client, topicArn string, message string) error {
    input := &sns.PublishInput{
        Message:  aws.String(message),
        TopicArn: aws.String(topicArn),
    }

    _, err := snsClient.Publish(ctx, input)
    if err != nil {
        return fmt.Errorf("failed to publish message: %w", err)
    }

    return nil
}

func main() {
    cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
        log.Fatalf("unable to load SDK config, %v", err)
    }

    snsClient := NewSNSClient(cfg)

    err = PublishMessage(context.TODO(), snsClient, "arn:aws:sns:region:account-id:topicName", "Hello, SNS!")
    if err != nil {
        log.Fatalf("failed to publish message: %v", err)
    }
}
