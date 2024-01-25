items := []map[string]dynamotypes.AttributeValue{
    {
        "StringField": &dynamotypes.AttributeValueMemberS{Value: "StringValue"},
        "NumberField": &dynamotypes.AttributeValueMemberN{Value: "123"}, // DynamoDB numbers are strings
        "FloatField": &dynamotypes.AttributeValueMemberN{Value: "123.45"}, // Representing a float as a string
        "MapField": &dynamotypes.AttributeValueMemberM{
            Value: map[string]dynamotypes.AttributeValue{
                "NestedString": &dynamotypes.AttributeValueMemberS{Value: "NestedValue"},
                "NestedNumber": &dynamotypes.AttributeValueMemberN{Value: "456"},
            },
        },
        "ListField": &dynamotypes.AttributeValueMemberL{
            Value: []dynamotypes.AttributeValue{
                &dynamotypes.AttributeValueMemberS{Value: "ListString"},
                &dynamotypes.AttributeValueMemberN{Value: "789"},
            },
        },
        "StringSetField": &dynamotypes.AttributeValueMemberSS{Value: []string{"str1", "str2"}},
        "NumberSetField": &dynamotypes.AttributeValueMemberNS{Value: []string{"100", "200"}},
        "BinaryField": &dynamotypes.AttributeValueMemberB{Value: []byte("BinaryData")},
        "BoolField": &dynamotypes.AttributeValueMemberBOOL{Value: true},
        "NullField": &dynamotypes.AttributeValueMemberNULL{Value: true},
    },
    // Additional items can be added here
}

wantJSON := []byte(`[
    {
        "StringField": {
            "S": "StringValue"
        },
        "NumberField": {
            "N": "123"
        },
        "FloatField": {
            "N": "123.45"
        },
        "MapField": {
            "M": {
                "NestedString": {
                    "S": "NestedValue"
                },
                "NestedNumber": {
                    "N": "456"
                }
            }
        },
        "ListField": {
            "L": [
                {
                    "S": "ListString"
                },
                {
                    "N": "789"
                }
            ]
        },
        "StringSetField": {
            "SS": [
                "str1",
                "str2"
            ]
        },
        "NumberSetField": {
            "NS": [
                "100",
                "200"
            ]
        },
        "BinaryField": {
            "B": "QmluYXJ5RGF0YQ=="
        },
        "BoolField": {
            "BOOL": true
        },
        "NullField": {
            "NULL": true
        }
    }
]`)
