package ssm

import (
	"errors"
	"fmt"

	"github.com/andybar2/team-env/store"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// Store wraps the logic for storing environment variables in the SSM Parameters Store (AWS)
type Store struct {
	project    string
	ssmSession *ssm.SSM
}

// NewStore returns a new store.IStore that stores environment variables on SSM
func NewStore(proj string) (store.IStore, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, errors.New("could not create aws session")
	}

	return &Store{
		project:    proj,
		ssmSession: ssm.New(sess),
	}, nil
}

// Set sets a variable value
func (s *Store) Set(environment, variable, value string) error {
	name := s.getParameterName(environment, variable)

	_, err := s.ssmSession.PutParameter(&ssm.PutParameterInput{
		Name:      aws.String(name),
		Overwrite: aws.Bool(true),
		Type:      aws.String("String"),
		Value:     aws.String(value),
	})

	return err
}

// Get gets a variable value
func (s *Store) Get(environment, variable string) (string, error) {
	name := s.getParameterName(environment, variable)

	output, err := s.ssmSession.GetParameter(&ssm.GetParameterInput{
		Name: aws.String(name),
	})
	if err != nil {
		return "", err
	}

	return aws.StringValue(output.Parameter.Value), nil
}

func (s *Store) getParameterName(environment, variable string) string {
	return fmt.Sprintf("/%s/%s/%s", s.project, environment, variable)
}
