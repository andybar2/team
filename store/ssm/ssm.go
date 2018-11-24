package ssm

import (
	"errors"
	"fmt"
	"path"

	"github.com/aws/aws-sdk-go/aws"
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

// Del deletes a variable
func (s *Store) Del(environment, variable string) error {
	name := s.getParameterName(environment, variable)

	_, err := s.ssmSession.DeleteParameter(&ssm.DeleteParameterInput{
		Name: aws.String(name),
	})

	return err
}

// Print prints all the variables and their values for the given environment
func (s *Store) Print(environment string) error {
	searchPath := s.getPath(environment)
	nextToken := ""

	for {
		input := &ssm.GetParametersByPathInput{
			Path:      aws.String(searchPath),
			Recursive: aws.Bool(true),
		}

		if nextToken != "" {
			input.SetNextToken(nextToken)
		}

		output, err := s.ssmSession.GetParametersByPath(input)
		if err != nil {
			return err
		}

		for _, par := range output.Parameters {
			name := path.Base(aws.StringValue(par.Name))
			value := aws.StringValue(par.Value)

			fmt.Printf("%s=%s\n", name, value)
		}

		if output.NextToken == nil {
			break
		}

		nextToken = aws.StringValue(output.NextToken)
	}

	return nil
}

func (s *Store) getPath(environment string) string {
	return fmt.Sprintf("/%s/%s", s.project, environment)
}

func (s *Store) getParameterName(environment, variable string) string {
	return fmt.Sprintf("%s/%s", s.getPath(environment), variable)
}
