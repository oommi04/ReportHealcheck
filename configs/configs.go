package configs

import (
	"github.com/joho/godotenv"
	"os"
)

type Configs struct {
	PORT               string
	CHANELID           string
	CHANELSECRET       string
	REDIRECTURL        string
	REPORTHEALCHECLURL string
}

func New() *Configs {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	return &Configs{
		os.Getenv("PORT"),
		os.Getenv("CHANELID"),
		os.Getenv("CHANELSECRET"),
		os.Getenv("REDIRECTURL"),
		os.Getenv("REPORTHEALCHECLURL"),
	}
}
