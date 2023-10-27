package config

import (
	"github.com/joho/godotenv"             // load .env file
	"github.com/kelseyhightower/envconfig" // manage configuration from env variables
	"log"
)

type Config struct {
	Db         Database
	ServerPort int `envconfig:"SERVER_PORT" default:"80"`
}

type Database struct {
	Host     string `envconfig:"DATABASE_HOST" required:"true"`
	Port     int    `envconfig:"DATABASE_PORT" required:"true"`
	User     string `envconfig:"DATABASE_USER" required:"true"`
	Password string `envconfig:"DATABASE_PASSWORD" required:"true"`
	Name     string `envconfig:"DATABASE_NAME" required:"true"`
}

func NewParsedConfig() (Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	cnfg := Config{}

	err = envconfig.Process("", &cnfg)
	if err != nil {
		log.Fatal(err)
	}
	return cnfg, err
}
