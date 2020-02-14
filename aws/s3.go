package aws

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func NewS3Uploader() (*s3manager.Uploader, error) {

	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "idelic-dev",
		Config: aws.Config{
			Region: aws.String("eu-central-1"),
			Credentials: credentials.NewStaticCredentials("AKIAJBP5562NBHPK23UQ", "6m6H1B44CkZ+wLPNRM1iqsqyi/HhEeKEHmqZ6HGu", ""),
		},
	})

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error creating new AWS session: %s", err))
	} else {
		uploader := s3manager.NewUploader(sess)
		return uploader, nil
	}
}
