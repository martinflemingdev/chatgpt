package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
)

func main() {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Unable to load SDK config, %v", err)
	}

	// Create an Organizations client
	svc := organizations.NewFromConfig(cfg)

	// Use the client to list accounts
	result, err := svc.ListAccounts(context.TODO(), &organizations.ListAccountsInput{})
	if err != nil {
		log.Fatalf("Error listing accounts, %v", err)
	}

	// Print the account details
	fmt.Println("Accounts under the AWS Organization:")
	for _, account := range result.Accounts {
		fmt.Printf("Account Name: %s, Account Email: %s, Account ID: %s\n",
			aws.ToString(account.Name), aws.ToString(account.Email), aws.ToString(account.Id))
	}
}
