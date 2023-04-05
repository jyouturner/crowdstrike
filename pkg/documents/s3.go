package main

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func generateSignedURL(bucketName, objectName string) (string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("your-region"),
		Credentials: credentials.NewStaticCredentials("your-access-key", "your-secret-key", ""),
	})
	if err != nil {
		return "", err
	}

	s3svc := s3.New(sess)

	// Set the expiry time for the signed URL
	expires := time.Now().Add(15 * time.Minute)

	// Generate the pre-signed URL
	req, _ := s3svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(objectName),
		ContentType: aws.String("application/octet-stream"),
	})
	signedURL, err := req.Presign(expires)

	if err != nil {
		return "", err
	}

	return signedURL, nil
}
