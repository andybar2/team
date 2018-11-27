package aws

import (
	"fmt"
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// EnvSet sets the value of an enviroment variable for the given stage
func (s *Store) EnvSet(stage, name, value string) error {
	parameter := s.getParameterName(stage, name)

	_, err := s.ssmSession.PutParameter(&ssm.PutParameterInput{
		Name:      aws.String(parameter),
		Overwrite: aws.Bool(true),
		Type:      aws.String("String"),
		Value:     aws.String(value),
	})

	return err
}

// EnvGet gets the current value of an enviroment variable for the given stage
func (s *Store) EnvGet(stage, name string) (string, error) {
	parameter := s.getParameterName(stage, name)

	output, err := s.ssmSession.GetParameter(&ssm.GetParameterInput{
		Name: aws.String(parameter),
	})
	if err != nil {
		return "", err
	}

	return aws.StringValue(output.Parameter.Value), nil
}

// EnvDelete deletes an environment variable from the given stage
func (s *Store) EnvDelete(stage, name string) error {
	parameter := s.getParameterName(stage, name)

	_, err := s.ssmSession.DeleteParameter(&ssm.DeleteParameterInput{
		Name: aws.String(parameter),
	})

	return err
}

// EnvPrint prints all the environment variables and their values for the given stage
func (s *Store) EnvPrint(stage string) error {
	searchPath := s.getParametersPath(stage)
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

func (s *Store) getParametersPath(stage string) string {
	return fmt.Sprintf("/team/%s/%s", s.project, stage)
}

func (s *Store) getParameterName(stage, name string) string {
	return fmt.Sprintf("%s/%s", s.getParametersPath(stage), name)
}
