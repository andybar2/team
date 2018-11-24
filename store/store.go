package store

import (
	"errors"
	"os"

	"github.com/andybar2/team-env/config"
	"github.com/andybar2/team-env/store/ssm"
)

// IStore is the interface implemented by the specific variables stores
type IStore interface {
	Set(environment, variable, value string) error
	Get(environment, variable string) (string, error)
	Del(environment, variable string) error
	Print(environment string) error
}

// New reads the app configuration from the given file and sets up the corresponding store
func New(configFile string) (IStore, error) {
	appConfig, err := config.ReadConfig(configFile)
	if err != nil {
		return nil, err
	}

	if appConfig.Project == "" {
		return nil, errors.New("invalid project name")
	}

	switch appConfig.Store {
	case "ssm":
		if appConfig.AWSProfile == "" {
			return nil, errors.New("invalid aws profile")
		}

		if appConfig.AWSRegion == "" {
			return nil, errors.New("invalid aws region")
		}

		os.Setenv("AWS_PROFILE", appConfig.AWSProfile)
		os.Setenv("AWS_REGION", appConfig.AWSRegion)

		return ssm.NewStore(appConfig.Project)
	default:
		return nil, errors.New("unsupported store")
	}
}
