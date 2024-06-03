package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	NAME string
	PORT string

	DB_PORT     string
	DB_NAME     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string

	GEMINI_API_KEY string

	REDIS_ADDR     string
	REDIS_USERNAME string
	REDIS_PASSWORD string
}

var ENV *Config

func isProduction() bool {
	return os.Getenv("NAME") == "production"
}

func LoadConfig() {
	if isProduction() {
		ENV = &Config{
			NAME: os.Getenv("NAME"),
			PORT: os.Getenv("PORT"),

			DB_PORT:     os.Getenv("DB_PORT"),
			DB_NAME:     os.Getenv("DB_NAME"),
			DB_USERNAME: os.Getenv("DB_USERNAME"),
			DB_PASSWORD: os.Getenv("DB_PASSWORD"),
			DB_HOST:     os.Getenv("DB_HOST"),

			GEMINI_API_KEY: os.Getenv("GEMINI_API_KEY"),

			REDIS_ADDR:     os.Getenv("REDIS_ADDR"),
			REDIS_USERNAME: os.Getenv("REDIS_USERNAME"),
			REDIS_PASSWORD: os.Getenv("REDIS_PASSWORD"),
		}

		return
	}

	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		log.Fatal(err)
	}
}
