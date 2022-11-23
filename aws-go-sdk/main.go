package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {

	// Get bucket name from command line arguments
	bucketName := flag.String("b", "", "The name of the bucket")
	region := flag.String("r", "", "The name of the region")
	flag.Parse()

	if *bucketName == "" {
		fmt.Println("You must supply the name of a bucket (-b BUCKET)")
		return
	}

	if *region == "" {
		fmt.Println("You must supply the region of a bucket (-r REGION)")
		return
	}

	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(*region))
	if err != nil {
		log.Fatalf("Failed to load configuration, %v", err)
	}

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)

	// Get the first page of results for ListObjectsV2 for a bucket
	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(*bucketName),
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("First Page results:")
	for _, object := range output.Contents {
		log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
	}
}
