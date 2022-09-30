package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Define OpenAWSS3Connection function to open AWS S3 connection return error input is access key, secret key and region name
func OpenAWSS3Connection(AccessKey string, SecretKey string, RegionName string) (ClientObj *s3.S3, err error) {
	//use aws-sdk-go  to open AWS S3 connection
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)
	// Create S3 service client
	svc := s3.New(sess)

	return svc, nil

}
