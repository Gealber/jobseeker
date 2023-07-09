package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	Database struct {
		DSN     string
		TESTDSN string
	}
}

var (
	cfg *AppConfig
)

func Config() *AppConfig {
	if cfg == nil {
		loadConfig()
	}

	return cfg
}

func loadConfig() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	// Ignore config file not found, perhaps we will use environment variables.
	_ = viper.ReadInConfig()

	cfg = &AppConfig{}

	// Database.
	cfg.Database.DSN = viper.GetString("DATABASE_DSN")
	cfg.Database.TESTDSN = viper.GetString("DATABASE_TEST_DSN")
}
