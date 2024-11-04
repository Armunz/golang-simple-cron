package config

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

const (
	CRON_TIME string = "CRON_TIME"
)

type Config struct {
	CronTime string `validate:"required"`
}

func New() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("failed to load .env files, loading host env instead")
	}

	cfg := Config{
		CronTime: os.Getenv(CRON_TIME),
	}

	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		log.Fatalln("config validation failed")
	}

	return cfg
}
