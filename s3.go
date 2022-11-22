package main

import (
	"bytes"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Config struct {
	AccessKeyID     string
	AccessKeySecret string
	Region          string
	BucketName      string
	Endpoint        string
}

func CreateSession(cfg *S3Config) *session.Session {
	sess := session.Must(session.NewSession(
		&aws.Config{
			Endpoint:         aws.String(cfg.Endpoint),
			Region:           aws.String(cfg.Region),
			DisableSSL:       aws.Bool(true),
			S3ForcePathStyle: aws.Bool(true),
			Credentials: credentials.NewStaticCredentials(
				cfg.AccessKeyID,
				cfg.AccessKeySecret,
				"",
			),
		},
	))
	return sess
}

func UploadObject(fileName string, body []byte, cfg *S3Config) (string, error) {
	session := CreateSession(cfg)
	uploader := s3manager.NewUploader(session)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(cfg.BucketName),
		Key:         aws.String(fileName),
		Body:        bytes.NewReader(body),
		ACL:         aws.String("public-read"),
		ContentType: aws.String(http.DetectContentType(body)),
	})
	if err != nil {
		return "", err
	}
	return result.Location, nil
}
