package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func uploadItem(sess *session.Session) {
	f, err := os.Open("downloads-and-uploads/my-file.txt")
	if err != nil {
		log.Fatal("could not open file")
	}

	defer f.Close()

	uploader := s3manager.NewUploader(sess)

	result, err := uploader.Upload(&s3manager.UploadInput{
		ACL:    aws.String("public-read"),
		Bucket: aws.String("go-aws-s3"),
		Key:    aws.String("my-file.txt"),
		Body:   f,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Upload Result: %+v\n", result)
}

func listItems(sess *session.Session) {
	svc := s3.New(sess)

	response, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String("go-aws-s3"),
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, item := range response.Contents {
		log.Printf("Name: %s\n", *item.Key)
		log.Printf("Size: %d\n", *item.Size)
	}
}

func downloadItems(sess *session.Session) {
	file, err := os.Create("downloads-and-uploads/downloaded.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer file.Close()

	downloader := s3manager.NewDownloader(sess)

	_, err = downloader.Download(file, &s3.GetObjectInput{Bucket: aws.String("go-aws-s3"), Key: aws.String("my-file.txt")})

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("succefully downloaded file")
}

func deleteItme(sess *session.Session) {
	svc := s3.New(sess)

	input := &s3.DeleteObjectInput{
		Bucket: aws.String("go-aws-s3"),
		Key:    aws.String("my-file.txt"),
	}

	result, err := svc.DeleteObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				log.Fatal(aerr.Error())
			}
		} else {
			log.Fatal(err.Error())
		}
	}

	log.Printf("Result: %+v\n ", result)
}

func main() {
	fmt.Println("Downloads and Uploads")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		log.Fatal("Could not get session")
	}

	uploadItem(sess)
	listItems(sess)
}
