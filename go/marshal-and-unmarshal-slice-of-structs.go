package main

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

type MyStruct struct {
	Field1 string
	Field2 int
	// Add other fields as necessary
}

func main() {
	// Example slice of structs
	myStructs := []MyStruct{
		{Field1: "Value1", Field2: 1},
		{Field1: "Value2", Field2: 2},
	}

	// Marshal the slice of structs to JSON
	jsonData, err := json.Marshal(myStructs)
	if err != nil {
		log.Fatalf("Error marshaling data: %v", err)
	}

	// Create an AWS session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create an S3 service client
	s3svc := s3.New(sess)

	// Define the bucket and object key
	bucket := "your-bucket-name"
	key := "your-object-key.json"

	// Upload the JSON data to S3
	_, err = s3svc.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        aws.ReadSeekCloser(strings.NewReader(string(jsonData))),
		ContentType: aws.String("application/json"),
	})

	if err != nil {
		log.Fatalf("Failed to upload data to %s/%s, %s", bucket, key, err)
	}
}

////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////

package main

import (
    "encoding/json"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "io/ioutil"
    "log"
)

type MyStruct struct {
    Field1 string
    Field2 int
    // Add other fields as necessary
}

func main() {
    // Create an AWS session
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    // Create an S3 service client
    s3svc := s3.New(sess)

    // Define the bucket and object key
    bucket := "your-bucket-name"
    key := "your-object-key.json"

    // Download the JSON data from S3
    obj, err := s3svc.GetObject(&s3.GetObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(key),
    })
    if err != nil {
        log.Fatalf("Failed to download object, %s", err)
    }
    defer obj.Body.Close()

    // Read the JSON data
    jsonData, err := ioutil.ReadAll(obj.Body)
    if err != nil {
        log.Fatalf("Failed to read object body, %s", err)
    }

    // Unmarshal the JSON data into a slice of structs
    var myStructs []MyStruct
    err = json.Unmarshal(jsonData, &myStructs)
    if err != nil {
        log.Fatalf("Error unmarshaling data: %v", err)
    }

    // Now `myStructs` contains the slice of structs, and you can use it as needed
}
