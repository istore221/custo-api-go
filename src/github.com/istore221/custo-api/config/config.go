package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	Environment string
	Server struct {
		Port     string
		LogLevel string
	}
}

func LoadConfiguration(cfgFile string) (*Config, error) {

	c := Config{
		Environment: getEnv("APP_ENV", "development") ,
	}

	bytes, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(bytes, &c); err != nil {
		return nil, err
	}


	return &c, nil
}

func getEnv(key, defaultValue string) string{

	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}