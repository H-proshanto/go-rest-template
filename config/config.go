package config

import (
	"fmt"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var appOnce = sync.Once{}

type Application struct {
	Host   string `mapstructure:"HOST"`
	Port   string `mapstructure:"PORT"`
	DBUser string `mapstructure:"DB_USER"`
	DBName string `mapstructure:"DB_NAME"`
	DBPass string `mapstructure:"DB_PASS"`
	DBHost string `mapstructure:"DB_HOST"`
	DBPort string `mapstructure:"DB_PORT"`
}

var appConfig *Application

func loadApp() error {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(".env file not found")
	}

	viper.AutomaticEnv()

	appConfig = &Application{
		Host:   viper.GetString("HOST"),
		Port:   viper.GetString("PORT"),
		DBUser: viper.GetString("DB_USER"),
		DBName: viper.GetString("DB_NAME"),
		DBPass: viper.GetString("DB_PASS"),
		DBHost: viper.GetString("DB_HOST"),
		DBPort: viper.GetString("DB_PORT"),
	}

	return nil
}

func GetApp() *Application {
	appOnce.Do(func() {
		loadApp()
	})

	return appConfig
}
