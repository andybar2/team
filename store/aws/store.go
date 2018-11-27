package aws

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// Store wraps the logic for storing environment variables in the SSM Parameters Store (AWS)
type Store struct {
	project      string
	region       string
	ssmSession   *ssm.SSM
	s3Session    *s3.S3
	s3Downloader *s3manager.Downloader
	s3Uploader   *s3manager.Uploader
	iamSession   *iam.IAM
}

// NewStore returns a new Store that stores environment variables on SSM
func NewStore(project, region string) (*Store, error) {
	sess, err := session.NewSession(aws.NewConfig().WithRegion(region))
	if err != nil {
		return nil, errors.New("could not create aws session")
	}

	return &Store{
		project:      project,
		region:       region,
		ssmSession:   ssm.New(sess),
		s3Session:    s3.New(sess),
		s3Downloader: s3manager.NewDownloader(sess),
		s3Uploader:   s3manager.NewUploader(sess),
		iamSession:   iam.New(sess),
	}, nil
}
