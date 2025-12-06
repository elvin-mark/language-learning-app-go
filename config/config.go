package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Filepath string
}

type LlmConfig struct {
	Type    string
	Model   string
	BaseUrl string
}

type ServerConfig struct {
	Port int
}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Llm      LlmConfig
	LogLevel string `mapstructure:"log_level"`
}

func LoadConfig() (*Config, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
