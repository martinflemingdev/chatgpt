// MARSHAL

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Item struct {
	ARN        string `json:"arn"`
	OtherField string `json:"otherField"`
}

func main() {
	itemsMap := map[string]Item{
		"arn1": {ARN: "arn1", OtherField: "value1"},
		"arn2": {ARN: "arn2", OtherField: "value2"},
	}

	jsonData, err := json.Marshal(itemsMap)
	if err != nil {
		log.Fatalf("Error marshalling map to JSON: %v", err)
	}

	fmt.Printf("%s\n", jsonData)
}

// UNMARSHAL

package main

import (
    "encoding/json"
    "fmt"
    "log"
)

type Item struct {
    ARN        string `json:"arn"`
    OtherField string `json:"otherField"`
}

func main() {
    jsonData := []byte(`{
        "arn1": {"arn": "arn1", "otherField": "value1"},
        "arn2": {"arn": "arn2", "otherField": "value2"}
    }`)

    var itemsMap map[string]Item
    err := json.Unmarshal(jsonData, &itemsMap)
    if err != nil {
        log.Fatalf("Error unmarshalling JSON to map: %v", err)
    }

    fmt.Println(itemsMap)
}
