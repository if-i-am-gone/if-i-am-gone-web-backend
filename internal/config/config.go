package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	Postgres struct {
		Host     string `validate:"required"`
		Port     int    `validate:"required"`
		DB_Name  string `validate:"required"`
		User     string `validate:"required"`
		Password string `validate:"required"`
	}
	Api struct {
		Hostname string `validate:"required"`
		Port     int    `validate:"required"`
	}
}

var cfg Config
var isLoaded bool = false

func Load() error {
	isLoaded = false

	CONFIG_FILE_NAME, configFileNameExists := os.LookupEnv("CONFIG_FILE_NAME")
	CONFIG_FILE_TYPE, configFileTypeExists := os.LookupEnv("CONFIG_FILE_TYPE")
	CONFIG_FILE_PATH, configFilePathExists := os.LookupEnv("CONFIG_FILE_PATH")

	if !configFileNameExists || !configFileTypeExists || !configFilePathExists {
		return fmt.Errorf("error while reading env variables about config file")
	}

	viper.SetConfigName(CONFIG_FILE_NAME)
	viper.SetConfigType(CONFIG_FILE_TYPE)
	viper.AddConfigPath(CONFIG_FILE_PATH)

	cfg = Config{}

	errReadConfig := viper.ReadInConfig()
	if errReadConfig != nil {
		return fmt.Errorf("unable to read config: %w", errReadConfig)
	}

	errUnmarshal := viper.Unmarshal(&cfg)
	if errUnmarshal != nil {
		return fmt.Errorf("unable to decode config into struct: %w", errUnmarshal)
	}

	errValidate := validate()
	if errValidate != nil {
		return errValidate
	}

	isLoaded = true
	return nil
}

func Get() (*Config, error) {
	if !isLoaded {
		return nil, fmt.Errorf("config is not loaded")
	}

	if errValidate := validate(); errValidate != nil {
		return nil, errValidate
	}

	return &cfg, nil
}

func validate() error {
	validate := validator.New()
	err := validate.Struct(cfg)

	if err != nil {
		return fmt.Errorf("config validation failed: %w", err)
	}
	return nil
}
