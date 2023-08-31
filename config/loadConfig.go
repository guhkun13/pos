package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func LoadEnvVar(path string) (envVar EnvironmentVariables, err error) {
	viper.AddConfigPath(path)

	viper.SetConfigName("app")
	viper.SetConfigType("env")

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
