package config

import (
	"github.com/go-errors/errors"
	"github.com/spf13/viper"
)

type (
	AppConfig struct {
		HTTP     HTTP     `mapstructure:"http"`
		Postgres Postgres `mapstructure:"postgres"`
	}

	HTTP struct {
		Port string `mapstructure:"port"`
	}

	Postgres struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
		SSLMode  string `mapstructure:"sslMode"`
	}
)

// LoadConfig - initializing app config from yaml file
func LoadConfig(configFile string) (*AppConfig, error) {
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, errors.New(err)
	}

	var appConfig AppConfig
	err = viper.Unmarshal(&appConfig)
	if err != nil {
		return nil, errors.New(err)
	}

	return &appConfig, err
}
