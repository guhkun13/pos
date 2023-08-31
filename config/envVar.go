package config

type EnvironmentVariables struct {
	Database struct {
		Username string `mapstructure:"USERNAME"`
		Password string `mapstructure:"PASSWORD"`
		Host     string `mapstructure:"HOST"`
		Port     string `mapstructure:"PORT"`
		Name     string `mapstructure:"NAME"`
	} `mapstructure:"DATABASE"`
}
