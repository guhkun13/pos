package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

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
