package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type EnvironmentVariables struct {
	Database struct {
		Username string `mapstructure:"USERNAME"`
		Password string `mapstructure:"PASSWORD"`
		Host     string `mapstructure:"HOST"`
		Port     string `mapstructure:"PORT"`
		Name     string `mapstructure:"NAME"`
	} `mapstructure:"DATABASE"`
}

func LoadEnv() (envVar EnvironmentVariables, err error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Error().Err(err).Msg("viper error read config")
	}

	err = viper.Unmarshal(&envVar)
	if err != nil {
		log.Error().Err(err).Msg("viper error unmarshal config")
	}

	return
}
