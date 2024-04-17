package main

import (
    "context"
    "fmt"
    "log"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/organizations"
    "github.com/aws/aws-sdk-go-v2/service/organizations/types"
)

// NewOrganizationsClient creates and returns an AWS Organizations client.
func NewOrganizationsClient(cfg aws.Config) *organizations.Client {
    return organizations.NewFromConfig(cfg)
}

// ListAccountsWithPagination retrieves all accounts and returns them in a map where the key is the account name and the value is the account ID.
func ListAccountsWithPagination(ctx context.Context, svc *organizations.Client) (map[string]string, error) {
    paginator := organizations.NewListAccountsPaginator(svc, &organizations.ListAccountsInput{})
    accountsMap := make(map[string]string)

    for paginator.HasMorePages() {
        page, err := paginator.NextPage(ctx)
        if err != nil {
            return nil, err // Return the error to the caller
        }

        for _, account := range page.Accounts {
            accountsMap[aws.ToString(account.Name)] = aws.ToString(account.Id)
        }
    }

    return accountsMap, nil
}

func main() {
    ctx := context.TODO()

    // Load the AWS Configuration
    cfg, err := config.LoadDefaultConfig(ctx)
    if err != nil {
        log.Fatalf("Unable to load SDK config, %v", err)
    }

    // Create an Organizations client
    svc := NewOrganizationsClient(cfg)

    // List accounts with pagination
    accounts, err := ListAccountsWithPagination(ctx, svc)
    if err != nil {
        log.Fatalf("Failed to list accounts: %v", err)
    }

    // Print the account details
    fmt.Println("Accounts under the AWS Organization:")
    for name, id := range accounts {
        fmt.Printf("Account Name: %s, Account ID: %s\n", name, id)
    }
}
