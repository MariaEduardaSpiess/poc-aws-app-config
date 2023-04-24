package config

import (
	"fmt"

	"github.com/Netflix/go-env"
	"github.com/subosito/gotenv"
)

type Environment struct {
	AwsIdentifier string `env:"AWS_ID"`
	AwsSecret     string `env:"AWS_SECRET"`
}

var Env Environment

func init() {
	gotenv.Load()

	_, err := env.UnmarshalFromEnviron(&Env)
	if err != nil {
		fmt.Print(err.Error())
	}
}
