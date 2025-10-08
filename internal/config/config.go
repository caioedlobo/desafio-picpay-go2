package config

import (
	"github.com/spf13/viper"
	"sync"
)

var (
	config *Config
	once   sync.Once
)

type Config struct {
	Port        string `mapstructure:"PORT"`
	Environment string `mapstructure:"ENVIRONMENT"`
	AppName     string `mapstructure:"APP_NAME"`
	Debug       bool   `mapstructure:"DEBUG"`
	PostgresDSN string `mapstructure:"DB_POSTGRES_DSN"`
	DriverName  string `mapstructure:"DRIVER_NAME"`
}

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigName(".env")
		viper.SetConfigType("env")
		viper.AddConfigPath(".")
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&config); err != nil {
			panic(err)
		}
	})
	return config
}
