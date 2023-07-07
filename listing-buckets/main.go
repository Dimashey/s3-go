package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	svc := s3.New(sess)

	result, err := svc.ListBuckets(nil)
	if err != nil {
		log.Println(err)
		log.Fatal("Error listing buckets")
	}

	for _, bucket := range result.Buckets {
		log.Printf("Bucket: %s\n", aws.StringValue(bucket.Name))
	}
}
