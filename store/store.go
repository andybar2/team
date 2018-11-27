package store

import (
	"errors"
	"os"

	"github.com/andybar2/team/config"
	"github.com/andybar2/team/store/aws"
)

// IStore is the interface that should be implemented by the specific stores
type IStore interface {
	// EnvSet sets the value of an enviroment variable for the given stage
	EnvSet(stage, name, value string) error

	// EnvGet gets the current value of an enviroment variable for the given stage
	EnvGet(stage, name string) (string, error)

	// EnvDelete deletes an environment variable from the given stage
	EnvDelete(stage, name string) error

	// EnvPrint prints all the environment variables and their values for the given stage
	EnvPrint(stage string) error

	// FileUpload uploads a file to the given stage
	FileUpload(stage, name, filePath string) error

	// FileDownload downloads a file from the given stage
	FileDownload(stage, name, filePath string) error

	// TODO:
	// FileDownloadAll
	// FileDel
}

// New reads the app configuration from the given file and sets up the corresponding store
func New() (IStore, error) {
	appConfig, err := config.ReadConfig()
	if err != nil {
		return nil, err
	}

	if appConfig.Project == "" {
		return nil, errors.New("invalid project name")
	}

	switch appConfig.Store {
	case "aws":
		if appConfig.AWSProfile == "" {
			return nil, errors.New("invalid aws profile")
		}

		if appConfig.AWSRegion == "" {
			return nil, errors.New("invalid aws region")
		}

		os.Setenv("AWS_PROFILE", appConfig.AWSProfile)
		os.Setenv("AWS_REGION", appConfig.AWSRegion)

		return aws.NewStore(appConfig.Project, appConfig.AWSRegion)
	default:
		return nil, errors.New("unsupported store")
	}
}
