package aws

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// Store wraps the logic for storing environment variables in the SSM Parameters Store (AWS)
type Store struct {
	project    string
	ssmSession *ssm.SSM
}

// NewStore returns a new Store that stores environment variables on SSM
func NewStore(proj string) (*Store, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, errors.New("could not create aws session")
	}

	return &Store{
		project:    proj,
		ssmSession: ssm.New(sess),
	}, nil
}
