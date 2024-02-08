package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// Initialize STS client with explicit IAM user credentials
func createSTSClientWithCredentials(accessKeyID, secretAccessKey, sessionToken string) *sts.Client {
	customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		if service == sts.ServiceID {
			return aws.Endpoint{
				URL:           "https://sts.amazonaws.com",
				SigningRegion: "us-east-1",
			}, nil
		}
		// Fallback to default resolver
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(aws.NewCredentialsCache(aws.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, sessionToken))),
		config.WithEndpointResolver(customResolver),
	)
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	return sts.NewFromConfig(cfg)
}
 // AssumeRole and return the credentials
func assumeRole(stsClient *sts.Client, roleARN, sessionName string) (*sts.Credentials, error) {
    input := &sts.AssumeRoleInput{
        RoleArn:         aws.String(roleARN),
        RoleSessionName: aws.String(sessionName),
    }

    result, err := stsClient.AssumeRole(context.TODO(), input)
    if err != nil {
        return nil, fmt.Errorf("unable to assume role: %w", err)
    }

    return result.Credentials, nil
}

func main() {
    // Example usage
    accessKeyID := "your-access-key-id"
    secretAccessKey := "your-secret-access-key"
    sessionToken := "your-session-token" // Leave empty if not using temporary credentials
    roleARN := "arn:aws:iam::123456789012:role/YourRole"
    sessionName := "YourSessionName"

    stsClient := createSTSClientWithCredentials(accessKeyID, secretAccessKey, sessionToken)
    credentials, err := assumeRole(stsClient, roleARN, sessionName)
    if err != nil {
        fmt.Println("Error assuming role:", err)
        return
    }

    fmt.Println("Assumed Role Credentials:", credentials)
}
