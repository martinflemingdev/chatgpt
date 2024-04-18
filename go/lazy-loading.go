// accountsmgr/accountsmgr.go

package accountsmgr

import (
    "sync"
)

var (
    accountMap map[string]string
    loadOnce   sync.Once
    mu         sync.Mutex
)

// Returns the accounts map, loading it if necessary
func GetAccountMap() map[string]string {
    mu.Lock()
    defer mu.Unlock()

    loadOnce.Do(func() {
        accountMap = loadAccountMapFromAWS()
    })
    return accountMap
}

// Actually loads the account map from AWS or another data source
func loadAccountMapFromAWS() map[string]string {
    // Here you would implement the actual loading logic.
    // This could involve calling AWS Organizations API and processing the response.
    // For example:
    // response, err := awsOrganizationsClient.ListAccounts(...)
    // if err != nil {
    //     log.Fatal("Failed to load accounts: ", err)
    // }
    // mapData := make(map[string]string)
    // for _, account := range response.Accounts {
    //     mapData[account.Name] = account.Id
    // }
    // return mapData
    return make(map[string]string) // Placeholder for actual implementation
}

package component

import (
    "fmt"
    "github.com/yourrepo/accountsmgr"
)

func ProcessSomething() {
    accounts := accountsmgr.GetAccountMap()
    // Use the accounts map for whatever processing is needed
    fmt.Println("Loaded accounts map: ", accounts)
}
