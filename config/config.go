package config

import "github.com/caarlos0/env/v8"

type Environment struct {
	AwsIdentifier string `env:"AWS_ID"`
	AwsSecret     string `env:"AWS_SECRET"`
}

var Env Environment

func init() {
	if err := env.Parse(&Env); err != nil {
		panic(err)
	}
}
