package aws

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// FileUpload uploads a file to the given stage
func (s *Store) FileUpload(stage, filePath string) error {
	// get bucket name
	bucketName, err := s.getBucketName()
	if err != nil {
		return err
	}

	// setup bucket
	if err = s.setupBucket(bucketName); err != nil {
		return err
	}

	// open file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// read file
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		return errors.New("file path can't be a directory")
	}

	size := fileInfo.Size()
	buffer := make([]byte, size)

	_, err = file.Read(buffer)
	if err != nil {
		return err
	}

	// upload file
	fileName := path.Base(filePath)
	fileKey := s.getFileKey(stage, fileName)

	_, err = s.s3Uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(fileKey),
		ACL:         aws.String("private"),
		Body:        bytes.NewReader(buffer),
		ContentType: aws.String(http.DetectContentType(buffer)),
	})

	return err
}

// FileDownload downloads a file from the given stage
func (s *Store) FileDownload(stage, filePath string) error {
	// get bucket name
	bucketName, err := s.getBucketName()
	if err != nil {
		return err
	}

	// setup bucket
	if err = s.setupBucket(bucketName); err != nil {
		return err
	}

	// create file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// download file
	fileName := path.Base(filePath)
	fileKey := s.getFileKey(stage, fileName)

	_, err = s.s3Downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileKey),
	})

	return err
}

func (s *Store) getBucketName() (string, error) {
	getUserOutput, err := s.iamSession.GetUser(&iam.GetUserInput{})
	if err != nil {
		return "", err
	}

	userARN := aws.StringValue(getUserOutput.User.Arn)
	accountID := strings.Split(userARN, ":")[4]
	bucketName := fmt.Sprintf("team-%s-%s", accountID, s.region)

	return bucketName, nil
}

func (s *Store) getFileKey(stage, name string) string {
	return fmt.Sprintf("%s/%s/%s", s.project, stage, name)
}

func (s *Store) setupBucket(bucketName string) error {
	_, err := s.s3Session.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		if strings.Contains(err.Error(), "BucketAlreadyOwnedByYou") {
			return nil
		}

		return err
	}

	return nil
}
