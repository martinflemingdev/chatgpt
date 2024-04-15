import (
    "testing"
    "reflect"
    "github.com/aws/aws-lambda-go/events"
)

func TestParseSNSEntity(t *testing.T) {
    tests := []struct {
        name        string
        input       []byte
        wantEntity  events.SNSEntity
        wantErr     bool
    }{
        {
            name: "valid input",
            input: []byte(`{"Type":"Notification","Message":"Hello"}`),
            wantEntity: events.SNSEntity{Type: "Notification", Message: "Hello"},
            wantErr: false,
        },
        {
            name: "invalid input",
            input: []byte(`{"Type":"Notification", "Message":}`),
            wantEntity: events.SNSEntity{},
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := parseSNSEntity(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("parseSNSEntity() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.wantEntity) {
                t.Errorf("parseSNSEntity() got = %v, want %v", got, tt.wantEntity)
            }
        })
    }
}
