package config

import (
	"log"

	"github.com/Netflix/go-env"
)

type Environment struct {
	Port       int64 `env:"PORT,default=8080"`
	Production bool  `env:"PRODUCTION,default=false"`

	BucketName string `env:"BUCKET_NAME,default=previews"`
	UrlPrefix  string `env:"URL_PREFIX"`

	S3Endpoint        string `env:"S3_ENDPOINT"`
	S3AccessKeyID     string `env:"S3_ACCESS_KEY_ID"`
	S3SecretAccessKey string `env:"S3_SECRET_ACCESS_KEY"`
	S3UseSSL          bool   `env:"S3_USE_SSL,default=false"`

	Extras env.EnvSet
}

func GetConfig() (*Environment, error) {
	var environment Environment

	es, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	environment.Extras = es

	return &environment, nil
}
