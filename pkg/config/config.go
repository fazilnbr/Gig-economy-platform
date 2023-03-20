package config

import (

	validator "github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost           string `mapstructure:"DB_HOST"`
	DBName           string `mapstructure:"DB_NAME"`
	DBUser           string `mapstructure:"DB_USER"`
	DBPort           string `mapstructure:"DB_PORT"`
	DBPassword       string `mapstructure:"DB_PASSWORD"`
	DBSOURCE         string `mapstructure:"DB_SOURCE"`
	SMTPPORT         string `mapstructure:"SMTP_PORT"`
	SMTPHOST         string `mapstructure:"SMTP_HOST"`
	SMTPPASSWORD     string `mapstructure:"SMTP_PASSWORD"`
	SMTPUSERNAME     string `mapstructure:"SMTP_USERNAME"`
	OauthStateString string `mapstructure:"OauthStateString"`
	ClientID         string `mapstructure:"ClientID"`
	ClientSecret     string `mapstructure:"ClientSecret"`
}

var envs = []string{
	"DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD", "DB_SOURCE", "SMTP_PORT", "SMTP_HOST", "SMTP_PASSWORD", "SMTP_USERNAME", "OauthStateString", "ClientID", "ClientSecret",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}
	return config, nil
}
