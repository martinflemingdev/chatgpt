// accountsmgr/accountsmgr.go

package accountsmgr

import (
    "sync"
    "errors"
)

var (
    accountMap map[string]string
    loadOnce   sync.Once
    mapErr     error // to store the error state
    mu         sync.Mutex
)

// GetAccountMap returns the accounts map, loading it if necessary. Also returns an error if loading fails.
func GetAccountMap() (map[string]string, error) {
    mu.Lock()
    defer mu.Unlock()

    if accountMap == nil && mapErr == nil {
        loadOnce.Do(func() {
            accountMap, mapErr = loadAccountMapFromAWS()
        })
    }
    return accountMap, mapErr
}


// Actually loads the account map from AWS or another data source
func loadAccountMapFromAWS() (map[string]string, error) {
    // Example AWS Organizations API call that might fail
    // response, err := awsOrganizationsClient.ListAccounts(...)
    // if err != nil {
        // return nil, err
    // }

    // mapData := make(map[string]string)
    // for _, account := range response.Accounts {
    //     mapData[account.Name] = account.Id
    // }
    // return mapData, nil
    return nil, errors.New("fetching account map failed") // Placeholder for actual implementation
}

package component

import (
    "fmt"
    "github.com/yourrepo/accountsmgr"
)

func ProcessSomething() {
    accounts, err := accountsmgr.GetAccountMap()
    if err != nil {
        fmt.Println("Error loading accounts map:", err)
        return
    }
    // Use the accounts map for whatever processing is needed
    fmt.Println("Loaded accounts map: ", accounts)
}
