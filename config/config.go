package config

import (
	"fmt"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var appOnce = sync.Once{}

type Application struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
}

var appConfig *Application

func loadApp() error {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(".env file not found")
	}

	viper.AutomaticEnv()

	appConfig = &Application{
		Host: viper.GetString("HOST"),
		Port: viper.GetString("PORT"),
	}

	return nil
}

func GetApp() *Application {
	appOnce.Do(func() {
		loadApp()
	})

	return appConfig
}
