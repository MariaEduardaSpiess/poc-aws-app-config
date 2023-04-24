package feature_flags

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/MariaEduardaSpiess/aws_app_config_test/config"
	"github.com/MariaEduardaSpiess/aws_app_config_test/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/appconfigdata"
)

type FeatureXPTO struct {
	Enabled          bool   `json:"enabled"`
	RegexCpfsRollout string `json:"regex_cpfs_rollout"`
}

type Configuration struct {
	FeatureXpto FeatureXPTO `json:"feature_xpto"`
}

var (
	onlyOnce                  sync.Once
	initialConfigurationToken *string
	client                    *appconfigdata.AppConfigData
)

func GetFeatureFlags() (Configuration, error) {

	onlyOnce.Do(func() {
		sess, err := session.NewSession(&aws.Config{
			Region:      aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials(config.Env.AwsIdentifier, config.Env.AwsSecret, ""),
		})
		if err != nil {
			panic(err)
		}

		client = appconfigdata.New(sess)

		sessionInput := appconfigdata.StartConfigurationSessionInput{
			ApplicationIdentifier:                aws.String("b3h1yng"),
			ConfigurationProfileIdentifier:       aws.String("w4crioj"),
			EnvironmentIdentifier:                aws.String("p7ckf5q"),
			RequiredMinimumPollIntervalInSeconds: aws.Int64(15),
		}
		startConfigurationSessionOutput, err := client.StartConfigurationSession(&sessionInput)
		if err != nil {
			panic(err)
		}

		initialConfigurationToken = startConfigurationSessionOutput.InitialConfigurationToken
	})

	start := time.Now()

	input := appconfigdata.GetLatestConfigurationInput{
		ConfigurationToken: initialConfigurationToken,
	}
	result, err := client.GetLatestConfiguration(&input)
	if err != nil {
		panic(err)
	}

	initialConfigurationToken = result.NextPollConfigurationToken

	elapsed := time.Since(start)
	logger.Log(fmt.Sprint(elapsed.Milliseconds()))

	var configuration Configuration
	if len(result.Configuration) > 0 {
		if err := json.Unmarshal(result.Configuration, &configuration); err != nil {
			panic(err)
		}
	}

	return configuration, nil
}
