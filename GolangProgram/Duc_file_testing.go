package main

import (
	"bytes"
	"fmt"
	"os"

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

// define ReadJPEGFile function to read JPEG file return byte array input is file path
func ReadJPEGFile(FilePath string) ([]byte, error) {
	//Check FilePath format is JPEG file
	if FilePath[len(FilePath)-4:] != ".jpg" {
		return nil, fmt.Errorf("File is not JPEG file")
	}
	//open file
	file, err := os.Open(FilePath)
	if err != nil {
		return nil, err
	}
	//read file
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	//get file size
	var size int64 = fileInfo.Size()
	//make a buffer
	buffer := make([]byte, size)
	//read file into buffer
	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}
	//Check JPEGFortmat
	if !bytes.HasPrefix(buffer, []byte("\xff\xd8")) {
		return nil, fmt.Errorf("Not a JPEG file")
	}
	//return byte array
	return buffer, nil
}

// Define UpdateJPEGFileToS3 function to update JPEG file to S3 return error input is AWS S3 connection, bucket name, file path, file name,File content
func UpdateJPEGFileToS3(ClientObj *s3.S3, BucketName string, FileName string, FileContent []byte) (err error) {
	//use aws-sdk-go  to update JPEG file to S3
	_, err = ClientObj.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(BucketName),
		Key:    aws.String(FileName),
		Body:   bytes.NewReader(FileContent),
	})
	if err != nil {
		return err
	}
	return nil
}
