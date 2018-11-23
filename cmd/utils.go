package cmd

import (
	"errors"
	"os"

	"github.com/andybar2/team-env/config"
	"github.com/andybar2/team-env/store"
	"github.com/andybar2/team-env/store/ssm"
)

func setupStore() (store.IStore, error) {
	appConfig, err := config.ReadConfig("team-env.json")
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
