package minio

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func InitS3Client(endpoint, accessKeyID, secretAccessKey string) (s3Client *s3.S3 ){
	s := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKeyID,secretAccessKey,""),
		Region: aws.String("us-west-2"),
	}))
	s3Config := aws.NewConfig().WithEndpoint(endpoint).WithS3ForcePathStyle(true)

	s3Client = s3.New(s, s3Config)

	return
}
