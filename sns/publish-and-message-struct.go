package main

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/sns"
)

// Message structure to be sent to SNS
type SNSMessage struct {
    Bucket         string `json:"bucket"`
    Key            string `json:"key"`
    Prefix         string `json:"prefix"`
    AdditionalInfo string `json:"additionalInfo"`
}

// PublishMessageToSNSTopic publishes a message to an SNS topic
func PublishMessageToSNSTopic(ctx context.Context, snsClient *sns.Client, topicArn, bucket, key, prefix, additionalInfo string) error {
    // Construct the message
    msg := SNSMessage{
        Bucket:         bucket,
        Key:            key,
        Prefix:         prefix,
        AdditionalInfo: additionalInfo,
    }

    // Marshal the message into JSON
    msgBytes, err := json.Marshal(msg)
    if err != nil {
        return fmt.Errorf("failed to marshal message: %w", err)
    }

    // Publish the message to the SNS topic
    _, err = snsClient.Publish(ctx, &sns.PublishInput{
        Message:  aws.String(string(msgBytes)),
        TopicArn: aws.String(topicArn),
    })

    if err != nil {
        return fmt.Errorf("failed to publish message to SNS topic: %w", err)
    }

    return nil
}

func main() {
    // Context to control cancellations and timeouts
    ctx := context.TODO()

    // Load the AWS configuration
    cfg, err := config.LoadDefaultConfig(ctx)
    if err != nil {
        fmt.Println("error loading AWS configuration:", err)
        return
    }

    // Create an SNS client
    snsClient := sns.NewFromConfig(cfg)

    // Example usage
    topicArn := "your-topic-arn" // Replace with your SNS topic ARN
    bucket := "your-bucket-name"
    key := "your-object-key"
    prefix := "optional-prefix"
    additionalInfo := "Any other relevant information"

    err = PublishMessageToSNSTopic(ctx, snsClient, topicArn, bucket, key, prefix, additionalInfo)
    if err != nil {
        fmt.Println("error publishing message to SNS topic:", err)
    } else {
        fmt.Println("Message published successfully")
    }
}
